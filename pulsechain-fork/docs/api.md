# Documentation API JSON-RPC - PulseChain Fork

> **Auteur et Développeur :** Martial Zinsou  
> **Dépôt officiel :** https://github.com/martialzinsou/pulsechain-fork  
> **Spécification :** JSON-RPC 2.0 (Compatible Ethereum & PulseChain)  
> **Licence :** MIT (c) 2026 Martial Zinsou  

---

## Vue d'ensemble

Le nœud PulseChain Fork conçu par **Martial Zinsou** expose une interface JSON-RPC 2.0 sur le port `8545`. Cette interface permet à tout client Web3 standard (MetaMask, Ethers.js, Web3.js, Foundry, Hardhat) d'interagir nativement avec la blockchain.

### Endpoint
- **URL standard :** `http://localhost:8545`
- **Path alternatif :** `http://localhost:8545/rpc`
- **Méthode HTTP :** `POST`
- **Headers :** `Content-Type: application/json`

---

## Méthodes Supportées

### 1. `eth_chainId`
Retourne l'identifiant unique du réseau (Chain ID) au format hexadécimal.

- **Requête :**
```json
{
  "jsonrpc": "2.0",
  "method": "eth_chainId",
  "params": [],
  "id": 1
}
```

- **Réponse :**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": "0x171" // 369 en décimal (PulseChain)
}
```

---

### 2. `web3_clientVersion`
Retourne la version du client du nœud blockchain.

- **Requête :**
```json
{
  "jsonrpc": "2.0",
  "method": "web3_clientVersion",
  "params": [],
  "id": 1
}
```

- **Réponse :**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": "PulseChain-Fork/v1.0.0-MartialZinsou/darwin-amd64/go1.21"
}
```

---

### 3. `eth_blockNumber`
Retourne le numéro du bloc le plus récent dans la chaîne.

- **Requête :**
```json
{
  "jsonrpc": "2.0",
  "method": "eth_blockNumber",
  "params": [],
  "id": 1
}
```

- **Réponse :**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": "0x0"
}
```

---

### 4. `eth_getBalance`
Retourne le solde en Wei (PLS) de l'adresse spécifiée.

- **Paramètres :**
  1. `Adresse` (DATA, 20 octets)
  2. `Bloc tag` ("latest", "earliest", "pending" ou numéro de bloc hex)

- **Requête :**
```json
{
  "jsonrpc": "2.0",
  "method": "eth_getBalance",
  "params": [
    "0x2b5AD5c4795c026514f8317c7a215E218DcCD6cF",
    "latest"
  ],
  "id": 1
}
```

- **Réponse :**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": "0xd3c21bcecceda1000000"
}
```

---

### 5. `eth_getBlockByNumber`
Retourne les données complètes d'un bloc selon son numéro ou son alias.

- **Paramètres :**
  1. `Numéro de bloc` ("0x0", "latest", etc.)
  2. `Transactions détaillées` (booléen : `true` pour les objets complets, `false` pour les hashs)

- **Requête :**
```json
{
  "jsonrpc": "2.0",
  "method": "eth_getBlockByNumber",
  "params": ["0x0", true],
  "id": 1
}
```

- **Réponse :**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "number": "0x0",
    "hash": "0x...",
    "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "timestamp": "0x64604500",
    "miner": "0x2b5AD5c4795c026514f8317c7a215E218DcCD6cF",
    "gasLimit": "0x1c9c380",
    "gasUsed": "0x0",
    "difficulty": "0x1",
    "transactionsRoot": "0x...",
    "stateRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
    "transactions": []
  }
}
```

---

### 6. `net_version`
Retourne l'ID réseau sous forme de chaîne décimale.

- **Réponse :** `"369"`

---

## Codes d'Erreurs JSON-RPC

| Code | Message | Description |
|------|---------|-------------|
| `-32700` | Parse error | Payload JSON invalide ou illisible |
| `-32600` | Invalid Request | Requête non conforme au format JSON-RPC 2.0 |
| `-32601` | Method not found | Méthode demandée non implémentée sur ce nœud |
| `-32602` | Invalid params | Paramètres manquants ou type incorrect |
| `-32603` | Internal error | Erreur d'exécution interne du moteur |

---

**Auteur du projet :** Martial Zinsou  
**Dépôt :** https://github.com/martialzinsou/pulsechain-fork  

---

## Captures d'écran des Endpoints RPC

### Représentation visuelle des réponses JSON-RPC :

**1. Exemple de réponse `eth_chainId` :**
```
HTTP POST http://localhost:8545
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "method": "eth_chainId",
  "params": [],
  "id": 1
}
```
→ **Affichage résultat** : `0x171` (représente la chaîne PulseChain Mainnet, ChainID 369)

**2. Exemple de réponse `eth_blockNumber` :**
```
HTTP POST http://localhost:8545
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "method": "eth_blockNumber",
  "params": [],
  "id": 1
}
```
→ **Affichage résultat** : `0x0` (premier bloc / bloc Genesis)

**3. Exemple de réponse `eth_getBalance` :**
```
HTTP POST http://localhost:8545
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "method": "eth_getBalance",
  "params": ["0x2b5AD5c4795c026514f8317c7a215E218DcCD6cF", "latest"],
  "id": 1
}
```
→ **Affichage résultat** : `0xd3c21bcecceda1000000` (solde en Wei du compte Genesis)

**4. Exemple de réponse `eth_getBlockByNumber` :**
```
HTTP POST http://localhost:8545
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "method": "eth_getBlockByNumber",
  "params": ["0x0", true],
  "id": 1
}
```
→ **Affichage résultat** : Objet JSON complet avec number, hash, miner, gaz, etc.

### Guidelines pour les captures d'écran documentation :

- **Format recommandé** : PNG (qualité élevée, texte lisible)
- **Zone de capture** : Se concentrer sur la partie JSON de la réponse dans l'outil de développement
- **Légende** : Toujours inclure la méthode RPC et un exemple de paramètre
- **Lisibilité** : Zoomer si nécessaire pour afficher les adresses hexadécimales complètes
- **Outils** : Utiliser les "Developer Tools" du navigateur (F12), section "Response" ou "Pretty Print"

---
