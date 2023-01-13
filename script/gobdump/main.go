package main

import (
	"compress/gzip"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/syumai/go-jpostcode"
)

func main() {
	m := make(map[string][]*jpostcode.Address)

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

			addresses, err := decodeAddresses(val)
			if err != nil {
				panic(err)
			}
			m[postCode] = addresses
		}
		return nil
	})

	const flag = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	f, err := os.OpenFile("./data/map.gob.gz", flag, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	gw, err := gzip.NewWriterLevel(f, gzip.BestCompression)
	if err != nil {
		panic(err)
	}

	genc := gob.NewEncoder(gw)
	if err := genc.Encode(m); err != nil {
		panic(err)
	}

	if err := gw.Close(); err != nil {
		panic(err)
	}
}

func decodeAddresses(addressData interface{}) ([]*jpostcode.Address, error) {
	var addresses []*jpostcode.Address
	switch reflect.TypeOf(addressData).Kind() {
	case reflect.Slice:
		rawAddrs, ok := addressData.([]interface{})
		if !ok {
			return nil, jpostcode.ErrInternal
		}
		for _, rawAddr := range rawAddrs {
			addr, err := jpostcode.AddressFromMap(rawAddr)
			if err != nil {
				return nil, err
			}
			addresses = append(addresses, addr)
		}
	default:
		addr, err := jpostcode.AddressFromMap(addressData)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, addr)
	}
	return addresses, nil
}
