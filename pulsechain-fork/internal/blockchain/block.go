package blockchain

import (
	"crypto/sha256"
	"encoding/binary"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/martialzinsou/pulsechain-fork/internal/consensus"
	"github.com/martialzinsou/pulsechain-fork/internal/ledger"
)

// Auteur : Martial Zinsou
// Structures et fonctions de hachage pour la blockchain PulseChain