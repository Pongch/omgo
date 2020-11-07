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
	"os"
)

type clienv struct {
	ccclient   string
	rcclient   string
	fwcontract string
	vcontract  string
	econtract  string
	pkey       string
}

func getCliEnv() (*clienv, error) {
	var ccclient string
	var rcclient string
	var vcontract string
	var fwcontract string
	var econtract string
	var pkey string
	var ok bool
	if ccclient, ok = os.LookupEnv("OMGO_CCCLIENT"); !ok {
		return nil, fmt.Errorf("env OMGO_CCCLIENT not found")
	}
	if rcclient, ok = os.LookupEnv("OMGO_RCCLIENT"); !ok {
		return nil, fmt.Errorf("env OMGO_RCCLIENT not found")
	}
	if fwcontract, ok = os.LookupEnv("OMGO_FWCONTRACT"); !ok {
		return nil, fmt.Errorf("env OMGO_FWCONTRACT not found")
	}
	if vcontract, ok = os.LookupEnv("OMGO_VCONTRACT"); !ok {
		return nil, fmt.Errorf("env OMGO_VCONTRACT not found")
	}
	if econtract, ok = os.LookupEnv("OMGO_ECONTRACT"); !ok {
		return nil, fmt.Errorf("env OMGO_ECONTRACT not found")
	}
	if pkey, ok = os.LookupEnv("OMGO_PKEY"); !ok {
		return nil, fmt.Errorf("env OMGO_PKEY not found")
	}
	clienv := clienv{
		ccclient:   ccclient,
		rcclient:   rcclient,
		fwcontract: fwcontract,
		vcontract:  vcontract,
		econtract:  econtract,
		pkey:       pkey,
	}
	return &clienv, nil
}
