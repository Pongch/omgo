//
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

package cli

import (
	"net/http"

	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pongch/omgo/childchain"
	"github.com/pongch/omgo/rootchain"
	"github.com/pongch/omgo/util"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"strconv"
)

var (
	// get              = kingpin.Command("get", "Get a resource.").Default()
	// getUTXO          = get.Command("utxos", "Retrieve UTXO data from the Watcher service")
	watcherURL = get.Flag("watcher", "FQDN of the Watcher in the format https://watcher.path.net").String()
	// ownerUTXOAddress = get.Flag("address", "Owner address to search UTXOs").String()
	// getBalance       = get.Command("balance", "Retrieve balance of an address from the Watcher service")
	// status           = get.Command("status", "Get status from the Watcher")
	// transaction      = get.Command("transaction", "get transaction details")
	// txHash           = get.Flag("txhash", "transaction hash of the transaction you would like to get the information for").String()

	getExit             = get.Command("exit", "Get UTXO exit information")
	getExitUTXOPosition = getExit.Flag("utxo", "Get UTXO exit information").Required().Int()

	// deposit         = kingpin.Command("deposit", "Deposit ETH or ERC20 into the Plamsa MoreVP smart contract.")
	// privateKey      = deposit.Flag("privatekey", "Private key of the account used to send funds into Plasma MoreVP").Required().String()
	// client          = deposit.Flag("client", "Address of the Ethereum client. Infura and local node supported https://rinkeby.infura.io/v3/api_key or http://localhost:8545").Required().String()
	// contract        = deposit.Flag("contract", "Address of the Plasma MoreVP smart contract").Required().String()
	// depositAmount   = deposit.Flag("amount", "Amount to deposit in wei").Required().Int()
	// depositCurrency = deposit.Flag("currency", "ERC20 token address for the deposit, leave blank for ETH").Default(childchain.EthCurrency).String()

	send             = kingpin.Command("send", "Create a transaction on the OmiseGO Plasma MoreVP network")
	to               = send.Flag("to", "Wallet address of the recipient").Required().String()
	privatekey       = send.Flag("privatekey", "Privatekey from a wallet to send from").Required().String()
	amount           = send.Flag("amount", "Amount to transact").Required().Uint64()
	currency         = send.Flag("currency", "currency of the amount to send, default to ETH").Default(childchain.EthCurrency).String()
	watcherSubmitURL = send.Flag("watcher", "FQDN of the Watcher in the format http://watcher.path.net").Required().String()
	feetoken         = send.Flag("feetoken", "set the token to be used as transaction fee, default to ETH").Default(childchain.EthCurrency).String()
	metadata         = send.Flag("metadata", "additional metadata to send with the transaction, up to 32 bytes").Default(childchain.DefaultMetadata).String()

	exit           = kingpin.Command("exit", "Standard exit a UTXO back to the root chain.")
	watcherExitURL = exit.Flag("watcher", "FQDN of the Watcher in the format http://watcher.path.net").Required().String()
	contractExit   = exit.Flag("contract", "Address of the Plasma MoreVP smart contract").Required().String()
	utxoPosition   = exit.Flag("utxo", "UTXO Position that will be exited").Required().String()
	exitPrivateKey = exit.Flag("privatekey", "Private key of the UTXO owner").Required().String()
	clientExit     = exit.Flag("client", "Address of the Ethereum client. Infura and local node supported https://rinkeby.infura.io/v3/api_key or http://localhost:8545").Required().String()

	process         = kingpin.Command("process", "Process all exits that have completed the challenge period")
	processContract = process.Flag("contract", "Address of the Plasma MoreVP smart contract").Required().String()
	// processToken      = process.Flag("token", "Token address to process for standard exits").Required().String()
	processPrivateKey = process.Flag("privatekey", "Private key used to fund the gas for the smart contract call").Required().String()
	processExitClient = process.Flag("client", "Address of the Ethereum client. Infura and local node supported https://rinkeby.infura.io/v3/api_key or http://localhost:8545").Required().String()

	create        = kingpin.Command("create", "Create a resource.")
	createAccount = create.Command("account", "Create an account consisting of Public and Private key")
)

func ParseArgs() {
	switch kingpin.Parse() {
	case getUtxo.FullCommand():
		if err := _getUtxo(); err != nil {
			log.Error(err)
		}
	case getBalance.FullCommand():
		if err := _getBalance(); err != nil {
			log.Error(err)
		}
	case getStatus.FullCommand():
		if err := _getStatus(); err != nil {
			log.Error(err)
		}
	case deposit.FullCommand():
		if err := _deposit(); err != nil {
			log.Error(err)
		}
	case send.FullCommand():
		chch, err := childchain.NewClient(*watcherSubmitURL, &http.Client{})
		if err != nil {
			log.Errorf("unexpected error from creating new client: %v", err)
		}
		ptx := chch.NewPaymentTx(
			childchain.AddOwner(common.HexToAddress(util.DeriveAddress(*privatekey))),
			childchain.AddPayment(string(*amount), common.HexToAddress(*to), common.HexToAddress(*currency)),
			childchain.AddMetadata(*metadata),
			childchain.AddFee(common.HexToAddress(*feetoken)),
		)
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

	case getTransaction.FullCommand():
		if err := _getTransaction(); err != nil {
			log.Error(err)
		}
	case exit.FullCommand():
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
		if err != nil {
			log.Errorf("error submiting transaction for exit =: %v", err)
		} else {
			res := tx.Hash().Hex()
			log.Printf("%v", res)
			log.Printf("standard exit tx hash: %s \n", res)
		}

	case process.FullCommand():
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
		if err != nil {
			log.Errorf("error submiting transaction for process exit =: %v", err)
		} else {
			res := tx.Hash().Hex()
			log.Printf("process exit tx hash: %s \n", res)
		}
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
		} else {
			log.Info("UTXO Position: ", exitData.Data.UtxoPos, " Proof: ", exitData.Data.Proof)
		}
	}
}
