package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dgraph-io/badger/v2"
)

func main() {
	opt := badger.DefaultOptions("").WithInMemory(true)
	opt.Logger = nil

	db, err := badger.Open(opt)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	filepath.Walk("./jpostcode-data/data/json", func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(info.Name(), ".json") {
			return nil
		}
		dataFile, err := os.Open("./jpostcode-data/data/json/" + info.Name())
		if err != nil {
			panic(err)
		}
		defer dataFile.Close()
		var addressMap map[string]interface{}
		if err := json.NewDecoder(dataFile).Decode(&addressMap); err != nil {
			panic(err)
		}
		firstPostCode := strings.TrimSuffix(info.Name(), ".json")
		for secondPostCode, val := range addressMap {
			postCode := firstPostCode + secondPostCode

			fmt.Println(postCode)
			addrJSON, err := json.Marshal(val)
			if err != nil {
				panic(err)
			}
			err = db.Update(func(txn *badger.Txn) error {
				err = txn.SetEntry(badger.NewEntry([]byte(postCode), addrJSON))
				return err
			})
			if err != nil {
				panic(err)
			}
		}
		return nil
	})

	f, err := os.OpenFile("./tmp/badger/dump.db", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = db.Backup(f, 0)
	if err != nil {
		panic(err)
	}

}
