package blockchain

import (
	"crypto/sha256"
	"encoding/binary"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/martialzinsou/pulsechain-fork/internal/consensus"
	"github.com/martialzinsou/pulsechain-fork/internal/ledger"
)

type Block struct {
	Number     uint64
	TimeStamp  int64
	ParentHash common.Hash
	Nonce      []byte
	Root       common.Hash
	TxsHash    common.Hash
	ReceiptsRoot common.Hash
	LogsBloom  []byte
	Difficulty *big.Int
	GasLimit   uint64
	GasUsed    uint64
	Coinbase   common.Address
	Transactions []*Transaction
	StateRoot  common.Hash
}

type Transaction struct {
	Hash      common.Hash
	Nonce     uint64
	GasPrice  *big.Int
	GasLimit  uint64
	To        common.Address
	Value     *big.Int
	Data      []byte
	V, R, S   *big.Int
}

func NewBlock(number uint64, timestamp int64, parentHash common.Hash, difficulty *big.Int, coinbase common.Address) *Block {
	return &Block{
		Number:     number,
		TimeStamp:  timestamp,
		ParentHash: parentHash,
		Difficulty: difficulty,
		Coinbase:   coinbase,
	}
}

func (b *Block) HashBlock() common.Hash {
	// Hash le bloc en utilisant RLP encoding et SHA256
	var encoded []byte
	
	// Encodage RLP basique
	rlpData := []interface{}{
		b.Number,
		b.TimeStamp,
		b.ParentHash,
		b.Nonce,
		b.Root,
		b.TxsHash,
		b.ReceiptsRoot,
		b.LogsBloom,
		b.Difficulty,
		b.GasLimit,
		b.GasUsed,
		b.Coinbase,
		b.Transactions,
		b.StateRoot,
	}
	
	encoded, err := rlp.EncodeToBytes(rlpData)
	if err != nil {
		// Fallback hachage manuel si RLP échoue
		return hashBlockData(encoded)
	}
	return hashBlockData(encoded)
}

func hashBlockData(data []byte) common.Hash {
	hash := sha256.Sum256(data)
	return common.Hash(hash[:])
}

func (b *Block) ValidateBlock(consensus *consensus.Consensus) bool {
	// Validation basique du bloc
	if b.Number == 0 && b.ParentHash == (common.Hash{}) {
		return true // Genesis block
	}
	
	// Vérifier la difficulté via l'algorithme de consensus
	target := consensus.CalculateTarget(b.Difficulty, b.TimeStamp, b.ParentHash)
	
	// Hasher le nonce et vérifier qu'il est inférieur à la target
	nonceHash := b.HashNonce()
	nonceBig := new(big.Int).SetBytes(nonceHash[:])
	
	return nonceBig.Cmp(target) <= 0
}

func (b *Block) HashNonce() []byte {
	// Hash du nonce pour la vérification de preuve
	hash := sha256.Sum256(append(b.Nonce, b.TimeStamp...))
	return hash[:]
}

func (b *Block) AddTransaction(tx *Transaction) {
	b.Transactions = append(b.Transactions, tx)
	b.TxsHash = b.CalculateTxsHash()
}

func (b *Block) CalculateTxsHash() common.Hash {
	// Calculer le hash de toutes les transactions
	var txHashes []common.Hash
	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.Hash)
	}
	
	if len(txHashes) == 0 {
		return common.Hash{}
	}
	
	// Hash simple en concaténant tous les hashes
	var data []byte
	for _, h := range txHashes {
		data = append(data, h[:]...)
	}
	
	hash := sha256.Sum256(data)
	return common.Hash(hash[:])
}

// GenesisBlock crée le bloc génèse
func GenesisBlock(coinbase common.Address) *Block {
	difficulty := big.NewInt(1)
	genesis := NewBlock(0, 0, common.Hash{}, difficulty, coinbase)
	
	// Création de la transaction de coinbase
	coinbaseTx := &Transaction{
		Nonce:  1,
		GasPrice: big.NewInt(0),
		GasLimit: uint64(21000),
		To:     coinbase,
		Value:  big.NewInt(1000000000000000000), // 1 ETH initial
		Data:   []byte{},
	}
	
	genesis.AddTransaction(coinbaseTx)
	genesis.Root = genesis.CalculateRoot()
	genesis.TxsHash = genesis.CalculateTxsHash()
	
	return genesis
}

func (b *Block) CalculateRoot() common.Hash {
	// Calculer l'état root à partir des transactions
	// Pour l'instant, retourne le hash des transactions
	if len(b.Transactions) == 0 {
		return common.Hash{}
	}
	
	// Agrégation simple des hashes de transactions
	var combinedHash common.Hash
	for i, tx := range b.Transactions {
		if i == 0 {
			combinedHash = tx.Hash
		} else {
			// Concaténer et hacher
			combinedHash = sha256Hash(append(combinedHash[:], tx.Hash[:]...))
		}
	}
	
	return combinedHash
}

func sha256Hash(data []byte) common.Hash {
	hash := sha256.Sum256(data)
	return common.Hash(hash[:])
}