// Copyright 2019 OmiseGO Pte Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package parser

import (
	"net/http"

	"github.com/omisego/plasma-cli/childchain"
	"github.com/omisego/plasma-cli/rootchain"
	"github.com/omisego/plasma-cli/util"
	log "github.com/sirupsen/logrus"
	"strconv"
	"context"
	"math/big"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	get              = kingpin.Command("get", "Get a resource.").Default()
	getUTXO          = get.Command("utxos", "Retrieve UTXO data from the Watcher service")
	watcherURL       = get.Flag("watcher", "FQDN of the Watcher in the format https://watcher.path.net").Required().String()
	ownerUTXOAddress = get.Flag("address", "Owner address to search UTXOs").String()
	getBalance       = get.Command("balance", "Retrieve balance of an address from the Watcher service")
	status           = get.Command("status", "Get status from the Watcher")
	transaction      = get.Command("transaction", "get transaction details")
	txHash           = get.Flag("txhash", "transaction hash of the transaction you would like to get the information for").String()

	getExit             = get.Command("exit", "Get UTXO exit information")
	getExitUTXOPosition = getExit.Flag("utxo", "Get UTXO exit information").Required().Int()

	deposit         = kingpin.Command("deposit", "Deposit ETH or ERC20 into the Plamsa MoreVP smart contract.")
	privateKey      = deposit.Flag("privatekey", "Private key of the account used to send funds into Plasma MoreVP").Required().String()
	client          = deposit.Flag("client", "Address of the Ethereum client. Infura and local node supported https://rinkeby.infura.io/v3/api_key or http://localhost:8545").Required().String()
	contract        = deposit.Flag("contract", "Address of the Plasma MoreVP smart contract").Required().String()
	depositAmount   = deposit.Flag("amount", "Amount to deposit in wei").Required().Int()
	depositCurrency = deposit.Flag("currency", "ERC20 token address for the deposit, leave blank for ETH").Default(childchain.EthCurrency).String()


	send             = kingpin.Command("send", "Create a transaction on the OmiseGO Plasma MoreVP network")
	to               = send.Flag("to", "Wallet address of the recipient").Required().String()
	privatekey       = send.Flag("privatekey", "Privatekey from a wallet to send from").Required().String()
	amount           = send.Flag("amount", "Amount to transact").Required().Uint64()
	currency         = send.Flag("currency", "currency of the amount to send, default to ETH").Default(childchain.EthCurrency).String()
	watcherSubmitURL = send.Flag("watcher", "FQDN of the Watcher in the format http://watcher.path.net").Required().String()
	feetoken         = send.Flag("feetoken", "set the token to be used as transaction fee, default to ETH").Default(childchain.EthCurrency).String()
	feeamount        = send.Flag("feeamount", "set the amount to be used as transaction fee, default to 0").Default("0").Uint64()
	metadata         = send.Flag("metadata", "additional metadata to send with the transaction, up to 32 bytes").Default(childchain.DefaultMetadata).String()

	exit           = kingpin.Command("exit", "Standard exit a UTXO back to the root chain.")
	watcherExitURL = exit.Flag("watcher", "FQDN of the Watcher in the format http://watcher.path.net").Required().String()
	contractExit   = exit.Flag("contract", "Address of the Plasma MoreVP smart contract").Required().String()
	utxoPosition   = exit.Flag("utxo", "UTXO Position that will be exited").Required().String()
	exitPrivateKey = exit.Flag("privatekey", "Private key of the UTXO owner").Required().String()
	clientExit     = exit.Flag("client", "Address of the Ethereum client. Infura and local node supported https://rinkeby.infura.io/v3/api_key or http://localhost:8545").Required().String()

	process           = kingpin.Command("process", "Process all exits that have completed the challenge period")
	processContract   = process.Flag("contract", "Address of the Plasma MoreVP smart contract").Required().String()
	// processToken      = process.Flag("token", "Token address to process for standard exits").Required().String()
	processPrivateKey = process.Flag("privatekey", "Private key used to fund the gas for the smart contract call").Required().String()
	processExitClient = process.Flag("client", "Address of the Ethereum client. Infura and local node supported https://rinkeby.infura.io/v3/api_key or http://localhost:8545").Required().String()

	create        = kingpin.Command("create", "Create a resource.")
	createAccount = create.Command("account", "Create an account consisting of Public and Private key")

	// challengeexit  = kingpin.Command("challenge exit", "Challenge an invalid exit on the root chain")
	// ceContract     = challengeexit.Flag("contract", "Address of the Plasma MoreVP smart contract").Required().String()
	// ceClient       = challengeexit.Flag("client", "Address of the Ethereum client. Infura and local node supported https://rinkeby.infura.io/v3/api_key or http://localhost:8545").Required().String()
	// cePrivateKey   = challengeexit.Flag("privatekey", "Private key used to fund the gas for the smart contract call").Required().String()
	// ceUTXOPosition = challengeexit.Flag("utxo", "the UTXO to call challenge exit").Required().String()
	// ceWatcherURL   = challengeexit.Flag("watcher", "FQDN of the Watcher in the format http://watcher.path.net").Required().String()
)

func ParseArgs() {
	switch kingpin.Parse() {
	case getUTXO.FullCommand():
		//plasma_cli get utxos --address=0x944A81BeECac91802787fBCFB9767FCBf81db1f5 --watcher=http://watcher.path.net
		chch, err := childchain.NewClient(*watcherURL, &http.Client{})
		if err != nil {
			log.Error(err)
		}
		utxos, err := chch.GetUTXOsFromAddress(*ownerUTXOAddress)
		if err != nil {
			log.Error(err)
		} else {
			DisplayUTXOS(utxos)
		}
	case getBalance.FullCommand():
		//plamsa_cli get balance --address=0x944A81BeECac91802787fBCFB9767FCBf81db1f5 --watcher=http://watcher.path.net
		chch, err := childchain.NewClient(*watcherURL, &http.Client{})
		if err != nil {
			log.Error(err)
		}
		balance, err := chch.GetBalance(*ownerUTXOAddress)
		if err != nil {
			log.Error(err)
		} else {
			DisplayBalance(balance)
		}
	case status.FullCommand():
		//plamsa_cli get status --watcher=http://watcher.path.net
		chch, err := childchain.NewClient(*watcherURL, &http.Client{})
		if err != nil {
			log.Error(err)
		}
		ws, err := chch.GetWatcherStatus()
		if err != nil {
			log.Error(err)
		} else {
			DisplayByzantineEvents(ws)
		}
	case deposit.FullCommand():
		//plasma_cli deposit --privatekey=0x944A81BeECac91802787fBCFB9767FCBf81db1f5 --client=https://rinkeby.infura.io/v3/api_key --contract=0x457e2ec4ad356d3cb449e3bd4ba640d720c30377 --currency=ETH
		// c := plasma.Deposit{PrivateKey: *privateKey, Client: *client, Contract: *contract, Amount: *depositAmount, Owner: util.DeriveAddress(*privateKey), Currency: *depositCurrency}
		// res, err := c.DepositEth()
		// if err != nil {
		// 	log.Error(err)
		// } else {
		// 	DisplayEthTransaction("Deposit successful, txhash: ", res)
		// }
		client, err := ethclient.Dial(*client)
		if err != nil {
			log.Errorf("error initializing Ethereum client for deposit, %v", err)
		}
		rc := rootchain.NewClient(client)

		depositTx := rc.NewDeposit(common.HexToAddress(*contract), common.HexToAddress(util.DeriveAddress(*privateKey)), common.HexToAddress(*depositCurrency), strconv.Itoa(*depositAmount))
		privatekey, err := crypto.HexToECDSA(util.FilterZeroX(*privateKey))
		if err != nil {
			log.Errorf("bad privatekey: %v", err)
		}
		gasPrice, _ := client.SuggestGasPrice(context.Background())
		txopts := bind.NewKeyedTransactor(privatekey)
		txopts.From = common.HexToAddress(util.DeriveAddress(*privateKey))
		txopts.GasLimit = 2000000
		txopts.Value = big.NewInt(int64(*depositAmount))
		txopts.GasPrice = gasPrice
		if err := rootchain.Options(depositTx, txopts); err != nil {
			log.Errorf("transaction options invalid, %v", err)
		}
		if err := rootchain.Build(depositTx); err != nil {
			log.Errorf("deposit build error, %v", err)
		}
		tx, err := rootchain.Submit(depositTx)
		if err != nil {log.Errorf("error submiting transaction for deposit =: %v", err)}

		log.Infof("%v", tx.Hash().Hex())

	case send.FullCommand():
		chch, err := childchain.NewClient(*watcherSubmitURL, &http.Client{})
		if err != nil {
			log.Errorf("unexpected error from creating new client: %v", err)
		}
		ptx := chch.NewPaymentTx()

		if err = ptx.AddOwner(util.DeriveAddress(*privatekey)); err != nil {
			log.Errorf("unexpected error from Adding owner: %v", err)
		}

		if err = ptx.AddPayment(*amount, *to, *currency); err != nil {
			log.Errorf("unexpected error from Adding payment: %v", err)
		}
		if err = ptx.AddMetadata(*metadata); err != nil {
			log.Errorf("unexpected error from Adding metadata: %v", err)
		}

		if err = ptx.AddFee(*feeamount, *feetoken); err != nil {
			log.Errorf("unexpected error from Adding metadata: %v", err)
		}

		if err = childchain.BuildTransaction(ptx); err != nil {
			log.Errorf("unexpected error : %v", err)
		}
		_, err = childchain.SignTransaction(ptx, childchain.SignWithRawKeys(*privatekey))
		if err != nil {
			log.Errorf("unexpected error : %v", err)
		}
		res, err := childchain.SubmitTransaction(ptx)
		if err != nil {
			log.Errorf("unexpected error : %v", err)
		} else {
			DisplaySubmitResponse(res)
		}

	case transaction.FullCommand():
		chch, err := childchain.NewClient(*watcherURL, &http.Client{})
		if err != nil {
			log.Error(err)
		}
		gtx, err := chch.GetTransaction(*txHash)
		if err != nil {
			log.Errorf("got error: %v", err)
		}
		DisplayGetResponse(gtx)
	case exit.FullCommand():
		//plasma_cli exit --utxo=1000000000 --privatekey=foo --contract=0x5bb7f2492487556e380e0bf960510277cdafd680 --watcher=ari.omg.network
		// s := plasma.StandardExit{UtxoPosition: util.ConvertStringToInt(*utxoPosition), Contract: *contractExit, PrivateKey: *exitPrivateKey, Client: *clientExit}
		// log.Info("Attempting to exit UTXO ", *utxoPosition)
		// s.StartStandardExit(*watcherExitURL)
		client, err := ethclient.Dial(*clientExit)
		rc := rootchain.NewClient(client)
		ste := rc.NewStandardExit(common.HexToAddress(*contractExit))
		c := &bind.CallOpts{}
		bondsize, err := ste.GetStandardExitBond(c)
		if err != nil {
			log.Errorf("could not get exit bond: %v", err)
		}
		gasPrice, _ := client.SuggestGasPrice(context.Background())
		privateKey, err := crypto.HexToECDSA(util.FilterZeroX(*exitPrivateKey))
		txopts := bind.NewKeyedTransactor(privateKey)
		txopts.From = common.HexToAddress(util.DeriveAddress(*exitPrivateKey))
		txopts.GasLimit = 2000000
		txopts.Value = bondsize
		txopts.GasPrice = gasPrice

		//fetches UTXO
		// ch, err := childchain.NewClient(*watcherExitURL, &http.Client{})
		// if err != nil {
		// 	log.Errorf("failed to start client: %v", err)
		// }
		// utxos, err := ch.GetUTXOsFromAddress(util.DeriveAddress(*exitPrivateKey))
		// if err != nil {
		// 	log.Errorf("error fetching utxos, %v", err)
		// }
		// fetches the first UTXO we find
		utxo, err := strconv.ParseInt(*utxoPosition, 10, 0)
		if err != nil {
			log.Errorf("issue parseing utxo position %v", err)
		}
		ed, err := rootchain.GetUTXOExitData(*watcherExitURL, int(utxo))
		if err != nil {
			log.Errorf("unexpected error from getting UTXO exit data %v", err)
		}

		if err := ste.Utxo(ed); err != nil {
			log.Errorf("error setting data for childchain transaction")
		}
		if err := rootchain.Options(ste, txopts); err != nil {
			log.Errorf("transaction options invalid, %v", err)
		}
		if err := rootchain.Build(ste); err != nil {
			log.Errorf("exit build error, %v", err)
		}
		tx, err := rootchain.Submit(ste)
		if err != nil {log.Errorf("error submiting transaction for exit =: %v", err)}

		res := tx.Hash().Hex()
		log.Printf("%v", res)
		log.Printf("standard exit tx hash: %s \n", res)

	case process.FullCommand():
		//plasma_cli process --contract=0x5bb7f2492487556e380e0bf960510277cdafd680 --token 0x0 --privatekey=foo --client=https://rinkeby.infura.io/v3/api_key
		// p := plasma.ProcessExit{Contract: *processContract, PrivateKey: *processPrivateKey, Token: *processToken, Client: *processExitClient}
		// log.Info("Calling process exits in the Plasma contract")
		// plasma.ProcessExits(1,1, p)
		client, _ := ethclient.Dial(*processExitClient)
		rc := rootchain.NewClient(client)
		vaultID := "1"
		numberOfExits := "1"
		topExit := "0"
		token := common.HexToAddress(util.EthCurrency)
		petx := rc.NewProcessExit(
			rootchain.PlasmaAddress(common.HexToAddress(*processContract)),
			rootchain.TokenAddress(token),
			rootchain.VaultId(vaultID),
			rootchain.TopExit(topExit),
			rootchain.NumberOfExits(numberOfExits),
		)
		gasPrice, _ := client.SuggestGasPrice(context.Background())
		privateKey, err := crypto.HexToECDSA(util.FilterZeroX(*processPrivateKey))
		txopts := bind.NewKeyedTransactor(privateKey)
		txopts.From = common.HexToAddress(util.DeriveAddress(*processPrivateKey))
		txopts.GasLimit = 2000000
		txopts.GasPrice = gasPrice

		if err := rootchain.Options(petx, txopts); err != nil {
			log.Errorf("transaction options invalid, %v", err)
		}

		if err := rootchain.Build(petx); err != nil {
			log.Errorf("error building process exit transaction: %v \n", err)
		}
		tx, err := rootchain.Submit(petx)
		if err != nil {log.Errorf("error submiting transaction for process exit =: %v", err)}
		res := tx.Hash().Hex()

		log.Printf("process exit tx hash: %s \n", res)

	case createAccount.FullCommand():
		//plasma_cli create account
		log.Info("Generating Keypair")
		util.GenerateAccount()
	case getExit.FullCommand():
		//plasma_cli get exit --watcher=https://watcher.ari.omg.network --utxo=1000000000000
		chch, err := childchain.NewClient(*watcherURL, &http.Client{})
		if err != nil {
			log.Error(err)
		}
		log.Info("Getting UTXO exit data")
		exitData, err := chch.GetUTXOExitData(*getExitUTXOPosition)
		if err != nil {
			log.Error(err)
		}
		log.Info("UTXO Position: ", exitData.Data.UtxoPos, " Proof: ", exitData.Data.Proof)
	// case challengeexit.FullCommand():
	// 	//plasma_cli challengeexit --contract="" --client="" --privatekey="" --utxo="" --watcher=""
	// 	challengeData, err := plasma.GetChallengeData(*ceWatcherURL, util.ConvertStringToInt(*ceUTXOPosition))
	// 	log.Info(challengeData)
	// 	if err != nil {
	// 		log.Fatalf("got error retrieving challenge data %v", err)
	// 	}
	// 	challengeData.ChallengeInvalidExit(*ceClient, *ceContract, *cePrivateKey)
	}
}
