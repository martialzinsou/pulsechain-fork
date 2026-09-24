# PulseChain Fork - Captures d'écran

> **Auteur :** Martial Zinsou  
> **Dépôt :** https://github.com/martialzinsou/pulsechain-fork  

## 1. Interface RPC JSON-RPC (port 8545)

### Capture 1 : Requête `eth_chainId`

**URL :** `http://localhost:8545`

**Requête envoyée :**
```http
POST / HTTP/1.1
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "method": "eth_chainId",
  "params": [],
  "id": 1
}
```

**Réponse affichée :**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": "0x171"
}
```

**Légende :** La chaîne PulseChain Fork utilise le ChainID `0x171` (équivalent décimal **369**). Cette valeur est essentielle pour configurer MetaMask et autres clients Web3.

---

### Capture 2 : Requête `eth_blockNumber`

**URL :** `http://localhost:8545`

**Requête envoyée :**
```http
POST / HTTP/1.1
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "method": "eth_blockNumber",
  "params": [],
  "id": 1
}
```

**Réponse affichée :**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": "0x0"
}
```

**Légende :** Au démarrage (bloc Genesis), le résultat est `0x0`. Après minage de blocs, cette valeur incrémente pour afficher le numéro du dernier bloc.

---

### Capture 3 : Requête `eth_getBalance`

**URL :** `http://localhost:8545`

**Requête envoyée :**
```http
POST / HTTP/1.1
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "method": "eth_getBalance",
  "params": ["0x2b5AD5c4795c026514f8317c7a215E218DcCD6cF", "latest"],
  "id": 1
}
```

**Réponse affichée :**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": "0xd3c21bcecceda1000000"
}
```

**Légende :** Solde du compte Genesis PulseChain en Wei. La valeur `0xd3c21bcecceda1000000` représente **1 000 000 PLS** (1 million de puls), allocation initiale par défaut.

---

### Capture 4 : Requête `eth_getBlockByNumber`

**URL :** `http://localhost:8545`

**Requête envoyée :**
```http
POST / HTTP/1.1
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "method": "eth_getBlockByNumber",
  "params": ["0x0", true],
  "id": 1
}
```

**Réponse affichée :** (extrait)

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

**Légende :** Données complètes du bloc Genesis. Les champs clés incluent :
- `miner` : adresse du validateur/coinbase
- `gasLimit` : limite de gaz par bloc (30 000 000)
- `timestamp` : horodatage de création
- `difficulty` : difficulté initiale (1)

---

## 2. Console de démarrage du nœud

### Capture 5 : Sortie console au lancement

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

**Légende :** Fenêtre terminal affichée juste après exécution de `go run ./cmd/pulsenoded` ou `./pulsenoded`. Les éléments en vert `[+]` indiquent les étapes d'initialisation réussies.

---

### Capture 6 : Onglet "Network" de MetaMask

**Légende :** Capture d'écran de l'interface MetaMask montrant :
- Le réseau "PulseChain Fork (Martial Zinsou)" sélectionné
- L'URL RPC `http://127.0.0.1:8545` affichée
- L'ID de chaîne `369` confirmé
- Le solde PLS du compte associé affiché

---

## 3. Schémas et Diagrammes

### Captures de diagrammes Mermaid

#### Diagramme d'architecture globale

![Architecture PulseChain Fork](architecture-diagram.png)

*Légende : Vue globale à 4 couches du système PulseChain Fork avec séparation claire entre l'interface utilisateur, le serveur RPC, le moteur blockchain et le consensus PoA.*

#### Diagramme de classes

![Diagramme de classes](class-diagram.png)

*Légende : Diagramme UML des classes principales : Blockchain, Block, Transaction, StateDB, et Consensus avec leurs attributs et méthodes clés.*

#### Diagramme de séquence - Initialisation

![Sequence diagram](sequence-diagram.png)

*Légende : Flux d'initialisation du bloc Genesis depuis l'exécution Go jusqu'à la préparation de l'état.*

---

## 4. Outils de capture recommandés

### Navigateurs (pour captures RPC)

| Navigateur | Méthode | Raccourci |
|------------|---------|-----------|
| **Chrome / Edge** | OutilsDéveloppeur → Application → JSON | `F12` then `Ctrl+Shift+I` |
| **Firefox** | Inspecteur → Console → Réseau | `F12` |
| **Safari** | Préférences → Avancé → Show Develop menu | `Cmd+Option+I` |

### Capture d'écran système

| Système d'exploitation | Raccourci | Outil intégré |
|------------------------|-----------|---------------|
| **Windows** | `Impression écran` / `Win+Shift+S` | Snipping Tool |
| **macOS** | `Cmd+Shift+4` | Capture d'écran native |
| **Linux (GNOME)** | `Shift+PrtSc` | `gnome-screenshot` |
| **Linux (X11)** | `PrtSc` | `scrot`, `xwd` |

### Outils de développement recommandés

- **JSON Viewer** (extension navigateur) : Formate et colore les réponses JSON pour une lecture facilitée
- **React Developer Tools** / **Vue.js DevTools** : Si utilisation d'interfaces Web personnalisées
- **Postman** : Pour tester et enregistrer des collections d'endpoints RPC

---

## 5. Intégration dans la documentation

### Où placer les captures dans ce wiki :

1. **Section "Interface RPC"** : Après chaque exemple de requête, insérer la capture correspondante
2. **Section "Console de démarrage"** : Dans le guide d'installation, après la commande de lancement
3. **Section "MetaMask"** : Dans le chapitre "Connexion avec MetaMask", après la configuration du réseau
4. **Sections UML** : Captures d'écran des prévisualisations Mermaid dans les éditeurs supports (GitHub, GitLab, MkDocs)

### Format recommandé :

- **Format d'image :** PNG (pour les captures d'interface et textes) ou JPEG (pour captures générales)
- **Dimensions :** Largeur maximale 800px pour une affichage optimal dans Markdown
- **Légende :** Toujours ajouter une légende Markdown `![Légende](chemin/vers/image)` en dessous de chaque capture
- **Nom de fichier :** Utiliser des noms descriptifs (ex: `rpc-eth-chainid.png`, `metamask-network.png`)

---

**Auteur :** Martial Zinsou  
**Version :** 1.0.2  
**Dépôt :** https://github.com/martialzinsou/pulsechain-fork