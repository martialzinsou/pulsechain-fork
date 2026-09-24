package main

// Auteur : Martial Zinsou
// Point d'entrée principal du nœud PulseChain Fork

import (
	"log"
	"net/http"

	"pulsechain-fork/genesis"
	"pulsechain-fork/internal/blockchain"
	"pulsechain-fork/internal/consensus"
	"pulsechain-fork/internal/ledger"
	"pulsechain-fork/internal/rpc"
)

func main() {
	// Initialisation du bloc génèse
	gen := genesis.DefaultGenesis()
	log.Println("======================================================")
	log.Println("     PULSECHAIN FORK NODE - MARTIAL ZINSOU            ")
	log.Println("======================================================")
	log.Printf("[+] Initialisation du bloc génèse PulseChain...")
	log.Printf("[+] ChainID: %d | Coinbase: %s", gen.ChainID, gen.Coinbase)

	// Création de la base de données du ledger
	ledgerDB := ledger.NewStateDB()

	// Initialisation du consensus Proof-of-Authority
	cons := consensus.NewProofOfAuthority()

	// Création de la blockchain
	bc := blockchain.NewBlockchain(gen, ledgerDB, cons)

	// Initialisation du serveur RPC JSON-RPC 2.0
	rpcServer := rpc.NewServer(bc)
	http.Handle("/", rpcServer)
	http.Handle("/rpc", rpcServer)

	log.Println("[+] Serveur JSON-RPC prêt à recevoir les connexions")
	log.Printf("[+] RPC Endpoint: http://localhost:8545 (compatible MetaMask/Ethers)")
	log.Printf("[+] Blocs initialisés: %d (Bloc Genesis actif)", bc.BlockCount())
	log.Println("[+] Démarrage de l'écoute sur le port :8545 ...")

	// Démarrage du serveur HTTP
	if err := http.ListenAndServe(":8545", nil); err != nil {
		log.Fatalf("[-] Erreur critique du serveur RPC: %v", err)
	}
}
