package address

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/syumai/go-jpostcode"
)

func Test_FromMap(t *testing.T) {
	tests := map[string]struct {
		input   interface{}
		want    *jpostcode.Address
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
			want: &jpostcode.Address{
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
			got, err := FromMap(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if d := cmp.Diff(got, tt.want); d != "" {
				t.Error(d)
			}
		})
	}
}
