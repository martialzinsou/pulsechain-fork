package ledger

// Auteur : Martial Zinsou
// Gestion des comptes, soldes et base de données d'état pour PulseChain Fork

import (
	"crypto/sha256"
	"errors"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

// Account représente un compte sur la blockchain PulseChain (Fork)
type Account struct {
	Address     common.Address `json:"address"`
	Balance     *big.Int       `json:"balance"`
	Nonce       uint64         `json:"nonce"`
	CodeHash    []byte         `json:"codeHash"`
	StorageHash []byte         `json:"storageHash"`
}

// StateDB gère l'état global des comptes et des contrats
type StateDB struct {
	mu       sync.RWMutex
	accounts map[common.Address]*Account
	stateRoot common.Hash
}

// NewStateDB initialise une nouvelle base de données d'état
func NewStateDB() *StateDB {
	return &StateDB{
		accounts: make(map[common.Address]*Account),
	}
}

// GetAccount récupère un compte par son adresse (crée le compte s'il n'existe pas)
func (s *StateDB) GetAccount(addr common.Address) *Account {
	s.mu.Lock()
	defer s.mu.Unlock()

	acc, exists := s.accounts[addr]
	if !exists {
		acc = &Account{
			Address: addr,
			Balance: big.NewInt(0),
			Nonce:   0,
		}
		s.accounts[addr] = acc
	}
	return acc
}

// GetBalance retourne le solde d'une adresse
func (s *StateDB) GetBalance(addr common.Address) *big.Int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	acc, exists := s.accounts[addr]
	if !exists {
		return big.NewInt(0)
	}
	return new(big.Int).Set(acc.Balance)
}

// SetBalance définit le solde d'une adresse
func (s *StateDB) SetBalance(addr common.Address, amount *big.Int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	acc, exists := s.accounts[addr]
	if !exists {
		acc = &Account{
			Address: addr,
			Balance: new(big.Int).Set(amount),
			Nonce:   0,
		}
		s.accounts[addr] = acc
	} else {
		acc.Balance = new(big.Int).Set(amount)
	}
}

// AddBalance crédite un compte
func (s *StateDB) AddBalance(addr common.Address, amount *big.Int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	acc, exists := s.accounts[addr]
	if !exists {
		acc = &Account{
			Address: addr,
			Balance: new(big.Int).Set(amount),
			Nonce:   0,
		}
		s.accounts[addr] = acc
	} else {
		acc.Balance = new(big.Int).Add(acc.Balance, amount)
	}
}

// SubBalance débite un compte après vérification de solde suffisant
func (s *StateDB) SubBalance(addr common.Address, amount *big.Int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	acc, exists := s.accounts[addr]
	if !exists || acc.Balance.Cmp(amount) < 0 {
		return errors.New("solde insuffisant")
	}

	acc.Balance = new(big.Int).Sub(acc.Balance, amount)
	return nil
}

// IncrementNonce incrémente le nonce d'un compte
func (s *StateDB) IncrementNonce(addr common.Address) {
	s.mu.Lock()
	defer s.mu.Unlock()

	acc, exists := s.accounts[addr]
	if exists {
		acc.Nonce++
	}
}

// StateRoot retourne le hash de la racine d'état
func (s *StateDB) StateRoot() common.Hash {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.stateRoot
}

// UpdateState met à jour la racine d'état après application d'un bloc
func (s *StateDB) UpdateState(blockHash common.Hash) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Hachage cumulé de l'état
	hasher := sha256.New()
	for addr, acc := range s.accounts {
		hasher.Write(addr.Bytes())
		hasher.Write(acc.Balance.Bytes())
	}
	hasher.Write(blockHash.Bytes())
	s.stateRoot = common.BytesToHash(hasher.Sum(nil))
	return nil
}
