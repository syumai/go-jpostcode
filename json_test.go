package jpostcode

import (
	"encoding/json"
	iofs "io/fs"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/rakyll/statik/fs"
	_ "github.com/syumai/go-jpostcode/statik"
)

func Test_searchAddressesFromJSON_AllFiles(t *testing.T) {
	staticFS, err := fs.New()
	if err != nil {
		t.Fatal(err)
	}
	var postCodes []string
	fs.Walk(staticFS, "/", func(path string, info iofs.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), ".json") {
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

func Test_convertJSONToAddress(t *testing.T) {
	tests := map[string]struct {
		input   interface{}
		want    *Address
		wantErr bool
	}{
		"ok with valid input": {
			input: map[string]interface{}{
				"postcode":         "1638001",
				"prefecture":       "東京都",
				"prefecture_kana":  "トウキョウト",
				"prefecture_code":  13,
				"city":             "新宿区",
				"city_kana":        "シンジュクク",
				"town":             "西新宿",
				"town_kana":        "ニシシンジュク",
				"street":           "２丁目８−１",
				"office_name":      "東京都庁",
				"office_name_kana": "トウキヨウトチヨウ",
			},
			want: &Address{
				PostCode:       "1638001",
				Prefecture:     "東京都",
				PrefectureCode: 13,
				PrefectureKana: "トウキョウト",
				City:           "新宿区",
				CityKana:       "シンジュクク",
				Town:           "西新宿",
				TownKana:       "ニシシンジュク",
				Street:         "２丁目８−１",
				OfficeName:     "東京都庁",
				OfficeNameKana: "トウキヨウトチヨウ",
			},
			wantErr: false,
		},
		"ng with invalid input": {
			input: map[string]interface{}{
				"post_code":       999999,
				"prefecture":      999999,
				"prefecture_kana": 999999,
				"prefecture_code": "abcde",
				"city":            []int{1, 2, 3},
				"city_kana":       []float64{1, 2, 3},
			},
			wantErr: true,
		},
	}
	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			got, err := convertJSONToAddress(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertJSONToAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if d := cmp.Diff(got, tt.want); d != "" {
				t.Error(d)
			}
		})
	}
}

func Test_openDataFile(t *testing.T) {
	_, err := openDataFile("/001.json")
	if err != nil {
		t.Fatalf("openDataFile() error = %v, wantErr %v", err, nil)
	}
}
