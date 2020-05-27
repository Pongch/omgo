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
	// "github.com/omisego/plasma-cli/rootchain"
	"fmt"
	"os"
	"strconv"
	"time"
	"context"
	"net/http"
	"math/big"
	"github.com/joho/godotenv"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
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
// TODO wrap all of the depositTx inside a rootchain wrapper interface function 
func TestDepositEthNew(t *testing.T) {
	env, err := loadTestEnv()
	if err != nil {
		t.Errorf("error loading test for deposit =: %v", err)
	}
	client, err := ethclient.Dial(env.EthClient)
	if err != nil {
		t.Errorf("error initializing Ethereum client for deposit =: %v", err)
	}
	rc := NewClient(client)

	depositTx := rc.NewDeposit(common.HexToAddress( env.EthVault ), common.HexToAddress(util.DeriveAddress(env.Privatekey)), common.HexToAddress(util.EthCurrency), strconv.Itoa(env.DepositAmount))
	privateKey, err := crypto.HexToECDSA(util.FilterZeroX(env.Privatekey))
	if err != nil {
		t.Errorf("bad privatekey: %v", err)
	}
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	txopts := bind.NewKeyedTransactor(privateKey)
	txopts.From = common.HexToAddress(util.DeriveAddress(env.Privatekey))
	txopts.GasLimit = 2000000
	txopts.Value = big.NewInt(int64(env.DepositAmount))
	txopts.GasPrice = gasPrice
	if err := Options(depositTx, txopts); err != nil {
		t.Errorf("transaction options invalid, %v", err)
	}
	if err := Build(depositTx); err != nil {
		t.Errorf("deposit build error, %v", err)
	}
	tx, err := Submit(depositTx)
	if err != nil {t.Errorf("error submiting transaction for deposit =: %v", err)}

	fmt.Printf("%v", tx.Hash().Hex())
	sleep(t)
	status := checkReceipt(tx.Hash().Hex(), t)
	if status == false {
		t.Error("transaction failed")
	}

}


func TestGetStandardExitBond(t *testing.T) {
	env, err := loadTestEnv()
	if err != nil {
		t.Errorf("error loading test env in get exit bond test: %v", err)
	}
	client, err := ethclient.Dial(env.EthClient)
	rc := NewClient(client)
	ste := rc.NewStandardExit(common.HexToAddress(env.ExitGame))
	c := &bind.CallOpts{}
	res, err := ste.GetStandardExitBond(c)
	if err != nil {
		t.Errorf("could not get exit bond: %v", err)
	}
	fmt.Printf("exit bond: %v \n", res)
}

func TestStartStandardExit(t *testing.T){
	env, err := loadTestEnv()
	if err != nil {
		t.Errorf("error loading test env in standard exit test: %v", err)
	}
	// get standard exit bond
	client, err := ethclient.Dial(env.EthClient)
	rc := NewClient(client)
	ste := rc.NewStandardExit(common.HexToAddress(env.ExitGame))
	c := &bind.CallOpts{}
	bondsize, err := ste.GetStandardExitBond(c)
	if err != nil {
		t.Errorf("could not get exit bond: %v", err)
	}
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	privateKey, err := crypto.HexToECDSA(util.FilterZeroX(env.Privatekey))
	txopts := bind.NewKeyedTransactor(privateKey)
	txopts.From = common.HexToAddress(util.DeriveAddress(env.Privatekey))
	txopts.GasLimit = 2000000
	txopts.Value = bondsize
	txopts.GasPrice = gasPrice

	//fetches UTXO
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
	ed, err := GetUTXOExitData(env.Watcher, int(utxo))
	if err != nil {
		t.Errorf("unexpected error from getting UTXO exit data %v", err)
	}

	if err := ste.Utxo(ed); err != nil {
		t.Errorf("error setting data for childchain transaction")
	}
	if err := Options(ste, txopts); err != nil {
		t.Errorf("transaction options invalid, %v", err)
	}
	if err := Build(ste); err != nil {
		t.Errorf("exit build error, %v", err)
	}
	tx, err := Submit(ste)
	if err != nil {t.Errorf("error submiting transaction for exit =: %v", err)}

	res := tx.Hash().Hex()
	fmt.Printf("%v", res)
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

	client, _ := ethclient.Dial(env.EthClient)
	rc := NewClient(client)
	vaultID := "1"
	numberOfExits := "1"
	topExit := "0"
	token := common.HexToAddress(util.EthCurrency)
	petx := rc.NewProcessExit(
		PlasmaAddress(common.HexToAddress(env.PlasmaFramework)),
		TokenAddress(token),
		VaultId(vaultID),
		TopExit(topExit),
		NumberOfExits(numberOfExits),
	)
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	privateKey, err := crypto.HexToECDSA(util.FilterZeroX(env.Privatekey))
	txopts := bind.NewKeyedTransactor(privateKey)
	txopts.From = common.HexToAddress(util.DeriveAddress(env.Privatekey))
	txopts.GasLimit = 2000000
	txopts.GasPrice = gasPrice

	if err := Options(petx, txopts); err != nil {
		t.Errorf("transaction options invalid, %v", err)
	}

	if err := Build(petx); err != nil {
		t.Errorf("error building process exit transaction: %v \n", err)
	}
	tx, err := Submit(petx)
	if err != nil {t.Errorf("error submiting transaction for process exit =: %v", err)}
	res := tx.Hash().Hex()

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

