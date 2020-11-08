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

Note: ERC20 tokens are not currently supported.

## Setup

Currently, the CLI must be built from source:

1. Run `go install` on root of the repository 
2. cd into `/cmd`
3. run `go build`

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
