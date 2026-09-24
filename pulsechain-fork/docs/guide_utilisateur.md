# Guide Utilisateur Complet - PulseChain Fork

> **Auteur :** Martial Zinsou  
> **Projet :** Fork de la Blockchain PulseChain  
> **Dépôt :** https://github.com/martialzinsou/pulsechain-fork  

---

## Sommaire

1. [Introduction](#1-introduction)
2. [Prérequis Système](#2-prérequis-système)
3. [Installation Pas-à-Pas](#3-installation-pas-à-pas)
4. [Démarrage du Nœud](#4-démarrage-du-nœud)
5. [Configuration du Réseau et du Fork](#5-configuration-du-réseau-et-du-fork)
6. [Connexion avec MetaMask et Outils Web3](#6-connexion-avec-metamask-et-outils-web3)
7. [Scripts Utilitaires](#7-scripts-utilitaires)
8. [Déploiement avec Docker](#8-déploiement-avec-docker)
9. [Foire Aux Questions (FAQ)](#9-foire-aux-questions-faq)

---

## 1. Introduction

PulseChain Fork est une implémentation modulaire en langage Go conçue pour recréer l'environnement de la blockchain PulseChain (fork d'Ethereum à haute cadence avec temps de bloc de ~3 secondes et frais de gaz réduits). 

Ce guide vous accompagne pas-à-pas pour initialiser, exécuter, administrer et interagir avec votre nœud blockchain.

---

## 2. Prérequis Système

| Composant | Version minimale recommandée |
|-----------|------------------------------|
| **Système d'exploitation** | Linux (Ubuntu 20.04+), macOS (11+) ou Windows (WSL2) |
| **Go** | 1.21+ (si compilation native) |
| **Docker & Docker Compose** | 20.10+ (si exécution conteneurisée) |
| **cURL & Python3** | Pour exécuter les scripts de test RPC |
| **Mémoire RAM** | 4 Go minimum (8 Go recommandés) |

---

## 3. Installation Pas-à-Pas

### Étape 3.1 : Clonage du dépôt GitHub
```bash
git clone https://github.com/martialzinsou/pulsechain-fork.git
cd pulsechain-fork
```

### Étape 3.2 : Téléchargement des dépendances Go
```bash
go mod tidy
```

### Étape 3.3 : Compilation du binaire
```bash
go build -o pulsenoded ./pulsechain-fork/cmd/pulsenoded
```

---

## 4. Démarrage du Nœud

### Option A : Démarrage direct avec le script d'initialisation
```bash
./scripts/init-node.sh
```

### Option B : Démarrage en ligne de commande Go
```bash
go run ./pulsechain-fork/cmd/pulsenoded
```

### Sortie attendue dans la console :
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

## 5. Configuration du Réseau et du Fork

### Fichier `genesis.json`
Le fichier `pulsechain-fork/genesis/genesis.json` paramètre l'état initial :
- **ChainID :** `369` (Mainnet) ou `943` (Testnet-v4)
- **Consensus Period :** `3` secondes
- **Gas Limit par bloc :** `30 000 000`
- **Allocation initiale :** Répartition des soldes natifs (PLS) pour les validateurs et comptes du fork

Pour modifier le Chain ID ou l'adresse du validateur, éditez `genesis.json` avant le démarrage.

---

## 6. Connexion avec MetaMask et Outils Web3

### Ajouter le Réseau dans MetaMask :

1. Ouvrez MetaMask > Menu Réseaux > **Ajouter un réseau manuellement**
2. Remplissez les champs :
   - **Nom du réseau :** `PulseChain Fork (Martial Zinsou)`
   - **Nouvelle URL RPC :** `http://127.0.0.1:8545`
   - **ID de chaîne :** `369` (ou `0x171`)
   - **Symbole de devise :** `PLS`
   - **URL de l'explorateur :** `http://localhost:8545` (optionnel)
3. Cliquez sur **Enregistrer**. Votre portefeuille est désormais connecté à votre nœud local !

---

## 7. Scripts Utilitaires

Le dossier `scripts/` contient des outils prêts à l'emploi :

### Tester les endpoints RPC :
```bash
./scripts/send-tx.sh
```
Ce script vérifie automatiquement :
- La connectivité réseau
- Le numéro de version du client
- Le dernier bloc
- Le solde du compte Genesis

---

## 8. Déploiement avec Docker

Si vous ne souhaitez pas installer Go localement :

```bash
# Construction et lancement en tâche de fond
docker compose up -d

# Visualisation des journaux en direct
docker compose logs -f

# Arrêt du nœud
docker compose down
```

---

## 8. Captures d'écran et Documentation Visuelle

### Comment capturer des captures d'écran du nœud PulseChain Fork :

**1. Interface RPC (Port 8545) :**
- Ouvrez votre navigateur à l'adresse `http://localhost:8545`
- Utilisez l'outil de développement (F12) ou l'extension "JSON Viewer" pour formater les réponses
- Captures recommandées :
  - La réponse de `eth_chainId` montrant `0x171` (369)
  - La réponse de `eth_blockNumber` montrant le dernier bloc
  - La réponse de `eth_getBalance` montrant le solde du compte Genesis

**2. Console de démarrage du nœud :**
- Lancez le nœud avec `go run ./cmd/pulsenoded` ou `./pulsenoded`
- Les captures suivantes sont utiles pour la documentation :
  - La bannière `PULSECHAIN FORK NODE - MARTIAL ZINSOU`
  - Les logs d'initialisation du bloc génèse
  - Le message de prêt du serveur RPC sur le port 8545

**3. Interface MetaMask :**
- Après avoir ajouté le réseau PulseChain Fork (ChainID 369)
- Capture de la sélection du réseau dans le menu déroulant de MetaMask
- Capture du solde initial affiché pour le compteGenesis

### Outils recommandés :
- **Navigateur** : Chrome, Firefox, Edge (outils de développement F12)
- **Extensions** : JSON Viewer, MetaMask interface
- **Outils système** : PrtScn (Windows), Cmd+Shift+4 (Mac), `gnome-screenshot` ou `scrot` (Linux)

---

## 9. Foire Aux Questions (FAQ)

### Q : Le port 8545 est déjà occupé ?
> Modifiez le port dans `cmd/pulsenoded/main.go` ou configurez la variable `RPC_PORT` dans `docker-compose.yml`.

### Q : Comment réinitialiser la blockchain à zéro ?
> Supprimez le dossier de données `./data` et relancez le nœud pour recréer le bloc Genesis.

---

**Auteur :** Martial Zinsou  
**Licence :** MIT  
**Support :** Ouvrez une issue sur https://github.com/martialzinsou/pulsechain-fork/issues
