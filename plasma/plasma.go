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
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/omisego/plasma-cli/rootchain"
	"github.com/omisego/plasma-cli/util"
	log "github.com/sirupsen/logrus"
)


type Deposit struct {
	PrivateKey string
	Client     string
	Contract   string
	Amount     uint64
	Owner      string
	Currency   string
}

type blockNumber struct {
	Hash string `json:"hash"`
}

type blockNumberError struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		Object   string `json:"object"`
		Messages struct {
			ValidationError struct {
				Validator string `json:"validator"`
				Parameter string `json:"parameter"`
			} `json:"validation_error"`
		} `json:"messages"`
		Description string `json:"description"`
		Code        string `json:"code"`
	} `json:"data"`
}

type ProcessExit struct {
	Contract   string
	PrivateKey string
	Token      string
	Client     string
}

type StandardExit struct {
	UtxoPosition int
	PrivateKey   string
	Contract     string
	Client       string
}

type standardExitUTXOError struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		Object      string `json:"object"`
		Code        string `json:"code"`
		Description string `json:"description"`
		Messages    struct {
			ErrorKey string `json:"error_key"`
		} `json:"messages"`
	} `json:"data"`
}

type StandardExitUTXOData struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		UtxoPos *big.Int `json:"utxo_pos"`
		Txbytes string   `json:"txbytes"`
		Proof   string   `json:"proof"`
	} `json:"data"`
}

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

type watcherError struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		Object      string `json:"object"`
		Description string `json:"description"`
		Code        string `json:"code"`
	} `json:"data"`
}


// Start a standard exit from user provided UTXO & private key
func (s *StandardExit) StartStandardExit(watcher string) {
	log.Info("Getting data needed to exit the UTXO from the Watcher")
	exit, err := GetUTXOExitData(watcher, s.UtxoPosition)
	if err != nil {
		log.Fatal(err)
	}
	bond, err := GetStandardExitBond(s.Client, s.Contract, s.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	exit.StartStandardExit(s.Client, s.Contract, s.PrivateKey, bond)
}


//Retrieve the UTXO exit data from the UTXO position
func GetUTXOExitData(watcher string, utxoPosition int) (StandardExitUTXOData, error) {
	// Build request
	var url strings.Builder
	url.WriteString(watcher)
	url.WriteString("/utxo.get_exit_data")
	postData := map[string]interface{}{"utxo_pos": utxoPosition}
	js, _ := json.Marshal(postData)
	r, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(js))
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Unmarshall the response
	response := StandardExitUTXOData{}

	rstring, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		return response, errors.New("No exit data found for UTXO provided")
	}

	return response, nil
}

// Start standard exit by calling the method in the smart contract
func (s *StandardExitUTXOData) StartStandardExit(ethereumClient , contract , private string, bondsize *big.Int) (string, error) {
	client, err := ethclient.Dial(ethereumClient)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(util.FilterZeroX(private))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(31415926535) // in wei
	auth.GasLimit = uint64(210000)       // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(contract)
	instance, err := rootchain.NewPaymentExitGame(address, client)
	if err != nil {
		log.Fatal(err)
	}
	t := &bind.TransactOpts{}
	t.From = fromAddress
	t.Signer = auth.Signer
	t.Value = bondsize //STANDARD_EXIT_BOND in the smart contract
	t.GasLimit = 2000000

	txBytesHex, txErr := hex.DecodeString(util.RemoveLeadingZeroX(s.Data.Txbytes))
	if txErr != nil {
		log.Fatal(txErr)
	}

	proofBytesHex, proofErr := hex.DecodeString(util.RemoveLeadingZeroX(s.Data.Proof))
	if proofErr != nil {
		log.Fatal(proofErr)
	}
	seargs := rootchain.PaymentStandardExitRouterArgsStartStandardExitArgs{
		UtxoPos: s.Data.UtxoPos,
		RlpOutputTx: []byte(txBytesHex),
		OutputTxInclusionProof: []byte(proofBytesHex),
	}
	tx, err := instance.StartStandardExit(t, seargs) 
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil 
}

// get standard exit bond 
func GetStandardExitBond(rootchainclient, contract, pkey string) (*big.Int, error){
client, err := ethclient.Dial(rootchainclient)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(util.FilterZeroX(pkey))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(6000000)           // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(contract)
	instance, err := rootchain.NewPaymentExitGame(address, client)
	if err != nil {
		log.Fatal(err)
	}
	c := &bind.CallOpts{}
	amount, err := instance.StartStandardExitBondSize(c)
	if err != nil {
		return nil, err
	} 
	return	amount, nil

}

// Deposit ETH into Plasma MoreVP ALD contract 
func (d *Deposit) DepositEth() (string, error){
	client, err := ethclient.Dial(d.Client)
	if err != nil {
		return "", err
	}

	privateKey, err := crypto.HexToECDSA(util.FilterZeroX(d.PrivateKey))
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		errors.New("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err 
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))

	auth.Value = big.NewInt(int64(d.Amount)) // in wei
	auth.GasLimit = uint64(6000000)           // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(d.Contract)
	rlpInputs, err := util.BuildRLPDeposit(util.RemoveLeadingZeroX(d.Owner), d.Currency, d.Amount, 1)
	if err != nil {
		return "", err 
	}
	instance, err := rootchain.NewEthvault(address, client)
	if err != nil {
		return "", err
	}
	t := &bind.TransactOpts{}
	t.From = fromAddress
	t.Signer = auth.Signer
	t.Value = big.NewInt(int64(d.Amount))
	tx, err := instance.Deposit(t, rlpInputs)
	if err != nil {
		return "", err
	}
	return	tx.Hash().Hex(), nil

}


// Calls the processExits in the Plasma smart contract to start processing exits that
// have completed the challenge period.
func ProcessExits(vaultID, numberExitsToProcess int64, p ProcessExit) (string, error) {
	client, err := ethclient.Dial(p.Client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(util.FilterZeroX(p.PrivateKey))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(210000) // in units
	auth.GasPrice = gasPrice

	contractAddress := p.Contract

	address := common.HexToAddress(contractAddress)
	instance, err := rootchain.NewPlasmaFramework(address, client)
	if err != nil {
		return "", err
	}
	t := &bind.TransactOpts{}
	t.From = fromAddress
	t.Signer = auth.Signer
	t.Value = big.NewInt(0)
	t.GasLimit = 2000000

	token := common.HexToAddress(p.Token)

	tx, err := instance.ProcessExits(t, big.NewInt(vaultID), token, big.NewInt(0), big.NewInt(numberExitsToProcess))
	if err != nil {
		return "", err
	} else {
		return tx.Hash().Hex(), nil
	}
}

//get challenge data from watcher
func GetChallengeData(watcher string, utxoPosition int) (ChallengeUTXOData, error) {
	// Build request
	var url strings.Builder
	url.WriteString(watcher)
	url.WriteString("/utxo.get_challenge_data")
	postData := map[string]interface{}{"utxo_pos": utxoPosition}
	js, _ := json.Marshal(postData)
	r, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(js))
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Unmarshall the response
	response := ChallengeUTXOData{}

	rstring, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonErr := json.Unmarshal([]byte(rstring), &response)
	if jsonErr != nil {
		return response, errors.New("No challenge data found for UTXO provided")
	} else {
		log.Info(resp.Status)
	}

	return response, nil
}

//challenge invalid exit on the root chain
func (c *ChallengeUTXOData) ChallengeInvalidExit(ethereumClient string, contract string, private string) {
	client, err := ethclient.Dial(ethereumClient)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(util.FilterZeroX(private))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice

	address := common.HexToAddress(contract)
	instance, err := rootchain.NewRootchain(address, client)
	if err != nil {
		log.Fatal(err)
	}
	t := &bind.TransactOpts{}
	t.From = fromAddress
	t.Signer = auth.Signer
	t.GasLimit = 2000000

	txBytesHex, txErr := hex.DecodeString(util.RemoveLeadingZeroX(c.Data.Txbytes))
	if txErr != nil {
		log.Fatal(txErr)
	}

	sigBytesHex, bytesErr := hex.DecodeString(util.RemoveLeadingZeroX(c.Data.Sig))
	if bytesErr != nil {
		log.Fatal(bytesErr)
	}
	tx, err := instance.ChallengeStandardExit(t, c.Data.ExitId, []byte(txBytesHex), c.Data.InputIndex, []byte(sigBytesHex))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Challenge exit to Plasma MoreVP sent. Transaction: ", tx.Hash().Hex())
	}
}
