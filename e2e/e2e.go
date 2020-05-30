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

package e2e

import (
	"testing"
	"fmt"
	"os"
	"strconv"
	"time"
	"context"
	"github.com/joho/godotenv"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
)

type TestEnv struct {
	Watcher, Privatekey, Publickey, EthClient, EthVault, ExitGame, PlasmaFramework string
	UtxoPos, DepositAmount, PaymentAmount, ExitToProcess, BlockTime, BlockConfirmation int

}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := os.Getenv(name)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func loadTestEnv() (*TestEnv, error){
	err := godotenv.Load("./e2e.env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file in test: %v", err)
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
		PaymentAmount: getEnvAsInt("PAYMENT_AMOUNT",1),
		BlockTime: getEnvAsInt("BLOCK_TIME", 12),
		BlockConfirmation: getEnvAsInt("BLOCK_CONFIRMATION", 7),
		ExitToProcess: getEnvAsInt("EXIT_TO_PROCESS",1),
	}
	return &env, nil
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
