#!/bin/bash
# ==============================================================================
# Script d'interaction RPC pour tester le nœud PulseChain Fork
# Auteur : Martial Zinsou
# ==============================================================================

RPC_URL=${RPC_URL:-"http://localhost:8545"}

echo "======================================================"
echo "    PULSECHAIN FORK - CLIENT RPC TEST                 "
echo "    Auteur : Martial Zinsou                           "
echo "======================================================"
echo "URL RPC : $RPC_URL"
echo ""

echo "[1] Test eth_chainId :"
curl -s -X POST "$RPC_URL" \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_chainId","params":[],"id":1}' | python3 -m json.tool || true

echo ""
echo "[2] Test web3_clientVersion :"
curl -s -X POST "$RPC_URL" \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"web3_clientVersion","params":[],"id":2}' | python3 -m json.tool || true

echo ""
echo "[3] Test eth_blockNumber :"
curl -s -X POST "$RPC_URL" \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":3}' | python3 -m json.tool || true

echo ""
echo "[4] Test eth_getBalance (Coinbase Genesis) :"
curl -s -X POST "$RPC_URL" \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_getBalance","params":["0x2b5AD5c4795c026514f8317c7a215E218DcCD6cF", "latest"],"id":4}' | python3 -m json.tool || true

echo ""
echo "[5] Test eth_getBlockByNumber (Bloc 0 - Genesis) :"
curl -s -X POST "$RPC_URL" \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x0", true],"id":5}' | python3 -m json.tool || true
