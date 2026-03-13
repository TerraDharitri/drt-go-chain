#!/usr/bin/env bash

export TerraDharitriTESTNETSCRIPTSDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

export TESTNETMODE=$1
export EXTRA=$2

source "$TerraDharitriTESTNETSCRIPTSDIR/variables.sh"
source "$TerraDharitriTESTNETSCRIPTSDIR/include/config.sh"
source "$TerraDharitriTESTNETSCRIPTSDIR/include/build.sh"
source "$TerraDharitriTESTNETSCRIPTSDIR/include/validators.sh"
source "$TerraDharitriTESTNETSCRIPTSDIR/include/observers.sh"
source "$TerraDharitriTESTNETSCRIPTSDIR/include/tools.sh"

prepareFolders

# Phase 1: build seednode and node executables
buildSeednode
buildNode

# Phase 2: generate configuration
if [ $ALWAYS_UPDATE_CONFIGS -eq 1 ]; then
  copyConfig
fi

copySeednodeConfig
updateSeednodeConfig

if [ $ALWAYS_UPDATE_CONFIGS -eq 1 ]; then
  copyNodeConfig
  updateNodeConfig
fi


# Phase 3: start the Seednode
startSeednode
showTerminalSession "TerraDharitri-tools"
echo "Waiting for the Seednode to start ($SEEDNODE_DELAY s)..."
sleep $SEEDNODE_DELAY

# Phase 4: start the Observer Nodes and Validator Nodes
startObservers
startValidators
showTerminalSession "TerraDharitri-nodes"
echo "Waiting for the Nodes to start ($NODE_DELAY s)..."
sleep $NODE_DELAY

# Phase 5: build the Proxy and TxGen
if [ $USE_PROXY -eq 1 ]; then
  prepareFolders_Proxy
  buildProxy
fi
if [ $USE_TXGEN -eq 1 ]; then
  prepareFolders_TxGen
  buildTxGen
fi

# Phase 6: start the Proxy
if [ $USE_PROXY -eq 1 ]; then
  startProxy
  echo "Waiting for the Proxy to start ($PROXY_DELAY s)..."
  sleep $PROXY_DELAY
fi

# Phase 7: start the TxGen, with or without regenerating the accounts
if [ $USE_TXGEN -eq 1 ]; then
  if [ -n "$TXGEN_REGENERATE_ACCOUNTS" ]
  then
    echo "Starting TxGen with account generation..."
    startTxGen_NewAccounts
  else
    echo "Starting TxGen with existing accounts..."
    startTxGen_ExistingAccounts
  fi
fi
