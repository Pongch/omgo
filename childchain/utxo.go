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
	"math/big"
)

// StandardExitUTXOData is a response from
// watcher from /utxo.get_exit_data endpoint
type StandardExitUTXOData struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		UtxoPos *big.Int `json:"utxo_pos"`
		Txbytes string   `json:"txbytes"`
		Sigs    string   `json:"sigs"`
		Proof   string   `json:"proof"`
	} `json:"data"`
}

// ChallengeUTXOData is a response from
// watcher from /utxo.get_challenge_data endpoint
type ChallengeUTXOData struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		ExitId     *big.Int `json:"exit_id"`
		InputIndex uint8    `json:"input_index"`
		Sig        string   `json:"sig"`
		Txbytes    string   `json:"txbytes"`
	} `json:"data"`
}

// GetUTXOExitData from watcher based on utxo position
// from /utxo.get_exit_data endpoint
// to be used in rootchain exit function call
func (c *Client) GetUTXOExitData(utxoPosition int) (*StandardExitUTXOData, error) {
	postData := map[string]interface{}{"utxo_pos": utxoPosition}
	rstring, err := c.do(
		"/utxo.get_exit_data",
		postData,
	)
	response := StandardExitUTXOData{}
	if err != nil {
		return nil, err
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return &response, nil
}

// GetChallengeData from the watcher based on
// UTXO position to be used in rootchain Challenge
// function call
func (c *Client) GetChallengeData(utxoPosition *big.Int) (*ChallengeUTXOData, error) {
	postData := map[string]interface{}{"utxo_pos": utxoPosition}
	rstring, err := c.do(
		"/utxo.get_challenge_data",
		postData,
	)
	response := ChallengeUTXOData{}
	if err != nil {
		return nil, err
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return &response, nil
}
