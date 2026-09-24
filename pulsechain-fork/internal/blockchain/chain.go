package blockchain

// Auteur : Martial Zinsou
// Moteur de blockchain PulseChain Fork avec gestion d'état, consensus et persistance

import (
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"pulsechain-fork/genesis"
	"pulsechain-fork/internal/consensus"
	"pulsechain-fork/internal/ledger"
)

// Blockchain structure principale représentant le registre de la chaîne
type Blockchain struct {
	mu        sync.RWMutex
	blocks    []*Block
	stateDB   *ledger.StateDB
	consensus consensus.Consensus
	chainID   uint64
	coinbase  common.Address
}

// NewBlockchain instancie la blockchain avec la configuration genesis PulseChain
func NewBlockchain(gen *genesis.Genesis, stateDB *ledger.StateDB, cons consensus.Consensus) *Blockchain {
	coinbaseAddr := common.HexToAddress(gen.Coinbase)
	genBlock := GenesisBlock(coinbaseAddr, gen.ChainID)

	// Attribution initiale au coinbase dans l'état
	initialBalance := new(big.Int).Mul(big.NewInt(1000000), big.NewInt(1e18))
	stateDB.SetBalance(coinbaseAddr, initialBalance)
	stateDB.UpdateState(genBlock.Hash)

	bc := &Blockchain{
		blocks:    []*Block{genBlock},
		stateDB:   stateDB,
		consensus: cons,
		chainID:   gen.ChainID,
		coinbase:  coinbaseAddr,
	}

	return bc
}

// AddBlock valide et intègre un nouveau bloc à la chaîne
func (bc *Blockchain) AddBlock(block *Block) error {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	lastBlock := bc.blocks[len(bc.blocks)-1]

	// Vérification du séquençage des blocs
	if block.Number != lastBlock.Number+1 {
		return fmt.Errorf("numéro de bloc incohérent : reçu %d, attendu %d", block.Number, lastBlock.Number+1)
	}

	// Vérification du parent hash
	if block.ParentHash != lastBlock.Hash {
		return errors.New("parentHash ne correspond pas au hash du dernier bloc")
	}

	// Validation du bloc via le consensus
	if !block.Validate(bc.consensus) {
		return errors.New("échec de validation du consensus pour ce bloc")
	}

	// Traitement des transactions et application de l'état
	for _, tx := range block.Transactions {
		if tx.From != (common.Address{}) {
			totalCost := new(big.Int).Add(tx.Value, new(big.Int).Mul(tx.GasPrice, big.NewInt(int64(tx.GasLimit))))
			if err := bc.stateDB.SubBalance(tx.From, totalCost); err != nil {
				return fmt.Errorf("transaction invalide : %v", err)
			}
			bc.stateDB.IncrementNonce(tx.From)
		}
		if tx.To != nil {
			bc.stateDB.AddBalance(*tx.To, tx.Value)
		}
	}

	// Récompense de minage au validateur/coinbase
	blockReward := new(big.Int).Mul(big.NewInt(2), big.NewInt(1e18)) // 2 PLS de récompense
	bc.stateDB.AddBalance(block.Coinbase, blockReward)

	// Mise à jour de la racine d'état
	if err := bc.stateDB.UpdateState(block.Hash); err != nil {
		return fmt.Errorf("erreur de mise à jour d'état : %v", err)
	}
	block.StateRoot = bc.stateDB.StateRoot()

	bc.blocks = append(bc.blocks, block)
	return nil
}

// MineBlock produit un nouveau bloc au sommet de la chaîne
func (bc *Blockchain) MineBlock(txs []*Transaction) (*Block, error) {
	bc.mu.RLock()
	lastBlock := bc.blocks[len(bc.blocks)-1]
	diff := bc.consensus.GetDifficulty()
	coinbase := bc.coinbase
	bc.mu.RUnlock()

	newBlock := NewBlock(
		lastBlock.Number+1,
		time.Now().Unix(),
		lastBlock.Hash,
		diff,
		coinbase,
	)

	for _, tx := range txs {
		newBlock.AddTransaction(tx)
	}

	if err := bc.AddBlock(newBlock); err != nil {
		return nil, err
	}

	return newBlock, nil
}

// GetBlockByNumber retourne le bloc correspondant au numéro indiqué
func (bc *Blockchain) GetBlockByNumber(number uint64) (*Block, error) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	if number >= uint64(len(bc.blocks)) {
		return nil, errors.New("bloc introuvable")
	}
	return bc.blocks[number], nil
}

// GetBlockByHash recherche un bloc par son hash
func (bc *Blockchain) GetBlockByHash(hash common.Hash) (*Block, error) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	for _, b := range bc.blocks {
		if b.Hash == hash {
			return b, nil
		}
	}
	return nil, errors.New("bloc non trouvé pour ce hash")
}

// BlockCount retourne le nombre total de blocs minés
func (bc *Blockchain) BlockCount() uint64 {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	return uint64(len(bc.blocks))
}

// GetLastBlock retourne le bloc à la tête de la chaîne
func (bc *Blockchain) GetLastBlock() *Block {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	if len(bc.blocks) == 0 {
		return nil
	}
	return bc.blocks[len(bc.blocks)-1]
}

// GetBalance retourne le solde d'un compte
func (bc *Blockchain) GetBalance(addr common.Address) *big.Int {
	return bc.stateDB.GetBalance(addr)
}

// GetChainID retourne l'identifiant du réseau PulseChain Fork
func (bc *Blockchain) GetChainID() uint64 {
	return bc.chainID
}
