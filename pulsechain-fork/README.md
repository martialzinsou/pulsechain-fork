# PulseChain Fork

Un prototype de blockchain fork de PulseChain en Go.

## Structure du projet

```
pulsechain-fork/
├── cmd/                  # Points d'entrée d'application
├── genesis/              # Configuration du bloc génèse
├── go.mod               # Dépendances Go
└── internal/
    ├── blockchain/      # Moteur de blockchain
    ├── consensus/       # Algorithme de consensus
    ├── ledger/          # Gestion du registre
    └── rpc/             # Points de terminaison RPC
```

## Démarrage rapide

```bash
# Installation
go mod tidy

# Lancement du nœud
go run ./cmd/pulsenoded
```

## Configuration du Fork

Le projet fork de PulseChain utilise les paramètres suivants :

- **ChainID**: Identifiant du réseau (par défaut: 1)
- **Coinbase**: Adresse du mineur par défaut
- **Hard Forks**: Homestead, EIP155, EIP158, Byzantium, Constantinople activés par défaut

## Points de terminaison RPC

Le serveur RPC est disponible sur le port 8545 par défaut. Les méthodes supportées :

- `eth_blockNumber` - Obtenir le numéro du dernier bloc
- `eth_getBalance` - Obtenir le solde d'une adresse
- `eth_getBlockByNumber` - Obtenir les informations d'un bloc

## Développement

Pour ajouter de nouvelles fonctionnalités :

1. Modifier le fichier correspondant dans `internal/`
2. Mettre à jour la configuration du génèse si nécessaire
3. Tester avec `go test ./...`