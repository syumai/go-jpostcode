package jpostcode

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/rakyll/statik/fs"
)

func TestAll_searchAddressesFromJSON_Files(t *testing.T) {
	staticFS, err := fs.New()
	if err != nil {
		t.Fatal(err)
	}
	var postCodes []string
	fs.Walk(staticFS, "/", func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(info.Name(), ".json") {
			return nil
		}
		dataFile, err := staticFS.Open("/" + info.Name())
		if err != nil {
			t.Fatal(err)
		}
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
	for _, postCode := range postCodes {
		t.Run(postCode[0:3], func(t *testing.T) {
			t.Parallel()
			addrs, err := searchAddressesFromJSON(postCode)
			if err != nil {
				t.Fatal(err)
			}
			if len(addrs) == 0 {
				t.Fatal("at least 1 address must be found")
			}
		})
	}
}
