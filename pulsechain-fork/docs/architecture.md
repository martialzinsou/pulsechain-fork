# Architecture Documentation

## Auteur : Martial Zinsou

### Vue d'ensemble

Ce document décrit l'architecture du projet PulseChain Fork, une implémentation blockchain en Go compatible avec l'écosystème Ethereum.

### Structure des composants

#### 1. Moteur de blockchain (`internal/blockchain/`)

Le moteur de blockchain est le cœur du système, responsable de :

- **Structure des blocs** : définition du format Block et Transaction
- **Validation** : vérification de la validité des blocs via le consensus
- **Gestion de chaîne** : stockage et récupération de la chaîne de blocs
- **État du réseau** : gestion des comptes, soldes et stockage

Composants clés :
- `block.go` : Structures Block et Transaction, méthodes de hachage
- `chain.go` : Gestion de la chaîne, ajout de blocs, requêtes d'état

#### 2. Consensus (`internal/consensus/`)

L'algorithme de consensus suit le modèle PulseChain Proof-of-Authority :

- **Validation de blocs** : vérification de la preuve de travail/d'autorité
- **Ajustement de difficulté** : calcul de la cible de hachage
- **Paramètres de réseau** : temps de bloc, difficulté cible

Fichier principal : `proof.go` - Implémentation PoA avec calcul de difficulté

#### 3. Ledger (`internal/ledger/`)

Gestion du registre des comptes et des soldes :

- **StateDB** : base de données d'état avec comptes, nonces et stockage
- **Account management** : création, mise à jour, consultation de comptes
- **Solde des tokens** : gestion ETH et tokens PRC-20

Fichier principal : `account.go` - Gestion des comptes utilisateur

#### 4. Serveur RPC (`internal/rpc/`)

Interface JSON-RPC compatible Ethereum :

- **Points de terminaison** : méthodes eth_* standards
- **Format des réponses** : JSON-RPC 2.0
- **Compatibilité** : geth, metrics, autres clients Ethereum

Fichier principal : `server.go` - Exposition des méthodes RPC

### Flux de données

1. **Nouveau bloc** : reçu → validation consensus → mise à jour état → ajout chaîne
2. **Requête RPC** : entrante → traitement → réponse JSON
3. **Transaction** : reçue → inclusion dans bloc → validation → état mis à jour

### Points d'extension

- Ajouter de nouveaux hard forks dans la configuration genesis
- Étendre les méthodes RPC dans internal/rpc/server.go
- Implémenter de nouveaux algorithmes de consensus
- Ajouter des fonctionnalités de couche de réseau P2P