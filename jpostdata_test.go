package jpostcode

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Find(t *testing.T) {
	tests := map[string]struct {
		postCode string
		want     *Address
		wantErr  error
	}{
		"OK": {
			postCode: "1638001",
			want: &Address{
				PostCode:       "1638001",
				Prefecture:     "東京都",
				PrefectureKana: "トウキョウト",
				PrefectureCode: 13,
				City:           "新宿区",
				CityKana:       "シンジュクク",
				Town:           "西新宿",
				TownKana:       "ニシシンジュク",
				Street:         "２丁目８−１",
				OfficeName:     "東京都庁",
				OfficeNameKana: "トウキヨウトチヨウ",
			},
		},
		"NG with invalid too long postCode": {
			postCode: "12345678", // too long
			want:     nil,
			wantErr:  ErrInvalidArgument,
		},
		"NG with invalid too short postCode": {
			postCode: "123456", // too short
			want:     nil,
			wantErr:  ErrInvalidArgument,
		},
		"NG with not found by firstPostCode": {
			postCode: "0006125", // not found by first 3 digits
			want:     nil,
			wantErr:  ErrNotFound,
		},
		"NG with not found by lastPostCode": {
			postCode: "1060000", // not found by last 4 digits
			want:     nil,
			wantErr:  ErrNotFound,
		},
	}

	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			got, err := Find(tt.postCode)
			if tt.wantErr != nil {
				if tt.wantErr != err {
					t.Fatalf("want err: %v, got: %v", tt.wantErr, err)
				}
				return
			}
			if d := cmp.Diff(tt.want, got); d != "" {
				t.Fatal(d)
			}
		})
	}
}
