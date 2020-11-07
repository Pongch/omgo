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
	"fmt"
	"github.com/pongch/omgo/childchain"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
)

var (
	get            = kingpin.Command("get", "Get a resource.").Default()
	getUtxo        = get.Command("utxos", "Retrieve UTXO data from the Watcher service")
	getBalance     = get.Command("balance", "Retrieve balance of an address from the Watcher service")
	getStatus      = get.Command("status", "Get status from the Watcher")
	getTransaction = get.Command("transaction", "get transaction details")
	from           = get.Flag("from", "address to search UTXOs").String()
	txHash         = get.Flag("txhash", "transaction hash of the transaction you would like to get the information for").String()
)

func _getUtxo() error {
	env, err := getCliEnv()
	if err != nil {
		return fmt.Errorf("error fetching environment variables: %v", err)
	}
	chch, err := childchain.NewClient(env.ccclient, &http.Client{})
	if err != nil {
		return fmt.Errorf("error initializing child chain client: %v", err)
	}
	utxos, err := chch.GetUTXOsFromAddress(*from)
	if err != nil {
		return fmt.Errorf("error fetching utxos: %v", err)
	}
	DisplayUTXOS(utxos)
	return nil
}

func _getBalance() error {
	env, err := getCliEnv()
	if err != nil {
		return fmt.Errorf("error fetching environment variables: %v", err)
	}
	chch, err := childchain.NewClient(env.ccclient, &http.Client{})
	if err != nil {
		return fmt.Errorf("error initializing child chain client: %v", err)
	}
	balance, err := chch.GetBalance(*from)
	if err != nil {
		return fmt.Errorf("error fetching balance: %v", err)
	}
	DisplayBalance(balance)
	return nil
}

func _getStatus() error {
	env, err := getCliEnv()
	if err != nil {
		return fmt.Errorf("error fetching environment variables: %v", err)
	}
	chch, err := childchain.NewClient(env.ccclient, &http.Client{})
	if err != nil {
		return fmt.Errorf("error initializing child chain client: %v", err)
	}
	ws, err := chch.GetWatcherStatus()
	if err != nil {
		return fmt.Errorf("error fetching watcher status", err)
	}
	DisplayByzantineEvents(ws)
	return nil
}

func _getTransaction() error {
	env, err := getCliEnv()
	if err != nil {
		return fmt.Errorf("error fetching environment variables: %v", err)
	}
	chch, err := childchain.NewClient(env.ccclient, &http.Client{})
	if err != nil {
		return fmt.Errorf("error initializing child chain client: %v", err)
	}
	ws, err := chch.GetTransaction(*txHash)
	if err != nil {
		return fmt.Errorf("error fetching transaction", err)
	}
	DisplayGetResponse(ws)
	return nil
}
