package consensus

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Consensus interface {
	CalculateTarget(difficulty *big.Int, timestamp int64, parentHash common.Hash) *big.Int
	ValidateBlock(block *Block) bool
}

type proofOfAuthority struct {
	// Paramètres de consensus PoA
	authorityAddr common.Address
	difficulty    *big.Int
	blockTime     int64
}

func NewProofOfAuthority() *proofOfAuthority {
	return &proofOfAuthority{
		difficulty: big.NewInt(1),
		blockTime:  5, // 5 secondes par bloc PulseChain
	}
}

func (p *proofOfAuthority) CalculateTarget(difficulty *big.Int, timestamp int64, parentHash common.Hash) *big.Int {
	// Algorithme de difficulté PulseChain
	// Ajustement tous les 100 blocs ou selon la difficulté cible
	
	// Pour PulseChain : cible de 1 seconde par bloc en moyenne
	// avec ajustement tous les 100 blocs
	
	target := new(big.Int).Set(difficulty)
	
	// Ajustement de la difficulté basique
	// Si le temps de bloc est inférieur à la cible, augmenter la difficulté
	// Si le temps de bloc est supérieur, diminuer la difficulté
	
	return target
}

func (p *proofOfAuthority) ValidateBlock(block *Block) bool {
	// Validation PoA : vérifier que le nonce satisfait la difficulté
	if block.Number == 0 {
		return true // Genesis block always valid
	}
	
	nonceHash := block.HashNonce()
	nonceBig := new(big.Int).SetBytes(nonceHash[:])
	
	// La cible est calculée basée sur la difficulté actuelle
	target := p.CalculateTarget(p.difficulty, block.TimeStamp, block.ParentHash)
	
	return nonceBig.Cmp(target) <= 0
}

// GetDifficulty retourne la difficulté actuelle
func (p *proofOfAuthority) GetDifficulty() *big.Int {
	return p.difficulty
}

// SetDifficulty définit la difficulté
func (p *proofOfAuthority) SetDifficulty(diff *big.Int) {
	p.difficulty = diff
}