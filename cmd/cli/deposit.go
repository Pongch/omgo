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
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pongch/omgo/childchain"
	"github.com/pongch/omgo/rootchain"
	"github.com/pongch/omgo/util"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"math/big"
	"strconv"
)

var (
	deposit         = kingpin.Command("deposit", "Deposit ETH or ERC20 into the Plamsa MoreVP smart contract.")
	depositAmount   = deposit.Flag("amount", "Amount to deposit in wei").Required().Int()
	depositCurrency = deposit.Flag("currency", "ERC20 token address for the deposit, leave blank for ETH").Default(childchain.EthCurrency).String()
)

func _deposit() error {
	env, err := getCliEnv()
	if err != nil {
		return fmt.Errorf("error fetching environment variables:", err)
	}
	client, err := ethclient.Dial(env.rcclient)
	if err != nil {
		return fmt.Errorf("error initializing Ethereum client for deposit, %v", err)
	}
	rc := rootchain.NewClient(client)

	depositTx := rc.NewDeposit(common.HexToAddress(env.vcontract), common.HexToAddress(util.DeriveAddress(env.pkey)), common.HexToAddress(*depositCurrency), strconv.Itoa(*depositAmount))
	privatekey, err := crypto.HexToECDSA(util.FilterZeroX(env.pkey))
	if err != nil {
		return fmt.Errorf("bad privatekey: %v", err)
	}
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	txopts := bind.NewKeyedTransactor(privatekey)
	txopts.From = common.HexToAddress(util.DeriveAddress(env.pkey))
	txopts.GasLimit = 2000000
	txopts.Value = big.NewInt(int64(*depositAmount))
	txopts.GasPrice = gasPrice
	if err := rootchain.Options(depositTx, txopts); err != nil {
		return fmt.Errorf("transaction options invalid, %v", err)
	}
	if err := rootchain.Build(depositTx); err != nil {
		return fmt.Errorf("deposit build error, %v", err)
	}
	tx, err := rootchain.Submit(depositTx)
	if err != nil {
		return fmt.Errorf("error submiting transaction for deposit =: %v", err)
	} else {
		log.Infof("deposit txhash: %v", tx.Hash().Hex())
	}
	return nil

}
