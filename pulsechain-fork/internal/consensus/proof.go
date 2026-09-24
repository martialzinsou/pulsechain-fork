package consensus

// Auteur : Martial Zinsou
// Algorithme de consensus Proof-of-Authority (PoA) adapté pour PulseChain Fork

import (
	"crypto/sha256"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// Consensus interface définissant les méthodes nécessaires pour la validation des blocs
type Consensus interface {
	CalculateTarget(difficulty *big.Int, timestamp int64, parentHash common.Hash) *big.Int
	ValidateBlock(blockNumber uint64, timestamp int64, parentHash common.Hash, nonce []byte, difficulty *big.Int) bool
	GetDifficulty() *big.Int
	SetDifficulty(diff *big.Int)
}

// ProofOfAuthority implémente le consensus PoA avec validateur et temps de bloc fixe
type ProofOfAuthority struct {
	mu           sync.RWMutex
	authorities  map[common.Address]bool
	difficulty   *big.Int
	blockTimeSec int64
	period       uint64
}

// NewProofOfAuthority initialise le consensus avec les paramètres PulseChain
func NewProofOfAuthority() *ProofOfAuthority {
	poa := &ProofOfAuthority{
		authorities:  make(map[common.Address]bool),
		difficulty:   big.NewInt(1),
		blockTimeSec: 3, // PulseChain vise un block time rapide (~3s)
		period:       3,
	}
	return poa
}

// AddAuthority ajoute un validateur autorisé
func (p *ProofOfAuthority) AddAuthority(addr common.Address) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.authorities[addr] = true
}

// IsAuthority vérifie si une adresse est un validateur actif
func (p *ProofOfAuthority) IsAuthority(addr common.Address) bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.authorities[addr]
}

// CalculateTarget calcule la cible de difficulté
func (p *ProofOfAuthority) CalculateTarget(difficulty *big.Int, timestamp int64, parentHash common.Hash) *big.Int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	// En PoA PulseChain, la difficulté sert de score de signature
	target := new(big.Int).Set(difficulty)
	return target
}

// ValidateBlock valide les paramètres de consensus d'un bloc
func (p *ProofOfAuthority) ValidateBlock(blockNumber uint64, timestamp int64, parentHash common.Hash, nonce []byte, difficulty *big.Int) bool {
	if blockNumber == 0 {
		return true // Bloc Genesis toujours valide
	}

	// Vérification de la cohérence temporelle
	now := time.Now().Unix()
	if timestamp > now+15 { // Tolérance max 15s dans le futur
		return false
	}

	return true
}

// GetDifficulty retourne la difficulté actuelle
func (p *ProofOfAuthority) GetDifficulty() *big.Int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return new(big.Int).Set(p.difficulty)
}

// SetDifficulty définit la difficulté courante
func (p *ProofOfAuthority) SetDifficulty(diff *big.Int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.difficulty = new(big.Int).Set(diff)
}

// HashNonce calcule le hash du nonce pour la validation
func HashNonce(nonce []byte, timestamp int64) []byte {
	buf := append(nonce, byte(timestamp), byte(timestamp>>8), byte(timestamp>>16), byte(timestamp>>24))
	h := sha256.Sum256(buf)
	return h[:]
}
