export DHARITRITESTNETSCRIPTSDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

source "$DHARITRITESTNETSCRIPTSDIR/variables.sh"
source "$DHARITRITESTNETSCRIPTSDIR/include/config.sh"
source "$DHARITRITESTNETSCRIPTSDIR/include/build.sh"


echo "Preparing folders..."
prepareFolders
echo "Building config generator..."
buildConfigGenerator
echo "Generating config..."
generateConfig
echo "Copying config..."
copyConfig
echo "Copying seednode config..."
copySeednodeConfig
echo "Updating seednode config..."
updateSeednodeConfig

echo "Copying node config..."
copyNodeConfig
echo "Updating node config..."
updateNodeConfig

if [ $USE_PROXY -eq 1 ]; then
  echo "Preparing proxy folders..."
  prepareFolders_Proxy
  echo "Copying proxy config..."
  copyProxyConfig
  echo "Updating proxy config..."
  updateProxyConfig
fi

if [ $USE_TXGEN -eq 1 ]; then
  echo "Preparing TxGen folders..."
  prepareFolders_TxGen
  echo "Copying TxGen config..."
  copyTxGenConfig
  echo "Updating TxGen config..."
  updateTxGenConfig
fi

if [ $USE_HARDFORK -eq 1 ]; then
  echo "Changing config for hardfork..."
  changeConfigForHardfork
fi

if [ $COPY_BACK_CONFIGS -eq 1 ]; then
  echo "Copying back configs..."
  copyBackConfigs
fi#!/usr/bin/env bash

# Prepare the configuration files for a testnet:
#		Build the drt-chain-deploy-go binary and run it
#		Copy the generated files to the testnet folder
#		Copy configuration for the seednode, nodes, proxy and txgen into the testnet folder

export DHARITRITESTNETSCRIPTSDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

source "$DHARITRITESTNETSCRIPTSDIR/variables.sh"
source "$DHARITRITESTNETSCRIPTSDIR/include/config.sh"
source "$DHARITRITESTNETSCRIPTSDIR/include/build.sh"

prepareFolders

buildConfigGenerator

generateConfig

copyConfig

copySeednodeConfig
updateSeednodeConfig

copyNodeConfig
updateNodeConfig

if [ $USE_PROXY -eq 1 ]; then
  prepareFolders_Proxy
  copyProxyConfig
  updateProxyConfig
fi

if [ $USE_TXGEN -eq 1 ]; then
  prepareFolders_TxGen
  copyTxGenConfig
  updateTxGenConfig
fi

if [ $USE_HARDFORK -eq 1 ]; then
  changeConfigForHardfork
fi

if [ $COPY_BACK_CONFIGS -eq 1 ]; then
  copyBackConfigs
fi
