# PulseChain Fork - Documentation API JSON-RPC

> **Auteur :** Martial Zinsou  
> **Dépôt :** https://github.com/martialzinsou/pulsechain-fork  
> **Version :** 1.0.2  

## 1. Introduction à l'API JSON-RPC

Le nœud **PulseChain Fork** expose une interface **JSON-RPC 2.0** compatible avec les standards Ethereum. Cette API permet à tout client Web3 (MetaMask, Ethers.js, Web3.js, Foundry, Hardhat) d'interagir avec la blockchain via des requêtes HTTP POST.

### Point d'entrée

- **URL :** `http://localhost:8545` (ou `http://localhost:8545/rpc`)
- **Méthode HTTP :** `POST`
- **En-tête Content-Type :** `application/json`
- **Format payload :** JSON conformément à la spécification JSON-RPC 2.0

---

## 2. Endpoints RPC détaillés

### 2.1 `eth_chainId`

**Description :** Retourne l'identifiant du réseau au format hexadécimal.

**Requête :**
```http
POST http://localhost:8545
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "method": "eth_chainId",
  "params": [],
  "id": 1
}
```

**Réponse réussie :**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": "0x171"
}
```

**Notes :**
- `0x171` = 369 en décimal (ChainID PulseChain Mainnet)
- Utilisé par MetaMask pour détecter et ajouter le réseau

---

### 2.2 `web3_clientVersion`

**Description :** Retourne la version du client nœud.

**Requête :**
```http
POST http://localhost:8545
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "method": "web3_clientVersion",
  "params": [],
  "id": 1
}
```

**Réponse réussie :**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": "PulseChain-Fork/v1.0.0-MartialZinsou/darwin-amd64/go1.21"
}
```

**Notes :** Indique la version du logiciel nœud en cours d'exécution.

---

### 2.3 `eth_blockNumber`

**Description :** Retourne le nombre du bloc le plus récent.

**Requête :**
```http
POST http://localhost:8545
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "method": "eth_blockNumber",
  "params": [],
  "id": 1
}
```

**Réponse réussie (au démarrage) :**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": "0x0"
}
```

**Réponse réussie (après minage) :**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": "0x1"
}
```

**Notes :** Augmente de 1 à chaque nouveau bloc miné.

---

### 2.4 `eth_getBalance`

**Description :** Retourne le solde d'une adresse en Wei.

**Paramètres :**
1. `address` (DATA, 20 octets) - Adresse du compte
2. `quantityOrTag` (QUANTITY|TAG) - "latest", "earliest", "pending" ou nombre de bloc

**Requête :**
```http
POST http://localhost:8545
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "method": "eth_getBalance",
  "params": ["0x2b5AD5c4795c026514f8317c7a215E218DcCD6cF", "latest"],
  "id": 1
}
```

**Réponse réussie :**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": "0xd3c21bcecceda1000000"
}
```

**Valeur décimale :** 1 000 000 PLS (1 million de puls) - allocation initiale par défaut.

---

### 2.5 `eth_getBlockByNumber`

**Description :** Retourne les informations d'un bloc par son numéro.

**Paramètres :**
1. `tag` (STRING) - "latest", "pending", "earliest" ou nombre hexadécimal "0x0", etc.
2. `fullTxObjects` (BOOLEAN) - `true` pour objets transaction complets, `false` pour hashs uniquement

**Requête :**
```http
POST http://localhost:8545
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "method": "eth_getBlockByNumber",
  "params": ["0x0", true],
  "id": 1
}
```

**Réponse réussie (extrait) :**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "number": "0x0",
    "hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "miner": "0x2b5AD5c4795c026514f8317c7a215E218DcCD6cF",
    "gasLimit": "0x1c9c380",
    "gasUsed": "0x0",
    "timestamp": "0x0",
    "difficulty": "0x1",
    "transactionsRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
    "stateRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
    "transactions": []
  }
}
```

---

### 2.6 `net_version`

**Description :** Retourne l'ID réseau sous forme de chaîne de caractères.

**Requête :**
```http
POST http://localhost:8545
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "method": "net_version",
  "params": [],
  "id": 1
}
```

**Réponse réussie :**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": "369"
}
```

**Notes :** Identique à `eth_chainId` mais retourne une chaîne décimale au lieu de hexadécimal.

---

### 2.7 `eth_mining`

**Description :** Indique si le nœud est en cours de minage.

**Requête :**
```http
POST http://localhost:8545
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "method": "eth_mining",
  "params": [],
  "id": 1
}
```

**Réponse réussie :**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": false
}
```

**Notes :** `false` par défaut car le consensus PoA fonctionne en arrière-plan automatisé.

---

### 2.8 `eth_gasPrice`

**Description :** Retourne le prix du gaz de base en Wei.

**Requête :**
```http
POST http://localhost:8545
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "method": "eth_gasPrice",
  "params": [],
  "id": 1
}
```

**Réponse réussie :**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": "0x3b9aca00"
}
```

**Valeur décimale :** 1 000 000 000 Wei = **1 Gwei** (unité recommandée pour les transactions).

---

## 3. Codes de réponse d'erreur

| Code JSON-RPC | Message | Description |
|---------------|---------|-------------|
| `-32700` | Parse error | Payload JSON invalide ou illisible |
| `-32600` | Invalid Request | Requête non conforme au format JSON-RPC 2.0 |
| `-32601` | Method not found | Méthode demandée non implémentée sur ce nœud |
| `-32602` | Invalid params | Paramètres manquants, mauvais type, ou nombre incorrect |
| `-32603` | Internal error | Erreur lors de l'exécution interne du moteur |

### Exemple d'erreur :

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "error": {
    "code": -32602,
    "message": "Invalid params"
  }
}
```

---

## 4. Exemples cURL complets

### Exemple 1 : Vérifier le ChainID

```bash
curl -X POST http://localhost:8545 \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_chainId","params":[],"id":1}'
```

**Sortie attendue :**
```json
{"jsonrpc":"2.0","id":1,"result":"0x171"}
```

### Exemple 2 : Obtenir le solde

```bash
curl -X POST http://localhost:8545 \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_getBalance","params":["0x2b5AD5c4795c026514f8317c7a215E218DcCD6cF","latest"],"id":1}'
```

**Sortie attendue :**
```json
{"jsonrpc":"2.0","id":1,"result":"0xd3c21bcecceda1000000"}
```

### Exemple 3 : Obtenir les infos du bloc Genesis

```bash
curl -X POST http://localhost:8545 \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x0",true],"id":1}'
```

**Sortie attendue :** Objet JSON complet avec fields number, hash, miner, gasLimit, timestamp, difficulty, transactionsRoot, stateRoot, transactions.

---

## 5. Intégration avec MetaMask

Pour ajouter le réseau PulseChain Fork manuellement dans MetaMask :

1. Ouvrir MetaMask > "Ajouter un réseau"
2. Renseigner :
   - **Nom du réseau :** `PulseChain Fork (Martial Zinsou)`
   - **Nouvelle URL RPC :** `http://127.0.0.1:8545`
   - **ID de chaîne :** `369`
   - **Symbole de devise :** `PLS`
   - **URL de l'explorateur** : (optionnel) `http://localhost:8545`
3. Cliquer "Enregistrer"

MetaMask utilisera automatiquement `eth_chainId` pour valider la connexion et afficher le bon ID réseau.

---

## 6. Limites et considérations

| Limite | Détails |
|--------|---------|
| **Port RPC** | Par défaut 8545, configurable via `cmd/pulsenoded/main.go` |
| **Concurrentité** | Géré par mutex `sync.RWMutex` dans StateDB |
| **Taille bloc** | Gas limit par défaut : 30 000 000 |
| **Temps de bloc** | ~3 secondes (consensus PoA) |
| **CORS** | En-têtes `Access-Control-Allow-Origin: *` pour compatibilité navigateur |

---

## 7. Dépannage des requêtes RPC

### Problème courant : "Connection refused"

**Cause :** Le nœud n'est pas en cours d'exécution ou le port est bloqué.

**Solution :**
```bash
# Vérifier que le nœud tourne
./pulsenoded &

# Ou via Docker
docker compose up -d

# Vérifier le port
lsof -i :8545
```

### Problème courant : "Method not found"

**Cause :** Le nœud ne supporte pas cette méthode ou la version JSON-RPC est incorrecte.

**Solution :**
- Mettre à jour vers la dernière version du nœud
- Vérifier la console d'erreurs (`docker logs` ou `go run ./cmd/pulsenoded`)

### Problème courant : "Invalid params"

**Cause :** Format de requête JSON incorrect ou paramètres manquants.

**Solution :**
- Respecter l'ordre et les types de paramètres définis dans la documentation
- Utiliser les exemples cURL fournis dans cette documentation

---

**Auteur :** Martial Zinsou  
**Version :** 1.0.2  
**Dépôt :** https://github.com/martialzinsou/pulsechain-fork  
**Licence :** MIT