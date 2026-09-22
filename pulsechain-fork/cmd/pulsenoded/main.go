package main

import (
	"log"
	"net/http"

	"pulsechain-fork/genesis"
	"pulsechain-fork/internal/blockchain"
	"pulsechain-fork/internal/consensus"
	"pulsechain-fork/internal/ledger"
	"pulsechain-fork/internal/rpc"
	"github.com/ethereum/go-ethereum/rlp"
)

// Auteur : Martial Zinsou
// Point d'entrée principal du nœud PulseChain Fork

func main() {
	// Initialisation du bloc génèse
	gen := genesis.DefaultGenesis()
	log.Println("Initialisation du bloc génèse PulseChain")
	log.Printf("ChainID: %d, Coinbase: %s", gen.ChainID, gen.Coinbase)

	// Création de la base de données du ledger
	ledgerDB := ledger.NewStateDB()

	// Initialisation du consensus
	cons := consensus.NewProofOfAuthority()

	// Création de la blockchain
	bc := blockchain.NewBlockchain(gen, ledgerDB, cons)

	// Démarrage du serveur RPC
	rpcServer := rpc.NewServer(bc)
	http.Handle("/rpc", rpcServer)

	log.Println("Démarrage du nœud PulseChain Fork...")
	log.Printf("RPC disponible sur http://localhost:8545/rpc")
	log.Printf("Blockchain initialisée avec %d blocs", bc.BlockCount())

	// Démarrage du serveur HTTP
	if err := http.ListenAndServe(":8545", nil); err != nil {
		log.Fatalf("Erreur du serveur RPC: %v", err)
	}
}