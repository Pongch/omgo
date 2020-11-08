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
	"strconv"
)

var (
	exit         = kingpin.Command("exit", "Standard exit a UTXO back to the root chain.")
	utxoPosition = exit.Flag("utxopos", "UTXO Position that will be exited").Required().String()
)

func _exit() error {
	env, err := getCliEnv()
	if err != nil {
		return fmt.Errorf("error fetching environment variables: %v", err)
	}
	client, err := ethclient.Dial(env.rcclient)
	if err != nil {
		fmt.Errorf("error dialing ethclient: %v", err)
	}
	rc := rootchain.NewClient(client)
	ste := rc.NewStandardExit(common.HexToAddress(env.econtract))
	c := &bind.CallOpts{}
	bondsize, err := ste.GetStandardExitBond(c)
	if err != nil {
		return fmt.Errorf("could not get exit bond: %v", err)
	}
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	privateKey, err := crypto.HexToECDSA(util.FilterZeroX(env.pkey))
	txopts := bind.NewKeyedTransactor(privateKey)
	txopts.From = common.HexToAddress(util.DeriveAddress(env.pkey))
	txopts.GasLimit = 2000000
	txopts.Value = bondsize
	txopts.GasPrice = gasPrice

	utxo, err := strconv.ParseInt(*utxoPosition, 10, 0) //TODO handle big num
	if err != nil {
		return fmt.Errorf("issue parseing utxo position %v", err)
	}
	ed, err := rootchain.GetUTXOExitData(env.ccclient, int(utxo))
	if err != nil {
		return fmt.Errorf("unexpected error from getting UTXO exit data %v", err)
	}

	if err := ste.Utxo(ed); err != nil {
		return fmt.Errorf("error setting data for childchain transaction: %v", err)
	}
	if err := rootchain.Options(ste, txopts); err != nil {
		return fmt.Errorf("transaction options invalid, %v", err)
	}
	if err := rootchain.Build(ste); err != nil {
		return fmt.Errorf("exit build error, %v", err)
	}
	tx, err := rootchain.Submit(ste)
	if err != nil {
		return fmt.Errorf("error submiting transaction for exit: %v", err)
	}
	log.Printf("standard exit tx hash: %s \n", tx.Hash().Hex())
	return nil
}
