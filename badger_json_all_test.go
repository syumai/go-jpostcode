package jpostcode

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAll_searchAddressesFromJSON_Files(t *testing.T) {
	var postCodes []string
	filepath.Walk("./jpostcode-data/data/json", func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".json") {
			return nil
		}
		dataFile, err := os.Open("./jpostcode-data/data/json/" + info.Name())
		if err != nil {
			t.Fatal(err)
		}
		defer dataFile.Close()
		var addressMap map[string]interface{}
		if err := json.NewDecoder(dataFile).Decode(&addressMap); err != nil {
			t.Fatal(err)
		}
		firstPostCode := strings.TrimSuffix(info.Name(), ".json")
		for secondPostCode := range addressMap {
			postCodes = append(postCodes, firstPostCode+secondPostCode)
			return nil
		}
		return nil
	})
	a, err := newBadgerAdapter()
	if err != nil {
		t.Fatal(err)
	}
	for _, postCode := range postCodes {
		t.Run(postCode[0:3], func(t *testing.T) {
			t.Parallel()
			addrs, err := a.SearchAddressesFromPostCode(postCode)
			if err != nil {
				t.Fatal(err)
			}
			if len(addrs) == 0 {
				t.Fatal("at least 1 address must be found")
			}
		})
	}
}
