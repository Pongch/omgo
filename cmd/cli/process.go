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
)

var (
	process         = kingpin.Command("process", "Process all exits that have completed the challenge period")
	processNumber   = process.Flag("number", "number of exits to process").Required().String()
	processVaultId  = process.Flag("id", "vault ID to process exit, 1 for ETH, 2 for ERC20 tokens").Default("1").String()
	processCurrency = process.Flag("currency", "currency to process exit from").Default(util.EthCurrency).String()
)

func _process() error {
	env, err := getCliEnv()
	if err != nil {
		return fmt.Errorf("error fetching environment variables: %v", err)
	}
	client, err := ethclient.Dial(env.rcclient)
	if err != nil {
		return fmt.Errorf("error dialing root chain client: %v", err)
	}
	rc := rootchain.NewClient(client)
	vaultID := *processVaultId
	numberOfExits := *processNumber
	topExit := "0"
	token := common.HexToAddress(*processCurrency)
	petx := rc.NewProcessExit(
		rootchain.PlasmaAddress(common.HexToAddress(env.fwcontract)),
		rootchain.TokenAddress(token),
		rootchain.VaultId(vaultID),
		rootchain.TopExit(topExit),
		rootchain.NumberOfExits(numberOfExits),
	)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("error getting gas price suggestion: %v", err)
	}
	privateKey, err := crypto.HexToECDSA(util.FilterZeroX(env.pkey))
	if err != nil {
		fmt.Errorf("error deriving privatekey: %v", err)
	}
	txopts := bind.NewKeyedTransactor(privateKey)
	txopts.From = common.HexToAddress(util.DeriveAddress(env.pkey))
	txopts.GasLimit = 2000000
	txopts.GasPrice = gasPrice

	if err := rootchain.Options(petx, txopts); err != nil {
		return fmt.Errorf("transaction options invalid, %v", err)
	}

	if err := rootchain.Build(petx); err != nil {
		return fmt.Errorf("error building process exit transaction: %v \n", err)
	}
	tx, err := rootchain.Submit(petx)
	if err != nil {
		return fmt.Errorf("error submiting transaction for process exit =: %v", err)
	}
	res := tx.Hash().Hex()
	log.Printf("process exit tx hash: %s \n", res)
	return nil
}
