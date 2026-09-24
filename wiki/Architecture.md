# PulseChain Fork - Architecture Système

> **Auteur :** Martial Zinsou  
> **Dépôt :** https://github.com/martialzinsou/pulsechain-fork  

## 1. Vue d'ensemble de l'architecture

Le projet **PulseChain Fork** adopte une architecture modulaire en couches, séparant clairement les responsabilités entre le moteur de blockchain, le consensus, le registre d'état et l'interface RPC. Cette conception facilite l'extension, le débogage et le déploiement.

```mermaid
graph TB
    %% Styles
    classDef system fill:#e1f5fe,stroke:#01579b,stroke-width:2px
    classDef module fill:#fff3e0,stroke:#ef6c00,stroke-width:1px
    classDef interface fill:#e8f5e9,stroke:#2e7d32,stroke-width:1px

    %% Couche Utilisateur
    subgraph UI["Couche Interface Utilisateur"]
        direction LR
        U1[MetaMask] -->|JSON-RPC| S1[Serveur RPC]
        U2[Ethers.js] -->|JSON-RPC| S1
        U3[Foundry/Hardhat] -->|JSON-RPC| S1
    end

    %% Couche Serveur
    subgraph SR["Couche Serveur RPC"]
        direction LR
        S1[Serveur JSON-RPC Go] -->|Route /rpc| S2[Gestionnaire requêtes]
    end

    %% Couche Blockchain
    subgraph BC["Couche Moteur Blockchain"]
        direction LR
        B1[Blockchain Engine] -->|Stocke/Retrieve| S2
        B1 -->|Validate| CO[Consensus PoA]
        B1 -->|Mise à jour| LD[StateDB Ledger]
    end

    %% Courage Consensus
    subgraph CO["Courage Consensus PoA"]
        direction LR
        C1[Proof-of-Authority] -->|Calcul difficulty| P1[Paramètres PulseChain]
        C1 -->|Valide blocs| B1
    end

    %% Courage Ledger
    subgraph LD["Courage StateDB Ledger"]
        direction LR
        L1[Gestion comptes] -->|Solde transfer| B1
        L1 -->|Mise à jour état| S2
    end

    %% Liens cross-couche
    S2 -->|Requêtes RPC| B1
    S2 -->|État comptes| L1

    %% Styles appliqués
    class UI,S1,R1 system
    class B1,CO,LD,BC module
    class C1,P1 interface
```

---

## 2. Diagramme des composants (Component Diagram)

```mermaid
componentDiagram
    %% Composants principaux
    class PulseNode "PulseChain Fork Node"
    class RPCServer "Serveur JSON-RPC (internal/rpc/server.go)"
    class BlockEngine "Moteur Blockchain (internal/blockchain)"
    class Consensus "Consensus PoA (internal/consensus)"
    class StateMgr "StateDB Ledger (internal/ledger)"
    class Genesis "Configuration Génèse (genesis/)"

    %% Interfaces et dépendances
    PulseNode --> RPCServer : expose API
    PulseNode --> BlockEngine : gérer blocs
    PulseNode --> Consensus : valider blocs
    PulseNode --> StateMgr : état comptes
    PulseNode --> Genesis : paramètres initial

    %% Flux de données
    RPCServer <|-- BlockEngine : appelle méthodes
    RPCServer <|-- Consensus : valide blocs
    RPCServer <|-- StateMgr : requête soldes

    %% Styles
    PulseNode * "v1.0.2" : version
    RPCServer "Auteur: Martial Zinsou"
    BlockEngine "Go 1.21"
    Consensus "PoA 3s bloc"
    StateMgr "sync.RWMutex"
    Genesis "ChainID 369"
```

---

## 3. Diagramme de déploiement (Deployment Diagram)

```mermaid
deploymentDiagram
    %% Nœuds matériel
    node Server "Serveur Linux/macOS/Windows" {
        node Docker "Conteneur Docker" {
            PulseNode : PulseChain Fork v1.0.2
        }
    }

    %% Interfaces
    PulseNode -->|Port 8545 (RPC)| ClientWeb3 : "MetaMask / DApps"
    PulseNode -->|Port 30303 (P2P)| PairNetwork : "Autres nœuds"

    %% Fichiers et volumes
    PulseNode --> VolumeData : "./data persistence"
    PulseNode --> ConfigFile : "genesis.json"

    %% Styles
    Server "Environment: Production/Development"
    Docker "Image: golang:1.21-alpine"
    ClientWeb3 "Navigateur / Application"
```

---

## 4. Flux de données complet

```mermaid
flowchart TD
    %% Déclenchements
    T1[Utilisateur soumet tx via MetaMask] --> T2
    T2[MetaMask → HTTP POST :8545] --> T3
    T3[Serveur RPC reçoit requête] --> T4
    T4[Validation params JSON] --> T5
    T5{Validation OK?} -- Non --> T6[Erreur RPC -32602]
    T5 -- Oui --> T6
    T6[Miner bloc] --> T7
    T7[Validation consensus PoA] --> T8
    T8{Bloc valide?} -- Non --> T9[Rejet bloc]
    T8 -- Oui --> T10
    T10[Mise à jour StateDB] --> T11
    T11[Broadcast réseau P2P] --> T12
    T12[Autres nœuds reçivent] --> T13
    T13[Mise à jour chaînes synchronisées] --> T14
    T14[Confirmation utilisateur MetaMask] --> Fine

    %% Styles
    classDef process fill:#e3f2fd,stroke:#1565c0,stroke-width:1px
    classDef decision fill:#fff9c4,stroke:#fbc02d,stroke-width:1px
    classDef error fill:#ffebee,stroke:#c62828,stroke-width:1px

    class T2,T4,T7,T10,T13 process
    class T5,T8 decision
    class T6,T9 error
```

---

## 4. Points d'extension architecturaux

| Point d'extension | Description | Fichier à modifier |
|-------------------|-------------|-------------------|
| Nouveau hard fork | Ajouter paramètres dans `genesis.json` et activer drapeaux | `genesis/genesis.go` |
| Nouvelle méthode RPC | Implémenter dans `internal/rpc/server.go` | `internal/rpc/server.go` |
| Nouveau algorithme consensus | Remplacer `proof.go` avec nouvelle logique | `internal/consensus/proof.go` |
| Couche réseau P2P | Intégrer libp2p ou libp2p-go | `cmd/pulsenoded/main.go` |
| Support multi-chain | Parameteriser ChainID dynamiquement | `genesis/genesis.go` |

---

**Auteur :** Martial Zinsou  
**Version :** 1.0.2  
**Dernière mise à jour :** Septembre 2026