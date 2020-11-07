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

package e2e

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
	"math/big"
	"net/http"
	"strconv"
	"testing"
)

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
	rc := rootchain.NewClient(client)

	depositTx := rc.NewDeposit(common.HexToAddress(env.EthVault), common.HexToAddress(util.DeriveAddress(env.Privatekey)), common.HexToAddress(util.EthCurrency), env.DepositAmount)
	privateKey, err := crypto.HexToECDSA(util.FilterZeroX(env.Privatekey))
	if err != nil {
		t.Errorf("bad privatekey: %v", err)
	}
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	txopts := bind.NewKeyedTransactor(privateKey)
	txopts.From = common.HexToAddress(util.DeriveAddress(env.Privatekey))
	txopts.GasLimit = 2000000
	bvalue, _ := new(big.Int).SetString(env.DepositAmount, 0)
	txopts.Value = bvalue
	txopts.GasPrice = gasPrice
	if err := rootchain.Options(depositTx, txopts); err != nil {
		t.Errorf("transaction options invalid, %v", err)
	}
	if err := rootchain.Build(depositTx); err != nil {
		t.Errorf("deposit build error, %v", err)
	}
	tx, err := rootchain.Submit(depositTx)
	if err != nil {
		t.Errorf("error submiting transaction for deposit =: %v", err)
	}

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
	rc := rootchain.NewClient(client)
	ste := rc.NewStandardExit(common.HexToAddress(env.ExitGame))
	c := &bind.CallOpts{}
	res, err := ste.GetStandardExitBond(c)
	if err != nil {
		t.Errorf("could not get exit bond: %v", err)
	}
	fmt.Printf("exit bond: %v \n", res)
}

func TestStartStandardExit(t *testing.T) {
	env, err := loadTestEnv()
	if err != nil {
		t.Errorf("error loading test env in standard exit test: %v", err)
	}
	// get standard exit bond
	client, err := ethclient.Dial(env.EthClient)
	rc := rootchain.NewClient(client)
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
	t.Errorf("%v", err)
	// fetches the first UTXO we find
	utxo, err := strconv.ParseInt(utxos.Data[0].UtxoPos.String(), 10, 0)
	if err != nil {
		t.Errorf("issue parseing utxo position %v", err)
	}
	ed, err := rootchain.GetUTXOExitData(env.Watcher, int(utxo))
	if err != nil {
		t.Errorf("unexpected error from getting UTXO exit data %v", err)
	}

	if err := ste.Utxo(ed); err != nil {
		t.Errorf("error setting data for childchain transaction")
	}
	if err := rootchain.Options(ste, txopts); err != nil {
		t.Errorf("transaction options invalid, %v", err)
	}
	if err := rootchain.Build(ste); err != nil {
		t.Errorf("exit build error, %v", err)
	}
	tx, err := rootchain.Submit(ste)
	if err != nil {
		t.Errorf("error submiting transaction for exit =: %v", err)
	}

	res := tx.Hash().Hex()
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
	rc := rootchain.NewClient(client)
	vaultID := "1"
	numberOfExits := "2"
	topExit := "0"
	token := common.HexToAddress(util.EthCurrency)
	petx := rc.NewProcessExit(
		rootchain.PlasmaAddress(common.HexToAddress(env.PlasmaFramework)),
		rootchain.TokenAddress(token),
		rootchain.VaultId(vaultID),
		rootchain.TopExit(topExit),
		rootchain.NumberOfExits(numberOfExits),
	)
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	privateKey, err := crypto.HexToECDSA(util.FilterZeroX(env.Privatekey))
	txopts := bind.NewKeyedTransactor(privateKey)
	txopts.From = common.HexToAddress(util.DeriveAddress(env.Privatekey))
	txopts.GasLimit = 2000000
	txopts.GasPrice = gasPrice

	if err := rootchain.Options(petx, txopts); err != nil {
		t.Errorf("transaction options invalid, %v", err)
	}

	if err := rootchain.Build(petx); err != nil {
		t.Errorf("error building process exit transaction: %v \n", err)
	}
	tx, err := rootchain.Submit(petx)
	if err != nil {
		t.Errorf("error submiting transaction for process exit =: %v", err)
	}
	res := tx.Hash().Hex()

	fmt.Printf("process exit tx hash: %s \n", res)
	sleep(t)
	status := checkReceipt(res, t)
	if status == false {
		t.Errorf("transaction failed")
	}
}

func TestErc20Approval(t *testing.T) {
	env, err := loadTestEnv()
	if err != nil {
		t.Errorf("error loading test env in standard exit test: %v", err)
	}
	client, _ := ethclient.Dial(env.EthClient)
	rc := rootchain.NewClient(client)
	privateKey, err := crypto.HexToECDSA(util.FilterZeroX(env.Privatekey))
	if err != nil {
		t.Errorf("bad privatekey: %v", err)
	}
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	amount, _ := new(big.Int).SetString(env.DepositAmount, 0)
	approvalTx := rc.NewApprove(
		common.HexToAddress(env.Erc20Vault),
		common.HexToAddress(env.Erc20token),
		amount,
	)
	gasPrice, _ = client.SuggestGasPrice(context.Background())
	txopts := bind.NewKeyedTransactor(privateKey)
	txopts.From = common.HexToAddress(util.DeriveAddress(env.Privatekey))
	txopts.GasLimit = 2000000
	txopts.GasPrice = gasPrice
	if err := rootchain.Options(approvalTx, txopts); err != nil {
		t.Errorf("transaction options invalid, %v", err)
	}
	if err := rootchain.Build(approvalTx); err != nil {
		t.Errorf("deposit build error, %v", err)
	}
	tx, err := rootchain.Submit(approvalTx)
	if err != nil {
		t.Errorf("error submiting transaction for token approval =: %v", err)
	}
	fmt.Printf("%v", tx.Hash().Hex())
	sleep(t)
	status := checkReceipt(tx.Hash().Hex(), t)
	if status == false {
		t.Error("transaction failed")
	}
}

func TestPaymentTransaction(t *testing.T) {
	env, err := loadTestEnv()
	if err != nil {
		t.Errorf("error loading test env in standard exit test: %v", err)
	}
	chch, err := childchain.NewClient(env.Watcher, &http.Client{})
	if err != nil {
		t.Errorf("unexpected error from creating new client: %v", err)
	}
	ptx := chch.NewPaymentTx(
		childchain.AddOwner(common.HexToAddress(util.DeriveAddress(env.Privatekey))),
		childchain.AddPayment(env.PaymentAmount, common.HexToAddress("0xedcf990e493f271020f3a2b2d6f17962683b2c45"), common.HexToAddress(childchain.EthCurrency)),
		childchain.AddMetadata(childchain.DefaultMetadata),
		childchain.AddFee(common.HexToAddress(childchain.EthCurrency)),
	)

	if err = childchain.BuildTransaction(ptx); err != nil {
		t.Errorf("unexpected error : %v", err)
	}
	_, err = childchain.SignTransaction(ptx, childchain.SignWithRawKeys(env.Privatekey))
	if err != nil {
		t.Errorf("unexpected error : %v", err)
	}
	res, err := childchain.SubmitTransaction(ptx)
	if err != nil {
		t.Errorf("unexpected error : %v", err)
	}
	fmt.Printf("txhash: %v \n", res)
}
