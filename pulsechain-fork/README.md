# PulseChain Fork

Un prototype complet de blockchain fork de PulseChain implémenté en Go. Ce projet sert de base pour créer votre propre nœud PulseChain avec une compatibilité totale avec l'écosystème Ethereum.

## Auteur

**Martial Zinsou** - https://github.com/martialzinsou

## Aperçu

Ce projet fork de PulseChain fournit :
- Un moteur de blockchain complet en Go
- Une compatibilité RPC Ethereum (JSON-RPC)
- La configuration du bloc génèse PulseChain
- Le support des hard forks (Homestead, EIP155, EIP158, Byzantium, Constantinople)
- Une architecture modulaire pour une extension facile

## Structure du Projet

```
pulsechain-fork/
├── cmd/                    # Points d'entrée d'application
│   └── pulsenoded/         # Commande principale du nœud
├── genesis/                # Configuration du bloc génèse
│   └── genesis.go          # Paramètres du bloc génèse
├── internal/               # Code interne modularisé
│   ├── blockchain/         # Moteur de blockchain
│   │   ├── block.go        # Structure et validation des blocs
│   │   └── chain.go        # Gestion de la chaîne de blocs
│   ├── consensus/          # Algorithme de consensus
│   │   └── proof.go        # Logique de consensus PoA
│   ├── ledger/             # Gestion du registre
│   │   └── account.go      # Gestion des comptes et soldes
│   └── rpc/                # Points de terminaison RPC
│       └── server.go       # Serveur JSON-RPC
├── go.mod                  # Dépendances Go
├── go.sum                  # Checksums de dépendances
└── docs/                   # Documentation détaillée
    ├── architecture.md     # Architecture du système
    └── api.md              # Documentation de l'API RPC
```

## Guide Utilisateur complet

### Installation

#### Prérequis

- Go 1.21 ou supérieur
- Git

#### Étapes d'installation

```bash
# 1. Cloner le dépôt
git clone https://github.com/martialzinsou/pulsechain-fork.git
cd pulsechain-fork

# 2. Installer les dépendances
go mod tidy

# 3. Compiler le nœud
go build -o pulsenoded ./cmd/pulsenoded
```

### Lancement du Nœud

```bash
# Démarrer le nœud en mode démon
./pulsenoded

# Ou en mode interactif pour le développement
go run ./cmd/pulsenoded
```

Le nœud démarre par défaut sur :
- **Port RPC** : 8545
- **Port P2P** : 30303
- **RPC CORS** : * (pour le développement)

### Configuration du Fork PulseChain

Le bloc génèse est configuré avec les paramètres PulseChain par défaut :

```go
type Genesis struct {
    ChainID       uint64  # Identifiant du réseau (1 pour PulseChain principal)
    Coinbase      string  # Adresse du contrat coinbase (miner)
    Homestead     bool    # Activation hard fork Homestead
    EIP155Enabled bool    # Activation EIP-155 (replay protection)
    EIP158Enabled bool    # Activation EIP-158 (gas pool)
    Byzantium     bool    # Activation hard fork Byzantium
    Constantinople bool   # Activation hard fork Constantinople
}

func DefaultGenesis() *Genesis {
    return &Genesis{
        ChainID:       1,
        Coinbase:      "0x0000000000000000000000000000000000000000",
        Homestead:     true,
        EIP155Enabled: true,
        EIP158Enabled: true,
        Byzantium:     true,
        Constantinople: true,
    }
}
```

### Points de terminaison RPC

Le serveur RPC expose toutes les méthodes standards Ethereum :

#### Requêtes de base

| Méthode | Description | Exemple |
|---------|-------------|---------|
| `eth_blockNumber` | Obtenir le numéro du dernier bloc | `curl -X POST http://localhost:8545 -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'` |
| `eth_getBalance` | Obtenir le solde d'une adresse | `curl -X POST http://localhost:8545 -d '{"jsonrpc":"2.0","method":"eth_getBalance","params":["0xaddr", "latest"],"id":1}'` |
| `eth_getBlockByNumber` | Obtenir les informations d'un bloc | `curl -X POST http://localhost:8545 -d '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["latest", false],"id":1}'` |

#### Méthodes supplémentaires

- `eth_getBlockTransactionCountByHash` - Nombre de transactions dans un bloc
- `eth_getBlockTransactionCountByNumber` - Nombre de transactions dans un bloc (par numéro)
- `eth_getCode` - Code de contrat à une adresse
- `eth_getLogs` - Récupérer les événements (logs)
- `eth_sendTransaction` - Envoyer une transaction
- `eth_call` - Appel de contrat lecture seule
- `eth_estimateGas` - Estimation du gas pour une transaction

### Développement

#### Ajouter de nouvelles fonctionnalités

1. **Moteur de blockchain** (`internal/blockchain/`):
   - Modifier `block.go` pour ajouter la structure des blocs
   - Modifier `chain.go` pour la validation et le stockage de la chaîne

2. **Consensus** (`internal/consensus/`):
   - Implémenter les règles de consensus PulseChain
   - Ajuster les paramètres de difficulté et de temps de bloc

3. **RPC** (`internal/rpc/`):
   - Ajouter de nouvelles méthodes dans `server.go`
   - Mettre à jour le schéma de réponse JSON

#### Structure des blocs

```go
type Block struct {
    Number     uint64
    Hash       []byte
    ParentHash []byte
    Nonce      []byte
    Timestamp  int64
    Difficulty *big.Int
    GasLimit   uint64
    GasUsed    uint64
    Coinbase   common.Address
    Transactions []*Transaction
    StateRoot  common.Hash
}
```

#### Gestion des comptes

```go
type Account struct {
    Address common.Address
    Balance *big.Int
    Nonce   uint64
    CodeHash []byte
    StorageHash []byte
}
```

### Dépannage

#### Problèmes courants

1. **Port déjà utilisé** : Modifier les ports dans `cmd/pulsenoded/main.go`
2. **Dépendance manquante** : Exécuter `go mod tidy`
3. **Erreur de compilation** : Vérifier la version Go (1.21+ requise)
4. **RPC non accessible** : Vérifier le firewall et la configuration CORS

#### Logs et débogage

- Les logs sont affichés en stdout par défaut
- Utiliser l'indicateur `-verbose` pour des logs détaillés
- Les métriques RPC sont disponibles sur `/debug/metrics`

### Contribution

1. Fork le dépôt
2. Créer une branche (`git checkout -b feature/nouvelle-fonctionnalite`)
3. Committer vos changements (`git commit -m 'Ajout: nouvelle fonctionnalité'`)
4. Pusher vers la branche (`git push origin feature/nouvelle-fonctionnalite`)
5. Ouvrir une Pull Request

### Licence

Ce projet est sous licence MIT. Voir le fichier `LICENSE` pour plus de détails.

### Contact

- GitHub : https://github.com/martialzinsou/pulsechain-fork
- Issues : https://github.com/martialzinsou/pulsechain-fork/issues
- Auteur : Martial Zinsou