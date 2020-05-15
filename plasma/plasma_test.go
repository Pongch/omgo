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

package plasma 

import (
	"testing"
	"github.com/omisego/plasma-cli/util"
	"github.com/omisego/plasma-cli/childchain"
	"fmt"
	"os"
	"strconv"
	"time"
	"context"
	"net/http"
	// "math/big"
	"github.com/joho/godotenv"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
)

type TestEnv struct {
	Watcher, Privatekey, Publickey, EthClient, EthVault, ExitGame, PlasmaFramework string
	UtxoPos, DepositAmount, ExitToProcess, BlockTime, BlockConfirmation int

}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := os.Getenv(name)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func loadTestEnv() (*TestEnv, error){
	err := godotenv.Load("./integration_test.env")
	if err != nil {
		return nil, fmt.Errorf("error loading env in test: %v", err)
	}
	env := TestEnv{
		Watcher: os.Getenv("WATCHER"),
		Privatekey: os.Getenv("PRIVKEY"),
		Publickey: os.Getenv("PUBKEY"),
		EthClient: os.Getenv("ETH_CLIENT"),
		EthVault: os.Getenv("ETH_VAULT_CONTRACT"),
		ExitGame: os.Getenv("EXIT_GAME_CONTRACT"),
		PlasmaFramework: os.Getenv("PLASMA_FRAMEWORK_CONTRACT"),
		UtxoPos: getEnvAsInt("UTXO_POS_FOR_EXIT", 1),
		DepositAmount: getEnvAsInt("DEPOSIT_AMOUNT",1),
		BlockTime: getEnvAsInt("BLOCK_TIME", 12),
		BlockConfirmation: getEnvAsInt("BLOCK_CONFIRMATION", 7),
		ExitToProcess: getEnvAsInt("EXIT_TO_PROCESS",1),
	}
	return &env, nil
}

// must start new ETH Client, then take in a signer to sign transactions
// rc, err := rc.Client(//rootchain interface)
// depositTx := rc.NewDeposit(// args)
// res, err := rc.Sign(depositTx, util.SignWithRawKeys(//pkey))
// txhash, err := rc.SubmitTransaction(depositTx) // the validation done before sending transaction


func TestDepositEth(t *testing.T) {
	env, err := loadTestEnv()
	if err != nil {
		t.Errorf("error loading test env in standard exit test: %v", err)
	}

	d := Deposit{PrivateKey: env.Privatekey, Client: env.EthClient, Contract: env.EthVault, Amount: uint64(env.DepositAmount), Owner: util.DeriveAddress(env.Privatekey), Currency: util.EthCurrency}
	res,  err := d.DepositEth()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("deposit tx hash: %s \n", res)
	sleep(t)
	status := checkReceipt(res, t)
	if status == false {
		t.Error("transaction failed")
	}
}

func TestGetStandardExitBond(t *testing.T) {
	env, err := loadTestEnv()
	if err != nil {
		t.Errorf("error loading test env in get exit bond test: %v", err)
	}
	res, err := GetStandardExitBond(env.EthClient, env.ExitGame, env.Privatekey)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	fmt.Printf("exit bond value: %v", res)
}

func TestStartStandardExit(t *testing.T){
	env, err := loadTestEnv()
	if err != nil {
		t.Errorf("error loading test env in standard exit test: %v", err)
	}
	// get standard exit bond
	bond, err := GetStandardExitBond(env.EthClient, env.ExitGame, env.Privatekey)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	// fetches all UTXOs available 
	ch, err := childchain.NewClient(env.Watcher, &http.Client{})
	if err != nil {
		t.Errorf("failed to start client: %v", err)
	}
	utxos, err := ch.GetUTXOsFromAddress(env.Publickey)
	if err != nil {
		t.Errorf("error fetching utxos, %v", err)
	}
	// fetches the first UTXO we find
	utxo, err := strconv.ParseInt(utxos.Data[0].UtxoPos.String(), 10, 0)
	if err != nil {
		t.Errorf("issue parseing utxo position %v", err)
	}
	exit, err := GetUTXOExitData(env.Watcher, int(utxo))
	if err != nil {
t.Errorf("unexpected error from getting UTXO exit data %v", err)
	}
	// call start standard exit 
	res, err := exit.StartStandardExit(env.EthClient, env.ExitGame, env.Privatekey, bond)
	if err != nil {
t.Errorf("unexpected error from, starting exit: %v", err)
	}
	fmt.Printf("standard exit tx hash: %s \n", res)
	sleep(t)
	status := checkReceipt(res, t)
	if status == false {
		t.Errorf("transaction failed")
	}
}


func TestProcessExit(t *testing.T) {
	env, err := loadTestEnv()
	if err != nil {
		t.Errorf("error loading test env in standard exit test: %v", err)
	}

	p := ProcessExit{Contract: env.PlasmaFramework, PrivateKey: env.Privatekey, Token: util.EthCurrency, Client: env.EthClient}
	res, err := ProcessExits(1, int64(env.ExitToProcess), p)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("process exit tx hash: %s \n", res)
	sleep(t)
	status := checkReceipt(res, t)
	if status == false {
		t.Errorf("transaction failed")
	}
}

// wait for 2 Ethereum block
func sleep(t *testing.T) {
	env, err := loadTestEnv()
	if err != nil {
		t.Errorf("error while sleep: %v", err)
	}
	time.Sleep(time.Duration(env.BlockTime * env.BlockConfirmation) * time.Second)
}

// get the result of transaction, return success or failure 
func checkReceipt(receipt string, t *testing.T) bool {
	env, err := loadTestEnv()
	if err != nil {
		t.Errorf("error loading env. %v ", err)
	}
	client, err := ethclient.Dial(env.EthClient) 
	if err != nil {
		t.Errorf("bad client: %v", err)
	}
	hash := common.HexToHash(receipt)
	tx, err := client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		t.Errorf("bad tx receipt: %v", err)
	}
	if tx.Status == 0 {
		return false
	} else {
		return true
	}
}

