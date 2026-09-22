package blockchain

import (
	"errors"
	"fmt"
	"sync"
	
	"github.com/ethereum/go-ethereum/common"
	"github.com/martialzinsou/pulsechain-fork/internal/consensus"
	"github.com/martialzinsou/pulsechain-fork/internal/ledger"
)

// Auteur : Martial Zinsou
// Moteur de blockchain PulseChain avec gestion d'état et validation de blocs