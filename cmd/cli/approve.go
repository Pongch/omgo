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
	"github.com/pongch/omgo/rootchain"
	"github.com/pongch/omgo/util"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"math/big"
)

var (
	approve         = kingpin.Command("approve", "approve erc20 token before deposit into the contract")
	approveCurrency = approve.Flag("currency", "address of ERC20 token").Required().String()
	approveAmount   = approve.Flag("amount", "address of ERC20 token").Required().String()
)

func _approve() error {
	env, err := getCliEnv()
	if err != nil {
		return fmt.Errorf("error fetching environment variables: %v", err)
	}
	client, err := ethclient.Dial(env.rcclient)
	if err != nil {
		return fmt.Errorf("error initializing Ethereum client for deposit, %v", err)
	}
	rc := rootchain.NewClient(client)
	privateKey, err := crypto.HexToECDSA(util.FilterZeroX(env.pkey))
	if err != nil {
		return fmt.Errorf("bad privatekey: %v", err)
	}
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	amount, _ := new(big.Int).SetString(*approveAmount, 0)
	approvalTx := rc.NewApprove(
		common.HexToAddress(env.vcontract),
		common.HexToAddress(*approveCurrency),
		amount,
	)
	gasPrice, _ = client.SuggestGasPrice(context.Background())
	txopts := bind.NewKeyedTransactor(privateKey)
	txopts.From = common.HexToAddress(util.DeriveAddress(env.pkey))
	txopts.GasLimit = 2000000
	txopts.GasPrice = gasPrice
	if err := rootchain.Options(approvalTx, txopts); err != nil {
		return fmt.Errorf("transaction options invalid, %v", err)
	}
	if err := rootchain.Build(approvalTx); err != nil {
		return fmt.Errorf("deposit build error, %v", err)
	}
	tx, err := rootchain.Submit(approvalTx)
	if err != nil {
		return fmt.Errorf("error submiting transaction for token approval =: %v", err)
	}
	log.Printf("txhash: %v \n", tx.Hash().Hex())
	return nil
}
