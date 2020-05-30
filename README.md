# OMGO Client Library

golang client library for OMG Network

**please note that this repo is still in experimental stage. use it at your own risk **

## Compatibility 

compatible with OMG Network V1 Beta

## Functionality 

root chain:
1. Deposit ETH to the root chain contract
2. Start standard exit with ETH
3. Process exit UTXO back to root chain 

*Note*: current functionalities are limited, but /abi package provides binder which give clients all available root chain contract calls from scratch, this can be found in:
`github.com/pongch/omgo/abi`

child chain:
1. get utxos via `/account.get_utxos`
2. get network status via `/status.get`
3. get balance via `/account.get_balance`
4. create payment transaction via convenient transaction API: `/transaction.create` 
5. get transaction information via: `/transaction.get` 

## Packages

1. githu.com/pongch/omgo/abi -> provides clients with all available root chain contract calls
2. github.com/pongch/omgo/childchain -> provides clients with child chain API functionality
3. github.com/pongch/omgo/childchain -> provides client with root chain contract calls
4. github.com/pongch/omgo/util -> provides client with utility functions
5. github.com/pongch/omgo/e2e -> end to end test
6. github.com/pongch/omgo/cmd/cli -> CLI package

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

## Installation

Currently, the CLI must be built from source:

1. Run `go install` on root of the repository 
2. cd into `/cmd/cli`
3. run `go build`

## Create an Account (Keypair)

```
plasma_cli create account
```

## Deposit ETH into the OMG Network

```
plasma_cli deposit --privatekey="private_key" --client="local_rpc_server_or_Infura_URL" --contract="contract_address" --amount=amount_in_wei --owner="owner_address"
```


## Retrieve a List of UTXOs

```
plasma_cli get utxos --watcher="watcher_URL" --address="public_address"
```

## Check the Balance of an Account

```
plasma_cli get balance --watcher="watcher_URL" --address="public_address"
```
## Send transaction

creates a send transaction from the balance of the privatekey to a recipient wallet 
(currently only support 1 recipient). Makes an ETH transaction by default

```
plasma_cli send --to="address" --privatekey="privatekey" --amount=amount --watcher="watcher_url"
```

optional arguements:
```
--currency="token_address" 
--feetoken="token_address"
--feeamount=fee_amount
--metadata="hex_string"
```

## Exit to Ethereum

```
plasma_cli exit --utxo=UTXO_position --privatekey="private_key" --contract="contract_address" --watcher="watcher_url" --client="local_rpc_server_or_Infura_URL"
```

## Process Exit

```
plasma_cli process --privatekey="private_key" --contract="contract_address" --watcher="watcher_url" --client="local_rpc_server_or_Infura_URL"
```
