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

package childchain

import (
	"encoding/json"
)

// WatcherStatus is a returned response from
// calling status.get
type ContractAddr struct {
	Erc20Vault      string `json:"erc20_vault"`
	EthVault        string `json:"eth_vault"`
	PaymentExitGame string `json:"payment_exit_game"`
	PlasmaFramework string `json:"plasma_framework"`
}

type WatcherStatus struct {
	Data struct {
		ByzantineEvents []struct {
			Event   string                 `json:"event"`
			Details map[string]interface{} `json:"details"`
		} `json:"byzantine_events"`
		ContractAddr  ContractAddr `json:"contract_addr"`
		EthSyncing    bool   `json:"eth_syncing"`
		InFlightExits []struct {
			EthHeight          json.Number           `json:"eth_height"`
			PiggybackedInputs  []interface{} `json:"piggybacked_inputs"`
			PiggybackedOutputs []interface{} `json:"piggybacked_outputs"`
			Txbytes            string        `json:"txbytes"`
			Txhash             string        `json:"txhash"`
		} `json:"in_flight_exits"`
		LastMinedChildBlockNumber        json.Number `json:"last_mined_child_block_number"`
		LastMinedChildBlockTimestamp     json.Number `json:"last_mined_child_block_timestamp"`
		LastSeenEthBlockNumber           json.Number `json:"last_seen_eth_block_number"`
		LastSeenEthBlockTimestamp        json.Number `json:"last_seen_eth_block_timestamp"`
		LastValidatedChildBlockNumber    json.Number `json:"last_validated_child_block_number"`
		LastValidatedChildBlockTimestamp json.Number `json:"last_validated_child_block_timestamp"`
		ServicesSyncedHeights            []struct {
			Height  json.Number    `json:"height"`
			Service string `json:"service"`
		} `json:"services_synced_heights"`
	} `json:"data"`
	Success bool   `json:"success"`
	Version string `json:"version"`
}

// GetWatcherStatus retrieves watcher status
// via status.get API
func (c *Client) GetWatcherStatus() (*WatcherStatus, error) {
	rstring, err := c.do(
		"/status.get",
		nil,
	)
	response := WatcherStatus{}
	if err != nil {
		return nil, err
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return &response, nil
}
