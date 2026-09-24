#!/bin/bash
# ==============================================================================
# Script d'initialisation et de démarrage du nœud PulseChain Fork
# Auteur : Martial Zinsou
# ==============================================================================

set -e

echo "======================================================"
echo "    PULSECHAIN FORK - DÉMARRAGE DU NŒUD               "
echo "    Auteur : Martial Zinsou                           "
echo "======================================================"

DATA_DIR=${DATA_DIR:-"./data"}
RPC_PORT=${RPC_PORT:-8545}
P2P_PORT=${P2P_PORT:-30303}
CHAIN_ID=${CHAIN_ID:-369}

echo "[1/4] Vérification des prérequis..."
mkdir -p "$DATA_DIR"

echo "[2/4] Chargement de la configuration genesis (ChainID: $CHAIN_ID)..."
if [ ! -f "pulsechain-fork/genesis/genesis.json" ]; then
    echo "[-] Erreur : genesis.json introuvable !"
    exit 1
fi

echo "[3/4] Préparation des dépendances..."
if command -v go >/dev/null 2>&1; then
    go mod tidy
    echo "[4/4] Lancement du nœud via Go..."
    go run ./pulsechain-fork/cmd/pulsenoded
else
    echo "[!] Go n'est pas détecté dans le PATH."
    echo "[!] Vous pouvez exécuter le nœud via Docker :"
    echo "    docker compose up -d"
fi
