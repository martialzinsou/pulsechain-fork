# Architecture Technique - PulseChain Fork

> **Auteur et Architecte :** Martial Zinsou  
> **Dépôt GitHub :** https://github.com/martialzinsou/pulsechain-fork  
> **Version :** 1.0.0  
> **Licence :** MIT (c) 2026 Martial Zinsou  

---

## 1. Vue d'ensemble du Système

Ce document décrit l'architecture complète du projet **PulseChain Fork**, conçu et développé par **Martial Zinsou**. Cette implémentation blockchain modulaire en Go est totalement interopérable avec l'écosystème Ethereum et PulseChain.

```text
+-------------------------------------------------------+
|          Clients Web3 (MetaMask, Ethers.js)           |
+-------------------------------------------------------+
                           | JSON-RPC (HTTP :8545)
                           v
+-------------------------------------------------------+
|        Serveur RPC (internal/rpc/server.go)           |
|                 Auteur: Martial Zinsou                |
+-------------------------------------------------------+
                           |
            +--------------+--------------+
            |                             |
            v                             v
+-----------------------+     +-------------------------+
|    Moteur Blockchain   |     |   Registre StateDB      |
| (internal/blockchain) |<--->|   (internal/ledger)     |
+-----------------------+     +-------------------------+
            |                             |
            +--------------+--------------+
                           |
                           v
+-------------------------------------------------------+
|       Consensus PoA (internal/consensus/proof.go)     |
|          Période 3s | Validation des blocs            |
+-------------------------------------------------------+
```

---

## 2. Structure des Composants et Responsabilités

### 2.1. Moteur de Blockchain (`internal/blockchain/`)
*Développé par Martial Zinsou.*

- **Structures Fondamentales (`block.go`) :**
  - Définition des blocs (`Block`) et transactions signées (`Transaction`).
  - Hachage cryptographique SHA-256 avec sérialisation déterministe.
  - Arbre de transactions (`TxsHash`) et racine d'état (`StateRoot`).
  - Fonction d'initialisation du bloc Genesis (`GenesisBlock`).

- **Gestionnaire de Chaîne (`chain.go`) :**
  - Validation séquentielle des blocs (`ParentHash`, index chronologique).
  - Intégration du consensus pour approbation des nouveaux blocs.
  - Exécution des transactions et mise à jour d'état atomique.
  - Fonction de minage de blocs (`MineBlock`) avec récompenses coinbase.

---

### 2.2. Algorithme de Consensus PoA (`internal/consensus/`)
*Développé par Martial Zinsou.*

- **Proof-of-Authority (`proof.go`) :**
  - Cadence de blocs rapide optimisée à **3 secondes** (calquée sur PulseChain).
  - Gestion de la liste des autorités et validateurs autorisés.
  - Ajustement dynamique de la difficulté et vérification de la dérive temporelle.
  - Algorithme de validation par signature et nonce.

---

### 2.3. Registre d'État & Ledger (`internal/ledger/`)
*Développé par Martial Zinsou.*

- **Gestionnaire `StateDB` (`account.go`) :**
  - Modèle de comptes avec `Address`, `Balance` en PLS (précision 18 décimales / `*big.Int`), `Nonce`, `CodeHash`.
  - Protection contre la concurrence via verrous lecture/écriture (`sync.RWMutex`).
  - Mécanisme anti-rejeu et débits sécurisés avec contrôle de solde préalable.
  - Calcul déterministe de la racine d'état (`stateRoot`).

---

### 2.4. Serveur JSON-RPC 2.0 (`internal/rpc/`)
*Développé par Martial Zinsou.*

- **Passerelle Web3 (`server.go`) :**
  - Support natif des requêtes JSON-RPC 2.0 sur `http://localhost:8545`.
  - En-têtes CORS universels pour intégration directe avec les navigateurs et DApps.
  - Implémentation des standards Ethereum : `eth_chainId`, `eth_blockNumber`, `eth_getBalance`, `eth_getBlockByNumber`, `web3_clientVersion`.

---

## 3. Paramètres du Bloc Genesis

Le réseau initialisé par **Martial Zinsou** repose sur la spécification PulseChain suivante :
- **Chain ID :** `369` (PulseChain Mainnet) / `943` (Testnet-v4)
- **Coinbase initial :** `0x2b5AD5c4795c026514f8317c7a215E218DcCD6cF`
- **Gas Limit :** `30 000 000`
- **Période Clique PoA :** `3 secondes`

---

## 4. Propriété Intellectuelle et Licence

Ce projet est la création originale de **Martial Zinsou**.  
Distribué sous licence MIT. Toute utilisation ou réutilisation doit mentionner l'auteur d'origine.

**Contact & Profil :** [https://github.com/martialzinsou](https://github.com/martialzinsou)
