package blockchain

// Auteur : Martial Zinsou
// Structures et fonctions de hachage pour la blockchain PulseChain Fork

import (
	"crypto/sha256"
	"encoding/binary"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"pulsechain-fork/internal/consensus"
)

// Transaction représente une transaction sur le réseau PulseChain Fork
type Transaction struct {
	Hash     common.Hash     `json:"hash"`
	Nonce    uint64          `json:"nonce"`
	From     common.Address  `json:"from"`
	To       *common.Address `json:"to"`
	Value    *big.Int        `json:"value"`
	GasLimit uint64          `json:"gasLimit"`
	GasPrice *big.Int        `json:"gasPrice"`
	Data     []byte          `json:"data"`
	V        *big.Int        `json:"v"`
	R        *big.Int        `json:"r"`
	S        *big.Int        `json:"s"`
}

// Block représente un bloc dans la blockchain PulseChain Fork
type Block struct {
	Number       uint64          `json:"number"`
	Hash         common.Hash     `json:"hash"`
	ParentHash   common.Hash     `json:"parentHash"`
	Nonce        []byte          `json:"nonce"`
	TimeStamp    int64           `json:"timestamp"`
	Difficulty   *big.Int        `json:"difficulty"`
	GasLimit     uint64          `json:"gasLimit"`
	GasUsed      uint64          `json:"gasUsed"`
	Coinbase     common.Address  `json:"miner"`
	StateRoot    common.Hash     `json:"stateRoot"`
	TxsHash      common.Hash     `json:"transactionsRoot"`
	ReceiptsRoot common.Hash     `json:"receiptsRoot"`
	Transactions []*Transaction  `json:"transactions"`
}

// NewBlock instancie un nouveau bloc avec les valeurs spécifiées
func NewBlock(number uint64, timestamp int64, parentHash common.Hash, difficulty *big.Int, coinbase common.Address) *Block {
	if timestamp == 0 {
		timestamp = time.Now().Unix()
	}
	if difficulty == nil {
		difficulty = big.NewInt(1)
	}
	b := &Block{
		Number:       number,
		TimeStamp:    timestamp,
		ParentHash:   parentHash,
		Difficulty:   new(big.Int).Set(difficulty),
		GasLimit:     30000000, // PulseChain gas limit élevé
		GasUsed:      0,
		Coinbase:     coinbase,
		Nonce:        make([]byte, 8),
		Transactions: make([]*Transaction, 0),
	}
	b.Hash = b.CalculateHash()
	return b
}

// CalculateHash génère l'empreinte cryptographique SHA-256 du bloc
func (b *Block) CalculateHash() common.Hash {
	hasher := sha256.New()

	numBuf := make([]byte, 8)
	binary.BigEndian.PutUint64(numBuf, b.Number)
	hasher.Write(numBuf)

	timeBuf := make([]byte, 8)
	binary.BigEndian.PutUint64(timeBuf, uint64(b.TimeStamp))
	hasher.Write(timeBuf)

	hasher.Write(b.ParentHash.Bytes())
	hasher.Write(b.Coinbase.Bytes())
	if b.Difficulty != nil {
		hasher.Write(b.Difficulty.Bytes())
	}
	hasher.Write(b.StateRoot.Bytes())
	hasher.Write(b.TxsHash.Bytes())
	hasher.Write(b.Nonce)

	return common.BytesToHash(hasher.Sum(nil))
}

// AddTransaction insère une transaction dans le bloc et actualise le tx root
func (b *Block) AddTransaction(tx *Transaction) {
	b.Transactions = append(b.Transactions, tx)
	b.TxsHash = b.CalculateTxsHash()
	b.Hash = b.CalculateHash()
}

// CalculateTxsHash calcule la racine de l'arbre des transactions
func (b *Block) CalculateTxsHash() common.Hash {
	if len(b.Transactions) == 0 {
		return common.Hash{}
	}

	hasher := sha256.New()
	for _, tx := range b.Transactions {
		hasher.Write(tx.Hash.Bytes())
	}
	return common.BytesToHash(hasher.Sum(nil))
}

// Validate valide le bloc avec les règles du consensus
func (b *Block) Validate(cons consensus.Consensus) bool {
	if b.Number == 0 && b.ParentHash == (common.Hash{}) {
		return true // Bloc genesis valide par convention
	}
	return cons.ValidateBlock(b.Number, b.TimeStamp, b.ParentHash, b.Nonce, b.Difficulty)
}

// GenesisBlock crée le bloc génèse officiel pour le fork PulseChain
func GenesisBlock(coinbase common.Address, chainID uint64) *Block {
	difficulty := big.NewInt(1)
	genesis := &Block{
		Number:       0,
		ParentHash:   common.Hash{},
		TimeStamp:    1684000000, // Date historique approximative du fork PulseChain
		Difficulty:   difficulty,
		GasLimit:     30000000,
		GasUsed:      0,
		Coinbase:     coinbase,
		Nonce:        make([]byte, 8),
		Transactions: make([]*Transaction, 0),
		StateRoot:    common.HexToHash("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"),
	}

	// Attribution initiale (Airdrop 1:1 PulseChain)
	coinbaseTx := &Transaction{
		Hash:     common.HexToHash("0x1111111111111111111111111111111111111111111111111111111111111111"),
		Nonce:    0,
		From:     common.Address{},
		To:       &coinbase,
		Value:    new(big.Int).Mul(big.NewInt(1000000), big.NewInt(1e18)), // 1,000,000 PLS
		GasLimit: 21000,
		GasPrice: big.NewInt(0),
	}

	genesis.Transactions = append(genesis.Transactions, coinbaseTx)
	genesis.TxsHash = genesis.CalculateTxsHash()
	genesis.Hash = genesis.CalculateHash()
	return genesis
}
