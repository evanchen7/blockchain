package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var genesisJson = `
{
  "genesis_time": "2019-03-18T00:00:00.000000000Z",
  "chain_id": "the-blockchain-bar-ledger",
  "balances": {
    "andrej": 1000000
  }
}`

type genesis struct {
	Balances map[Account]uint `json:"balances"`
}

func loadGenesis(path string) (genesis, error) {
	content, err := ioutil.ReadFile(path)
	handleError(err)

	var loadedGenesis genesis
	err = json.Unmarshal(content, &loadedGenesis)
	handleError(err)

	return loadedGenesis, nil
}

func handleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func writeGenesisToDisk(path string) error {
	return ioutil.WriteFile(path, []byte(genesisJson), 0644)
}
