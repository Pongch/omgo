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
	"net/http"
	"testing"
)

// Test getting status from watcher
func TestGetStatus(t *testing.T) {
	rs := `{ "version": "1.0", "success": true, "data": { "last_validated_child_block_timestamp": 1558535130, "last_validated_child_block_number": 10000, "last_mined_child_block_timestamp": 1558535190, "last_mined_child_block_number": 11000, "last_seen_eth_block_timestamp": 1558535190, "last_seen_eth_block_number": 4427041,"contract_addr": { "erc20_vault": "0x18e15c2cdc003b845b056f8d6b6a91ab33d3f182", "eth_vault": "0x895cc6f20d386f5c0deae08b08ccfec9f821e7d9", "payment_exit_game": "0x08c569c5774110eb84a80b292e6d6f039e18915a", "plasma_framework": "0x96d5d8bc539694e5fa1ec0dab0e6327ca9e680f9" }, "eth_syncing": true, "byzantine_events": [ { "event": "invalid_exit", "details": { "eth_height": 615440, "utxo_pos": 10000000010000000, "owner": "0xb3256026863eb6ae5b06fa396ab09069784ea8ea", "currency": "0x0000000000000000000000000000000000000000", "amount": 100 } }, { "event": "unchallenged_exit", "details": { "eth_height": 615440, "utxo_pos": 10000000010000000, "owner": "0xb3256026863eb6ae5b06fa396ab09069784ea8ea", "currency": "0x0000000000000000000000000000000000000000", "amount": 100 } } ], "services_synced_heights": [ { "service": "block_getter", "height": 4427041 }, { "service": "challenges_responds_processor", "height": 4427029 }, { "service": "competitor_processor", "height": 4427029 }, { "service": "convenience_deposit_processor", "height": 4427031 }, { "service": "convenience_exit_processor", "height": 4427029 }, { "service": "depositor", "height": 4427031 }, { "service": "exit_challenger", "height": 4427029 }, { "service": "exit_finalizer", "height": 4427029 }, { "service": "exit_processor", "height": 4427029 }, { "service": "ife_exit_finalizer", "height": 4427029 }, { "service": "in_flight_exit_processor", "height": 4427029 }, { "service": "piggyback_challenges_processor", "height": 4427029 }, { "service": "piggyback_processor", "height": 4427029 }, { "service": "root_chain_height", "height": 4427041 } ] } }`
	ms := createMockServer(t, rs, "/status.get", WatcherStatus{})
	defer ms.Close()
	chch, err := NewClient(ms.URL, &http.Client{})
	if err != nil {
		t.Errorf("unexpected error from creating new client: %v", err)
	}
	status, err := chch.GetWatcherStatus()
	if err != nil {
		t.Errorf("unexpected error from getting status: %v", err)
	}
	var ie int
	var ue int
	var uep interface{}
	var iep interface{}

	for _, v := range status.Data.ByzantineEvents {
		switch v.Event {
		case "unchallenged_exit":
			ue++
			uep = v.Details["utxo_pos"]
		case "invalid_exit":
			ie++
			iep = v.Details["utxo_pos"]
		}
	}

	if ie != 1 {
		t.Errorf("expected %v invalid exit, got %v", 1, ie)
	}

	if iep != 1.000000001e+16 {
		t.Errorf("expected %v UTXO_POS, got %v", 1.000000001e+16, iep)
	}

	if ue != 1 {
		t.Errorf("expected %v unchallenged exit, got %v", 1, ue)
	}

	if uep != 1.000000001e+16 {
		t.Errorf("expected %v UTXO_POS, got %v", 1.000000001e+16, uep)
	}
}
