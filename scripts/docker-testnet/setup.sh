#!/usr/bin/env bash

set -eux

export DOCKERTESTNETDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

TerraDharitriTESTNETSCRIPTSDIR="$(dirname "$DOCKERTESTNETDIR")/testnet"

source "$DOCKERTESTNETDIR/variables.sh"
source "$DOCKERTESTNETDIR/functions.sh"
source "$TerraDharitriTESTNETSCRIPTSDIR/include/config.sh"
source "$TerraDharitriTESTNETSCRIPTSDIR/include/build.sh"

cloneRepositories

prepareFolders

buildConfigGenerator

generateConfig

copyConfig

copySeednodeConfig
updateSeednodeConfig

copyNodeConfig
updateNodeConfig

createDockerNetwork

startSeedNode
startObservers
startValidators

if [ $USE_PROXY -eq 1 ]; then
  buildProxyImage
  prepareFolders_Proxy
  copyProxyConfig
  updateProxyConfigDocker
  startProxyDocker
fi

