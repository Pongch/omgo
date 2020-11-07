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
	"github.com/pongch/omgo/util"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"strings"
)

var (
	create = kingpin.Command("create", "Create a resource.")
	wallet = create.Command("keypair", "Create a test wallet consisting of Public and Private key")
)

func _createKeypair() error {
	address, pkey, err := util.GenerateAccount()
	if err != nil {
		return err
	}
	log.Infof("\n Address: %s \n Privatekey: 0x%s ", address, strings.ToUpper(pkey))
	return nil
}
