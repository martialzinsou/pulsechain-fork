# 🚀 PulseChain Fork - Nœud Blockchain Complet

[![Author](https://img.shields.io/badge/Auteur-Martial%20Zinsou-blue.svg)](https://github.com/martialzinsou)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://golang.org)
[![ChainID](https://img.shields.io/badge/ChainID-369%20%7C%20943-purple.svg)](https://pulsechain.com)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?logo=docker)](https://docker.com)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

Un projet complet et autonome d'implémentation et de fork de la blockchain **PulseChain** en langage **Go**, développé et documenté par **Martial Zinsou**.

Ce dépôt fournit l'ensemble des modules nécessaires : moteur de chaîne, consensus Proof-of-Authority (PoA) avec temps de bloc de 3 secondes, registre des comptes d'état (StateDB), serveur JSON-RPC 2.0 compatible Ethereum/MetaMask, fichiers de bloc genesis, scripts d'initialisation et configuration Docker.

---

## 👤 Auteur & Créateur du Projet

- **Auteur principal :** Martial Zinsou
- **Dépôt officiel :** [https://github.com/martialzinsou/pulsechain-fork](https://github.com/martialzinsou/pulsechain-fork)
- **Profil GitHub :** [@martialzinsou](https://github.com/martialzinsou)

---

## 📑 Sommaire

1. [Fonctionnalités Principales](#-fonctionnalités-principales)
2. [Structure Complète du Projet](#-structure-complète-du-projet)
3. [Guide d'Installation et Démarrage Rapide](#-guide-dinstallation-et-démarrage-rapide)
4. [Configuration du Bloc Genesis](#-configuration-du-bloc-genesis)
5. [Interface JSON-RPC 2.0 (Compatible Web3)](#-interface-json-rpc-20-compatible-web3)
6. [Connexion avec MetaMask](#-connexion-avec-metamask)
7. [Exécution Conteneurisée (Docker)](#-exécution-conteneurisée-docker)
8. [Scripts d'Exploitation & Tests](#-scripts-dexploitation--tests)
9. [Architecture Technique](#-architecture-technique)
10. [Documentation Complète](#-documentation-complète)

---

## ⚡ Fonctionnalités Principales

- **Moteur Blockchain Go :** Gestion de chaîne de blocs, hachage cryptographique SHA-256 / RLP, calcul de racines de transactions et validation d'état.
- **Consensus PulseChain PoA :** Algorithme Proof-of-Authority avec période de bloc rapide (~3 secondes) et vérification temporelle.
- **Registre StateDB :** Gestion des soldes, transferts de pièces (PLS), gestion des nonces et protection anti-rejeu.
- **Serveur JSON-RPC 2.0 :** Compatibilité totale avec les librairies Ethereum (`web3.js`, `ethers.js`, `Foundry`, `Hardhat`) et les portefeuilles comme MetaMask.
- **Spécification Genesis :** Fichier `genesis.json` officiel et structure `genesis.go` avec ChainID `369` (Mainnet) et `943` (Testnet-v4).
- **Prêt pour Docker :** Déploiement multi-étapes optimisé via `Dockerfile` et `docker-compose.yml`.

---

## 📁 Structure Complète du Projet

```text
pulsechain-fork/
├── Dockerfile                         # Image Docker multi-stage pour le nœud
├── docker-compose.yml                 # Environnement conteneurisé
├── go.mod                             # Dépendances du module Go
├── scripts/
│   ├── init-node.sh                   # Script d'initialisation et démarrage
│   └── send-tx.sh                     # Script de validation et test RPC
├── pulsechain-fork/
│   ├── cmd/
│   │   └── pulsenoded/
│   │       └── main.go                # Point d'entrée exécutable du démon
│   ├── genesis/
│   │   ├── genesis.go                 # Modèle de configuration Go
│   │   └── genesis.json               # Spécification standard du Genesis block
│   ├── internal/
│   │   ├── blockchain/
│   │   │   ├── block.go               # Structures Block & Transaction
│   │   │   └── chain.go               # Moteur de consensus et validation
│   │   ├── consensus/
│   │   │   └── proof.go               # Algorithme de consensus Proof-of-Authority
│   │   ├── ledger/
│   │   │   └── account.go             # StateDB, comptes et gestion des soldes
│   │   └── rpc/
│   │       └── server.go              # Serveur HTTP JSON-RPC 2.0
│   └── docs/
│       ├── architecture.md            # Spécifications de l'architecture
│       ├── api.md                     # Documentation complète des endpoints RPC
│       └── guide_utilisateur.md       # Manuel utilisateur exhaustif pas-à-pas
└── README.md                          # Documentation générale du projet
```

---

## 🚀 Guide d'Installation et Démarrage Rapide

### 1. Prérequis
- **Go 1.21+** (ou Docker si vous préférez exécuter en conteneur)
- **Git**
- **cURL** pour les requêtes de test

### 2. Téléchargement et compilation

```bash
# Cloner le projet
git clone https://github.com/martialzinsou/pulsechain-fork.git
cd pulsechain-fork

# Préparer les dépendances
go mod tidy

# Compiler le binaire
go build -o pulsenoded ./pulsechain-fork/cmd/pulsenoded
```

### 3. Lancement du Nœud

```bash
# Exécution directe
./pulsenoded

# Ou via le script automatisé
./scripts/init-node.sh
```

**Sortie console attendue :**
```text
======================================================
     PULSECHAIN FORK NODE - MARTIAL ZINSOU            
======================================================
[+] Initialisation du bloc génèse PulseChain...
[+] ChainID: 369 | Coinbase: 0x2b5AD5c4795c026514f8317c7a215E218DcCD6cF
[+] Serveur JSON-RPC prêt à recevoir les connexions
[+] RPC Endpoint: http://localhost:8545 (compatible MetaMask/Ethers)
[+] Blocs initialisés: 1 (Bloc Genesis actif)
[+] Démarrage de l'écoute sur le port :8545 ...
```

---

## ⚙️ Configuration du Bloc Genesis

Le fichier `pulsechain-fork/genesis/genesis.json` configure les paramètres initiaux du réseau fork :

```json
{
  "config": {
    "chainId": 369,
    "homesteadBlock": 0,
    "eip155Block": 0,
    "eip158Block": 0,
    "byzantiumBlock": 0,
    "constantinopleBlock": 0,
    "petersburgBlock": 0,
    "istanbulBlock": 0,
    "berlinBlock": 0,
    "londonBlock": 0,
    "clique": {
      "period": 3,
      "epoch": 30000
    }
  },
  "difficulty": "1",
  "gasLimit": "30000000",
  "alloc": {
    "0x2b5AD5c4795c026514f8317c7a215E218DcCD6cF": {
      "balance": "1000000000000000000000000000"
    }
  }
}
```

---

## 🌐 Interface JSON-RPC 2.0 (Compatible Web3)

Le serveur RPC écoute par défaut sur le port `8545`.

### Principaux Endpoints :

| Méthode | Description | Exemple cURL |
|---|---|---|
| `eth_chainId` | Retourne l'identifiant du réseau (369 = 0x171) | `curl -X POST http://localhost:8545 -d '{"jsonrpc":"2.0","method":"eth_chainId","params":[],"id":1}'` |
| `web3_clientVersion` | Identifiant du client blockchain | `curl -X POST http://localhost:8545 -d '{"jsonrpc":"2.0","method":"web3_clientVersion","params":[],"id":1}'` |
| `eth_blockNumber` | Numéro du dernier bloc miné | `curl -X POST http://localhost:8545 -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'` |
| `eth_getBalance` | Solde du compte en Wei (PLS) | `curl -X POST http://localhost:8545 -d '{"jsonrpc":"2.0","method":"eth_getBalance","params":["0x2b5AD5c4795c026514f8317c7a215E218DcCD6cF","latest"],"id":1}'` |
| `eth_getBlockByNumber` | Données complètes d'un bloc | `curl -X POST http://localhost:8545 -d '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x0",true],"id":1}'` |

Consultez la [Documentation API Complète](pulsechain-fork/docs/api.md) pour les détails et formats de réponse.

---

## 🦊 Connexion avec MetaMask

1. Ouvrez votre extension **MetaMask**.
2. Allez dans **Paramètres** > **Réseaux** > **Ajouter un réseau**.
3. Renseignez :
   - **Nom du réseau :** `PulseChain Fork (Martial Zinsou)`
   - **URL de RPC :** `http://127.0.0.1:8545`
   - **ID de chaîne :** `369`
   - **Symbole de devise :** `PLS`
4. Validez pour utiliser votre nœud local immédiatement avec vos contrats et transactions !

---

## 🐳 Exécution Conteneurisée (Docker)

Pour lancer le nœud sans installer Go :

```bash
# Construire et démarrer le conteneur
docker compose up -d

# Suivre les journaux d'exécution
docker compose logs -f

# Arrêter le conteneur
docker compose down
```

---

## 🧪 Scripts d'Exploitation & Tests

Un script de test automatique est mis à disposition dans `scripts/` :

```bash
# Vérifier tous les endpoints du nœud
./scripts/send-tx.sh
```

---

## 🏗️ Architecture Technique

Pour comprendre en profondeur les choix de conception :
- 📄 [Architecture Système](pulsechain-fork/docs/architecture.md) : Modèle du moteur, consensus, persistance.
- 📄 [Documentation API](pulsechain-fork/docs/api.md) : Spécification complète des requêtes RPC.
- 📄 [Guide Utilisateur](pulsechain-fork/docs/guide_utilisateur.md) : Guide détaillé pour débutants et développeurs.

---

## 📄 Licence

Ce projet est sous licence **MIT**. Vous êtes libre de l'utiliser, de le modifier et de le redistribuer.

**Auteur :** [Martial Zinsou](https://github.com/martialzinsou)  
**GitHub :** [https://github.com/martialzinsou/pulsechain-fork](https://github.com/martialzinsou/pulsechain-fork)
