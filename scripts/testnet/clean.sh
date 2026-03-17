#!/usr/bin/env bash

# Delete the entire testnet folder, which includes configuration, executables and logs.

export TerraDharitriTESTNETSCRIPTSDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

source "$TerraDharitriTESTNETSCRIPTSDIR/variables.sh"

echo "Removing $TESTNETDIR..."
rm -rf $TESTNETDIR
