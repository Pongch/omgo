# OMGO Client Library
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pongch/omgo)](https://pkg.go.dev/github.com/pongch/omgo)  [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/pongch/omgo/blob/master/LICENSE)  [![Go Report Card](https://goreportcard.com/badge/github.com/pongch/omgo)](https://goreportcard.com/report/github.com/pongch/omgo)


golang client library for OMG Network

**please note that this repo is still in experimental stage. use it at your own risk **

## Compatibility 

compatible with OMG Network V1 Beta

## Functionality 

root chain:
1. Deposit ETH/ERC20 to the root chain contract
2. Start standard exit with ETH/ERC20
3. Process exit ETH/ERC20 back to root chain 

*Note*: current functionalities are limited, but /abi package provides binder which give clients all available root chain contract calls from scratch, this can be found in:
`github.com/pongch/omgo/abi`

child chain:
1. get utxos via `/account.get_utxos`
2. get network status via `/status.get`
3. get balance via `/account.get_balance`
4. create payment transaction via convenient transaction API: `/transaction.create` 
5. get transaction information via: `/transaction.get` 

## Packages

| package                           | description                                                   |
|-----------------------------------|---------------------------------------------------------------|
| github.com/pongch/omgo/abi        | provides clients with all available root chain contract calls |
| github.com/pongch/omgo/childchain | provides clients with child chain API functionality           |
| github.com/pongch/omgo/childchain | provides client with root chain contract calls                |
| github.com/pongch/omgo/util       | provides client with utility functions                        |
| github.com/pongch/omgo/e2e        | end to end test                                               |
| github.com/pongch/omgo/cmd/cli    | CLI package                                                   |


## Getting Started

import the client library via golang import

```
import "github.com/pongch/omgo"
```

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE.md](LICENSE.md) file for details

# How to run via CLI

The OMGO CLI is used to interact with an OMG Network from the command line. Technical details about Ethereum and Plasma are abstracted from the user.
Connectivity to Ethereum via a local RPC node or Infura is required.

## Download Binary

`omgo` is available as binaries at https://github.com/pongch/omgo/releases. Download the binary that matches your operating system.

After downloading the latest binary file (on macOS):
1. rename the file `$ mv "name_of_your_binary_file" omgo`
2. set execute permissions `$ chmod u+x ./omgo`
3. run it `$ ./omgo --help` make sure your security settings allow you to run a downloaded executable file

## Build from source

Currently, the CLI must be built from source:

1. Run `go install` on root of the repository 
2. cd into `/cmd`
3. run `go build`

## Configuration

the CLI will read from the following environment variables

```
OMGO_CCCLIENT // client to interact with the child chain. This is currently Watcher-info endpoint
OMGO_RCCLIENT // root chain client to interact with the Ethereum network, ie. local node or node provider like Infura
OMGO_FWCONTRACT // plasma framework contract address
OMGO_VCONTRACT // plasma vault ID=1 (ETH Vault) contract
OMGO_V2CONTRACT // plasma vault ID=2 (ERC20 Vault) contract
OMGO_ECONTRACT // plasma exit game contract
OMGO_PKEY // private key of the wallet/account to transact from

```

example env:

```
export OMGO_CCCLIENT="https://watcher-info.ropsten.v1.omg.network"
export OMGO_RCCLIENT="https://ropsten.infura.io/v3/e2d2eee81e774cd3a4915d994dfec840"
export OMGO_FWCONTRACT="0x96d5d8bc539694e5fa1ec0dab0e6327ca9e680f9"
export OMGO_VCONTRACT="0x895cc6f20d386f5c0deae08b08ccfec9f821e7d9"
export OMGO_V2CONTRACT="0x18e15c2cdc003b845b056f8d6b6a91ab33d3f182"
export OMGO_ECONTRACT="0x08c569c5774110eb84a80b292e6d6f039e18915a"
export OMGO_PKEY="CD5994C7E2BF03202C59B529B76E5582266CEB384F02D32B470AC57112D0C6E7"
```

## Create a keypair

```
omgo create keypair
```

## Deposit into the OMG Network

```
omgo approve --currency=token_to_approve --amount=amount_to_approve
```

```
omgo deposit  --amount=amount_in_wei
```
optional flags:
```
--currency=non-eth-currency
```

## Retrieve a List of UTXOs

```
omgo get utxos --from=public_address
```

## Retrieve the Balance of an Account

```
omgo get balance --from=public_address
```

## Send transaction

creates a send transaction from the balance of the privatekey to a recipient wallet 
(currently only support 1 recipient). Makes an ETH transaction by default

```
omgo send --to=address --amount=amount 
```

optional flags:
```
--currency=token_address
--feetoken=token_address
--metadata=hex_string
```

## Exit to Ethereum

```
omgo exit --utxo=UTXO_position
```

## Process Exit

process matured exits, process ETH by default. use vault id=2 for ERC20 tokens
```
omgo process --number=number_of_exit_to_process 
```
optional flags:
```
--id=vault_id
--currency=currency_to_process
```
