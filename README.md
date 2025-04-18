<div style="text-align:center">
  <img
  src="https://raw.githubusercontent.com/TerraDharitri/drt-go-chain/main/dharitri-logo.svg"
  alt="Dharitri">
</div>
<br>

[![](https://img.shields.io/badge/made%20by-Dharitri-blue.svg)](http://dharitri.org/)
[![](https://img.shields.io/badge/project-Dharitri%20Mainnet-blue.svg)](https://explorer.dharitri.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/TerraDharitri/drt-go-chain)](https://goreportcard.com/report/github.com/TerraDharitri/drt-go-chain)
[![codecov](https://codecov.io/gh/TerraDharitri/drt-go-chain/branch/main/graph/badge.svg?token=MYS5EDASOJ)](https://codecov.io/gh/TerraDharitri/drt-go-chain)
[![Contributors](https://img.shields.io/github/contributors/TerraDharitri/drt-go-chain)](https://github.com/TerraDharitri/drt-go-chain/graphs/contributors)

# drt-go-chain

The go implementation for the Dharitri protocol

## Installation and running

In order to join the network as an observer or as a validator, the required steps to **build from source and setup explicitly** are explained below.

### Step 1: install & configure go:
The installation of go should proceed as shown in official golang installation guide https://golang.org/doc/install . In order to run the node, minimum golang version should be 1.17.6.

### Step 2: clone the repository and build the binaries:
The main branch that will be used is the main branch. Alternatively, an older release tag can be used.

```
# set $GOPATH if not set and export to ~/.profile along with Go binary path
$ if [[ $GOPATH=="" ]]; then GOPATH="$HOME/go" fi
$ mkdir -p $GOPATH/src/github.com/TerraDharitri
$ cd $GOPATH/src/github.com/TerraDharitri
$ git clone https://github.com/TerraDharitri/drt-go-chain
$ cd drt-go-chain && git checkout main
$ cd cmd/node && go build
```
The node depends on the Wasm Virtual Machine, which is automatically managed by the node.

### Step 3: creating the node’s identity:
In order to be registered in the Dharitri Network, a node must possess 2 types of (secret key, public key) pairs. One is used to identify the node’s credential used to generate transactions (having the sender field its account address) and the other is used in the process of the block signing. Please note that this is a preliminary mechanism, in the next releases the first (private, public key) pair will be dropped when the staking mechanism will be fully implemented. To build and run the keygenerator, the following commands will need to be run:

```
$ cd $GOPATH/src/github.com/TerraDharitri/drt-go-chain/cmd/keygenerator
$ go build
$ ./keygenerator
```

### Start the node 
#### Step 4a: Join Dharitri testnet:
Follow the steps outlined [here](https://docs.dharitri.org/validators/nodes-scripts/config-scripts/). This is because in order to join the testnet you need a specific node configuration.
______
OR
______
#### Step 4b: copying credentials and starting a node in a separate network:
The previous generated .pem file needs to be copied in the same directory where the node binary resides in order to start the node.

```
$  cp validatorKey.pem ./../node/config/
$  cd ../node && ./node
```

The node binary has some flags defined (for a brief description, the user can use --help flag). Those flags can be used to directly alter the configuration values defined in .toml/.json files and can be used when launching more than one instance of the binary. 

### Running the tests	
```	
$ go test ./...	
```

## Compiling new fields in .proto files (should be updated when required PR will be merged in gogo protobuf main branch):
1. Download protoc compiler: https://github.com/protocolbuffers/protobuf/releases 
 (if you are running under linux on a x64 you might want to download protoc-3.11.4-linux-x86_64.zip)
2. Expand archive, copy the /include/google folder in /usr/include using <br>
`sudo cp -r google /usr/include`
3. Copy bin/protoc using <br>
`sudo cp protoc  /usr/bin` 
4. Fetch the repo github.com/TerraDharitri/protobuf
5. Compile gogo slick & copy binary using
```
cd protoc-gen-gogoslick
go build
sudo cp protoc-gen-gogoslick /usr/bin/
```

Done

## Running p2p Prometheus dashboards
1. Start the node with `--p2p-prometheus-metrics` flag. This exposes a metrics collection at http://localhost:8080/debug/metrics/prometheus (port defined by -rest-api-interface flag, default 8080)
2. Clone libp2p repository: `git clone https://github.com/libp2p/go-libp2p`
3. `cd go-libp2p/dasboards/swarm` and under the 
```  
"templating": {
   "list": [
```
section, add the following lines:
```
{
  "hide": 0,
  "label": "datasource",
  "name": "DS_PROMETHEUS",
  "options": [],
  "query": "prometheus",
  "refresh": 1,
  "regex": "",
  "type": "datasource"
},
```
(this step will be removed once it will be fixed on libp2p)
4. `cd ..` to dashboards directory and update the port of `host.docker.internal` from `prometheus.yml` to node's Rest API port(default `8080`)
5. From this directory, run the following docker compose command: 
```
sudo docker compose -f docker-compose.base.yml -f docker-compose-linux.yml up --force-recreate
```
**Note:** If you choose to install the new Docker version manually, please make sure that installation is done for all users of the system. Otherwise, the docker command will fail because it needs the super-user privileges.
6. The preconfigured dashboards should be now available on Grafana at http://localhost:3000/dashboards

## Progress

### Done
- [x] Cryptography
  - [x] Schnorr Signature
  - [x] Belare-Neven Signature
  - [x] BLS Signature
  - [x] Modified BLS Multi-signature
- [x] Datastructures
  - [x] Transaction
  - [x] Block
  - [x] Account
  - [x] Trie
- [x] Execution
  - [x] Transaction
  - [x] Block
  - [x] State update
  - [x] Synchronization
  - [x] Shard Fork choice
- [x] Peer2Peer - libp2p
- [x] Consensus - SPoS
- [x] Sharding - fixed number
  - [x] Transaction dispatcher 
  - [x] Transaction
  - [x] State
  - [x] Network - Message dispatching
- [x] MetaChain
  - [x] Data Structures
  - [x] Block Processor
  - [x] Interceptors/Resolvers
  - [x] Consensus
- [x] Block K finality scheme
- [x] VM - K-Framework
  - [x] K Framework go backend
  - [x] IELE Core
  - [x] IELE Core tests
  - [x] IELE Adapter
- [x] Smart Contracts on a Sharded Architecture
  - [x] Concept reviewed
  - [x] VM integration
  - [x] SC Deployment
- [x] Governance
  - [x] Concept reviewed
- [x] Economics
  - [x] Concept reviewed  
- [x] Optimizations
  - [x] Randomness
  - [x] Consensus
- [x] Bootstrap from storage
- [x] Testing 
  - [x] Unit tests
  - [x] Integration tests
  - [x] TeamCity continuous integration
  - [x] Manual testing
- [x] Epochs
  - [x] Nodes dispatcher (shuffling)
- [x] Network sharding
  - [x] Optimized wiring protocol
- [x] VM
  - [x] EVM Core
  - [x] EVM Core tests
  - [x] EVM Adapter
- [x] Fee structure
- [x] Smart Contracts on a Sharded Architecture
  - [x] Async callbacks
- [x] Testing
  - [x] Automate tests with AWS
  - [x] Nodes Monitoring
- [x] DEX integration

### In progress

- [ ] Smart Contracts on a Sharded Architecture
  - [ ] Dependency checker + SC migration
  - [ ] Storage rent + SC backup & restore
- [ ] Adaptive State Sharding
  - [ ] Splitting
  - [ ] Merging 
  - [ ] Redundancy
- [ ] Privacy
- [ ] Interoperability
- [ ] Optimizations
  - [ ] Smart Contract 
- [ ] Governance
  - [ ] SC for DRT IP
  - [ ] Enforced Upgrade mechanism for voted DRT IP
- [ ] Bugfixing


## Contribution
Thank you for considering to help out with the source code! We welcome contributions from anyone on the internet, and are grateful for even the smallest of fixes to Dharitri!

If you'd like to contribute to Dharitri, please fork, fix, commit and send a pull request for the maintainers to review and merge into the main code base. If you wish to submit more complex changes though, please check up with the core developers first here on GitHub, to ensure those changes are in line with the general philosophy of the project and/or get some early feedback which can make both your efforts much lighter as well as our review and merge procedures quick and simple.

Please make sure your contributions adhere to our coding guidelines:

 - Code must adhere to the official Go [formatting](https://golang.org/doc/effective_go.html#formatting) guidelines.
 - Code must be documented adhering to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary) guidelines.
 - Pull requests need to be based on and opened against the main branch.
 - Commit messages should be prefixed with the package(s) they modify.
    - E.g. "outport/process: fixed a typo"

Please see the [documentation](https://docs.dharitri.org/) for more details on the Dharitri protocol.

