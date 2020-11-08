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
	"net/http"

	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pongch/omgo/childchain"
	"github.com/pongch/omgo/util"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	send     = kingpin.Command("send", "Create a transaction on the OmiseGO Plasma MoreVP network")
	to       = send.Flag("to", "Wallet address of the recipient").Required().String()
	amount   = send.Flag("amount", "Amount to transact").Required().String()
	currency = send.Flag("currency", "currency of the amount to send, default to ETH").Default(childchain.EthCurrency).String()
	feetoken = send.Flag("feetoken", "set the token to be used as transaction fee, default to ETH").Default(childchain.EthCurrency).String()
	metadata = send.Flag("metadata", "additional metadata to send with the transaction, up to 32 bytes").Default(childchain.DefaultMetadata).String()
)

func _send() error {
	env, err := getCliEnv()
	if err != nil {
		return fmt.Errorf("error fetching environment variables: %v", err)
	}
	chch, err := childchain.NewClient(env.ccclient, &http.Client{})
	if err != nil {
		return fmt.Errorf("unexpected error from creating new client: %v", err)
	}
	ptx := chch.NewPaymentTx(
		childchain.AddOwner(common.HexToAddress(util.DeriveAddress(env.pkey))),
		childchain.AddPayment(*amount, common.HexToAddress(*to), common.HexToAddress(*currency)),
		childchain.AddMetadata(*metadata),
		childchain.AddFee(common.HexToAddress(*feetoken)),
	)
	if err = childchain.BuildTransaction(ptx); err != nil {
		return fmt.Errorf("unexpected error : %v", err)
	}
	_, err = childchain.SignTransaction(ptx, childchain.SignWithRawKeys(env.pkey))
	if err != nil {
		return fmt.Errorf("unexpected error : %v", err)
	}
	res, err := childchain.SubmitTransaction(ptx)
	if err != nil {
		return fmt.Errorf("unexpected error : %v", err)
	}
	displaySubmitResponse(res)
	return nil
}
