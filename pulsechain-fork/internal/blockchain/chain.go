package blockchain

import (
	"errors"
	"fmt"
	"sync"
	
	"github.com/ethereum/go-ethereum/common"
	"github.com/martialzinsou/pulsechain-fork/internal/consensus"
	"github.com/martialzinsou/pulsechain-fork/internal/ledger"
)

type Blockchain struct {
	mu         sync.Mutex
	Blocks     []*Block
	StateDB    *ledger.StateDB
	Consensus  *consensus.Consensus
	ChainID    uint64
}

func NewBlockchain(gen *genesis.Genesis, stateDB *ledger.StateDB, cons *consensus.Consensus) *Blockchain {
	// Création du bloc génèse
	genesisBlock := GenesisBlock(common.HexToAddress(gen.Coinbase))
	genesisBlock.HashBlock()
	genesisBlock.ValidateBlock(cons)
	
	return &Blockchain{
		Blocks:     []*Block{genesisBlock},
		StateDB:    stateDB,
		Consensus:  cons,
		ChainID:    gen.ChainID,
	}
}

func (bc *Blockchain) AddBlock(block *Block) error {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	// validation du bloc
	if !block.ValidateBlock(bc.Consensus) {
		return errors.New("bloc invalide : échec de la validation")
	}

	// Vérifier que c'est le bon bloc suivant
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	if block.Number != lastBlock.Number+1 {
		return errors.New("numéro de bloc incorrect")
	}

	// Mettre à jour l'état
	if err := bc.StateDB.UpdateState(block); err != nil {
		return fmt.Errorf("échec de la mise à jour de l'état: %v", err)
	}

	bc.Blocks = append(bc.Blocks, block)
	return nil
}

func (bc *Blockchain) GetBlock(number uint64) (*Block, error) {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	if number >= uint64(len(bc.Blocks)) {
		return nil, errors.New("bloc non trouvé")
	}

	return bc.Blocks[number], nil
}

func (bc *Blockchain) GetBlockCount() uint64 {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	return uint64(len(bc.Blocks))
}

func (bc *Blockchain) GetLastBlock() *Block {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	
	if len(bc.Blocks) == 0 {
		return nil
	}
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) GetBalance(address common.Address) *big.Int {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	
	return bc.StateDB.GetBalance(address)
}

func (bc *Blockchain) GetAccount(address common.Address) *ledger.Account {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	
	return bc.StateDB.GetAccount(address)
}