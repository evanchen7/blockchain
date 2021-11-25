package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

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
