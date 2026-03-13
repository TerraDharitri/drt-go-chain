#!/usr/bin/env bash

export TerraDharitriTESTNETSCRIPTSDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

$TerraDharitriTESTNETSCRIPTSDIR/stop.sh
$TerraDharitriTESTNETSCRIPTSDIR/clean.sh
$TerraDharitriTESTNETSCRIPTSDIR/config.sh
$TerraDharitriTESTNETSCRIPTSDIR/start.sh $1
