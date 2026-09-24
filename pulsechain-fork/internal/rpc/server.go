package rpc

// Auteur : Martial Zinsou
// Serveur JSON-RPC 2.0 compatible Ethereum / PulseChain

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"pulsechain-fork/internal/blockchain"
)

// Request représente une requête standard JSON-RPC 2.0
type Request struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	ID      interface{}     `json:"id"`
}

// Response représente une réponse standard JSON-RPC 2.0
type Response struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RPCError   `json:"error,omitempty"`
	ID      interface{} `json:"id"`
}

// RPCError représente les détails d'une erreur JSON-RPC
type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Server encapsule les appels RPC vers la blockchain
type Server struct {
	bc *blockchain.Blockchain
}

// NewServer crée une nouvelle instance de serveur RPC
func NewServer(bc *blockchain.Blockchain) *Server {
	return &Server{bc: bc}
}

// ServeHTTP gère les requêtes HTTP JSON-RPC entrantes
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// CORS Headers pour compatibilité Web3 / MetaMask / DApps
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		s.writeError(w, nil, -32700, "Parse error")
		return
	}

	var req Request
	if err := json.Unmarshal(body, &req); err != nil {
		s.writeError(w, nil, -32700, "Parse error")
		return
	}

	res := s.handleMethod(req)
	json.NewEncoder(w).Encode(res)
}

func (s *Server) handleMethod(req Request) Response {
	switch req.Method {
	case "web3_clientVersion":
		return Response{JSONRPC: "2.0", ID: req.ID, Result: "PulseChain-Fork/v1.0.0-MartialZinsou/darwin-amd64/go1.21"}

	case "net_version":
		chainID := s.bc.GetChainID()
		return Response{JSONRPC: "2.0", ID: req.ID, Result: strconv.FormatUint(chainID, 10)}

	case "eth_chainId":
		chainID := s.bc.GetChainID()
		return Response{JSONRPC: "2.0", ID: req.ID, Result: fmt.Sprintf("0x%x", chainID)}

	case "eth_blockNumber":
		count := s.bc.BlockCount()
		if count == 0 {
			return Response{JSONRPC: "2.0", ID: req.ID, Result: "0x0"}
		}
		return Response{JSONRPC: "2.0", ID: req.ID, Result: fmt.Sprintf("0x%x", count-1)}

	case "eth_mining":
		return Response{JSONRPC: "2.0", ID: req.ID, Result: false}

	case "eth_gasPrice":
		return Response{JSONRPC: "2.0", ID: req.ID, Result: "0x3b9aca00"} // 1 Gwei

	case "eth_getBalance":
		var params []string
		if err := json.Unmarshal(req.Params, &params); err != nil || len(params) < 1 {
			return Response{JSONRPC: "2.0", ID: req.ID, Error: &RPCError{Code: -32602, Message: "Invalid params"}}
		}
		addr := common.HexToAddress(params[0])
		bal := s.bc.GetBalance(addr)
		return Response{JSONRPC: "2.0", ID: req.ID, Result: fmt.Sprintf("0x%x", bal)}

	case "eth_getBlockByNumber":
		var params []interface{}
		if err := json.Unmarshal(req.Params, &params); err != nil || len(params) < 1 {
			return Response{JSONRPC: "2.0", ID: req.ID, Error: &RPCError{Code: -32602, Message: "Invalid params"}}
		}

		var blockNum uint64
		tag, ok := params[0].(string)
		if ok {
			if tag == "latest" || tag == "pending" {
				blockNum = s.bc.BlockCount() - 1
			} else if strings.HasPrefix(tag, "0x") {
				val, _ := strconv.ParseUint(strings.TrimPrefix(tag, "0x"), 16, 64)
				blockNum = val
			}
		}

		block, err := s.bc.GetBlockByNumber(blockNum)
		if err != nil {
			return Response{JSONRPC: "2.0", ID: req.ID, Result: nil}
		}

		return Response{JSONRPC: "2.0", ID: req.ID, Result: map[string]interface{}{
			"number":           fmt.Sprintf("0x%x", block.Number),
			"hash":             block.Hash.Hex(),
			"parentHash":       block.ParentHash.Hex(),
			"timestamp":        fmt.Sprintf("0x%x", block.TimeStamp),
			"miner":            block.Coinbase.Hex(),
			"gasLimit":         fmt.Sprintf("0x%x", block.GasLimit),
			"gasUsed":          fmt.Sprintf("0x%x", block.GasUsed),
			"difficulty":       fmt.Sprintf("0x%x", block.Difficulty),
			"transactionsRoot": block.TxsHash.Hex(),
			"stateRoot":        block.StateRoot.Hex(),
			"transactions":     block.Transactions,
		}}

	default:
		return Response{
			JSONRPC: "2.0",
			ID:      req.ID,
			Error:   &RPCError{Code: -32601, Message: fmt.Sprintf("Method %s not found", req.Method)},
		}
	}
}

func (s *Server) writeError(w http.ResponseWriter, id interface{}, code int, message string) {
	res := Response{
		JSONRPC: "2.0",
		ID:      id,
		Error:   &RPCError{Code: code, Message: message},
	}
	json.NewEncoder(w).Encode(res)
}
