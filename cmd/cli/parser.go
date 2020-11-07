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
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

func ParseArgs() {
	switch kingpin.Parse() {
	case getUtxo.FullCommand():
		if err := _getUtxo(); err != nil {
			log.Error(err)
		}
	case getBalance.FullCommand():
		if err := _getBalance(); err != nil {
			log.Error(err)
		}
	case getStatus.FullCommand():
		if err := _getStatus(); err != nil {
			log.Error(err)
		}
	case getTransaction.FullCommand():
		if err := _getTransaction(); err != nil {
			log.Error(err)
		}
	case approve.FullCommand():
		if err := _approve(); err != nil {
			log.Error(err)
		}
	case deposit.FullCommand():
		if err := _deposit(); err != nil {
			log.Error(err)
		}
	case send.FullCommand():
		if err := _send(); err != nil {
			log.Error(err)
		}
	case exit.FullCommand():
		if err := _exit(); err != nil {
			log.Error(err)
		}
	case process.FullCommand():
		if err := _process(); err != nil {
			log.Error(err)
		}
	case wallet.FullCommand():
		if err := _createKeypair(); err != nil {
			log.Error(err)
		}
	}
}
