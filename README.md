# Heimdall

[![Go Report Card](https://goreportcard.com/badge/github.com/nexusblock/heimdall)](https://goreportcard.com/report/github.com/nexusblock/heimdall) [![CircleCI](https://circleci.com/gh/nexusblock/heimdall/tree/master.svg?style=shield)](https://circleci.com/gh/nexusblock/heimdall/tree/master) [![GolangCI](https://golangci.com/badges/github.com/nexusblock/heimdall.svg)](https://golangci.com/r/github.com/nexusblock/heimdall)


Validator node for Matic Network. It uses peppermint, customized [Tendermint](https://github.com/tendermint/tendermint).

### Install from source 

Make sure you have go1.11+ already installed

### Install 
```bash 
$ make install
```
### Init-heimdall 
```bash 
$ odind init
$ odind init --chain=mainnet        Will init with genesis.json for mainnet
$ odind init --chain=loki           Will init with genesis.json for loki
```
### Run-heimdall 
```bash 
$ odind start
```
#### Usage
```
$ odind start                       Will start for mainnet by default
$ odind start --chain=mainnet       Will start for mainnet
$ odind start --chain=loki          Will start for loki
$ odind start --chain=local         Will start for local with NewSelectionAlgoHeight = 0
```

### Run rest server
```bash 
$ odind rest-server 
```

### Run bridge
```bash 
$ odind bridge 
```

### Develop using Docker

You can build and run Heimdall using the included Dockerfile in the root directory:

```bash
docker build -t heimdall .
docker run heimdall
```

### Documentation 

Latest docs are [here](https://wiki.polygon.technology/docs/category/heimdall) 
