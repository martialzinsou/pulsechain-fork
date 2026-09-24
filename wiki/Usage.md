# PulseChain Fork - Guide d'Utilisation

> **Auteur :** Martial Zinsou  
> **Dépôt :** https://github.com/martialzinsou/pulsechain-fork  
> **Version :** 1.0.2  

## 1. Prérequis système

Avant d'installer ou d'exécuter PulseChain Fork, assurez-vous d'avoir les composants suivants :

| Composant | Version minimale | Notes |
|-----------|-----------------|-------|
| **Système d'exploitation** | Linux (Ubuntu 20.04+), macOS (11+), Windows (WSL2) | Testé sur les trois principales plateformes |
| **Go (Golang)** | 1.21+ | Requis pour compilation native |
| **Docker** | 20.10+ | Requis pour déploiement en conteneur |
| **Docker Compose** | 2.20+ | Optionnel mais recommandé |
| **cURL** | Mis en bundle | Pour tester les endpoints RPC |
| **Python3** | 3.8+ | Optionnel, requis pour certains scripts |
| **Mémoire RAM** | 4 Go minimum (8 Go recommandés) | Pour le fonctionnement du nœud |
| **Espace disque** | 500 Mo minimum | Pour les données de blockchain (state.db) |

---

## 2. Installation

### Option A : Installation via Go (compilation native)

```bash
# 1. Cloner le dépôt
git clone https://github.com/martialzinsou/pulsechain-fork.git

# 2. Accéder au répertoire du projet
cd pulsechain-fork

# 3. Installer les dépendances Go
go mod tidy

# 4. Compiler le binaire
go build -o pulsenoded ./cmd/pulsenoded

# 5. Vérification de la compilation
./pulsenoded --version  # Si implémenté, sinon passer à l'étape 6
```

### Option B : Déploiement via Docker

```bash
# 1. Cloner le dépôt
git clone https://github.com/martialzinsou/pulsechain-fork.git

# 2. Accéder au répertoire du projet
cd pulsechain-fork

# 3. Construire l'image Docker
docker build -t pulsechain-fork .

# 4. Lancer le conteneur
docker run -d \
  -p 8545:8545 \      # Port RPC JSON
  -p 30303:30303 \    # Port P2PDiscovery
  --name pulse-node \
  pulsechain-fork

# 5. Vérifier le démarrage
docker logs -f pulse-node
```

---

## 3. Démarrage du nœud

### 3.1 Démarrage direct (Go)

```bash
# Depuis le répertoire du projet
go run ./cmd/pulsenoded
```

**Sortie attendue :**
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

### 3.2 Démarrage via binaire compilé

```bash
# Exécutable déjà compilé
./pulsenoded
```

### 3.3 Démarrage via scriptShell

```bash
# Depuis la racine du projet
./scripts/init-node.sh
```

### 3.3 Démarrage via Docker Compose

```bash
docker compose up -d
```

**Résultat :** Le nœud tourne en tâche de fond. Pour voir les logs :

```bash
docker compose logs -f
```

Pour arrêter :

```bash
docker compose down
```

---

## 4. Configuration du réseau MetaMask

### 4.1 Ajout manuel du réseau

1. Ouvrir l'extension **MetaMask** dans votre navigateur
2. Cliquer sur le réseau actuel en haut à droite
3. Sélectionner **"Ajouter un réseau"** en bas de la liste
4. Remplir les champs suivants :

| Champ | Valeur à entrer |
|-------|-----------------|
| **Nom du réseau** | `PulseChain Fork (Martial Zinsou)` |
| **Nouvelle URL RPC** | `http://127.0.0.1:8545` |
| **ID de chaîne** | `369` |
| **Symbole de devise** | `PLS` |
| **URL de l'explorateur** | (optionnel) `http://localhost:8545` |

5. Cliquer sur le bouton **"Enregistrer"**
6. Le réseau PulseChain Fork apparaît désormais dans le sélecteur de réseau MetaMask

### 4.2 Utilisation avec le portefeuille

Une fois le réseau ajouté :

- **Voir le solde :** Le compte par défaut affichera le solde initial de **1 000 000 PLS** (allocation Genesis)
- **Effectuer des transactions :** Utiliser l'interface de transfert MetaMask en sélectionnant PulseChain Fork comme réseau actif
- **Interagir avec des contrats :** Utiliser l'adresse `0x2b5AD5c4795c026514f8317c7a215E218DcCD6cF` comme validateur/coinbase si nécessaire

---

## 5. Utilisation des endpoints RPC

### 5.1 Via cURL (ligne de commande)

```bash
# Vérifier le Chain ID
curl -X POST http://localhost:8545 \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_chainId","params":[],"id":1}'

# Obtenir le dernier bloc
curl -X POST http://localhost:8545 \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'

# Vérifier le solde d'une adresse
curl -X POST http://localhost:8545 \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_getBalance","params":["0x2b5AD5c4795c026514f8317c7a215E218DcCD6cF","latest"],"id":1}'
```

### 5.2 Via navigateur (Developer Tools)

1. Ouvrir les **Outils de développement** (F12)
2. Aller dans l'onglet **"Réseau"**
3. Sélectionner la requête POST vers `localhost:8545`
4. Vérifier l'onglet **"Réponse"** pour voir le JSON formaté
5. Utiliser l'extension **"JSON Viewer"** pour une meilleure lisibilité

### 5.3 Via Postman ou Insomnia

1. Créer une nouvelle requête de type **POST**
2. Définir l'URL : `http://localhost:8545`
3. Aller dans l'onglet **"Headers"** et ajouter :
   - `Content-Type: application/json`
4. Dans le corps de la requête (Body), sélectionner **raw** et **JSON**
5. Utiliser les exemples fournis dans la documentation API
6. Cliquer sur **"Send"** pour recevoir la réponse

---

## 6. Arrêt et redémarrage du nœud

### 6.1 Arrêt propre (Go)

Si le nœud a été lancé avec `go run ./cmd/pulsenoded` :

- Appuyer sur `Ctrl+C` dans le terminal
- Le nœud s'arrêtera proprement en sauvegardant l'état actuel

### 6.2 Arrêt via Docker

```bash
# Arrêt et suppression du conteneur
docker compose down

# Ou arrêt solo
docker stop pulse-node
```

### 6.3 Redémarrage

```bash
# Redémarrage rapide (Go)
go run ./cmd/pulsenoded

# Redémarrage via Docker
docker compose up -d
```

---

## 7. Dépannage courant

### Problème 1 : Le port 8545 est déjà utilisé

**Cause :** Un autre processus utilise déjà le port RPC.

**Solution :**
```bash
# Trouver quel processus utilise le port
lsof -i :8545

# Ou modifier le port dans cmd/pulsenoded/main.go
# puis recompiler : go build -o pulsenoded ./cmd/pulsenoded

# Ou redémarrer avec Docker (port différent)
docker run -p 8546:8545 ...
```

### Problème 2 : "Connection refused" lors des requêtes RPC

**Cause :** Le nœud ne tourne pas ou le firewall bloque le port.

**Solution :**
```bash
# Vérifier que le processus tourne
ps aux | grep pulsenoded

# Démarrer le nœud si arrêté
go run ./cmd/pulsenoded &

# Vérifier le firewall (macOS/Linux)
sudo ufw allow 8545/tcp

# Via Docker vérifier les ports exposés
docker ps -a
```

### Problème 3 : Solde affiché comme 0 dans MetaMask

**Cause :** Le réseau n'a pas été correctement ajouté ou le nœud n'a pas encore miné de blocs.

**Solution :**
```bash
# Vérifier que le nœud est bien en cours d'exécution
curl -X POST http://localhost:8545 -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","method":"eth_chainId","params":[],"id":1}'

# Vérifier le dernier bloc
curl -X POST http://localhost:8545 -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'

# Recharger MetaMask : décocher/recocher le réseau PulseChain Fork
```

### Problème 4 : Erreur "Method not found" sur les requêtes RPC

**Cause :** Incompatibilité version JSON-RPC ou méthode non supportée.

**Solution :**
- Vérifier que le nœud est à jour (dernier commit)
- Consulter la liste des méthodes supportées dans `docs/api.md`
- Mise à jour : `go get -u github.com/ethereum/go-ethereum`

---

## 7. Sauvegarde et restauration

### 7.1 Sauvegarde des données

```bash
# Via Docker (recommandé)
docker cp pulse-node:/app/data ./backup-data-$(date +%Y%m%d)

# ViaGo (si compilation native)
tar -czf backup-data-$(date +%Y%m%d).tar.gz ./data
```

### 7.2 Restauration des données

```bash
# Via Docker
docker stop pulse-node
docker cp backup-data-20260924.tar.gz pulse-node:/app/
docker start pulse-node

# Via Go
tar -xzf backup-data-20260924.tar.gz
# Redémarrer le nœud : go run ./cmd/pulsenoded
```

---

## 8. Mise à jour du projet

```bash
# Mise à jour depuis GitHub
git pull origin main

# Mise à jour des dépendances
go mod tidy

# Recompilation (optionnelle)
go build -o pulsenoded ./cmd/pulsenoded

# Redémarrage du nœud
./pulsenoded  # ou redémarrage Docker
```

---

**Auteur :** Martial Zinsou  
**Version :** 1.0.2  
**Dépôt :** https://github.com/martialzinsou/pulsechain-fork  
**Dernière mise à jour :** Septembre 2026  

---

**Besoin d'aide ?** Ouvrez une issue sur : https://github.com/martialzinsou/pulsechain-fork/issues