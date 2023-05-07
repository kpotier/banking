package boursorama

import (
	"reflect"
	"testing"
	"time"

	"github.com/kpotier/banking/pkg/money"
)

func Test_parseDate(t *testing.T) {
	tests := []struct {
		arg     string
		want    string
		wantErr bool
	}{
		{"\t13 janvier \t2021", "13/01/2021", false},
		{" 01\nfévrier 2022", "01/02/2022", false},
		{" 1 mars  2023", "01/03/2023", false},
		{"\t22 avril 2024\t ", "22/04/2024", false},
		{"\n 2   mai 2025 ", "02/05/2025", false},
		{"09 \tjuin 2018", "09/06/2018", false},
		{" 19   juillet 2035   \t ", "19/07/2035", false},
		{"\n 22 août\n2100", "22/08/2100", false},
		{"21 septembre 2160", "21/09/2160", false},
		{"17   octobre 2068", "17/10/2068", false},
		{" 11 novembre 2027   ", "11/11/2027", false},
		{"\t\n 3 décembre \n 2028", "03/12/2028", false},
		{"\t\n 3 décembre", "", true},
		{"\t\n 3 jan \n 2028", "", true},
	}
	for _, tt := range tests {
		t.Run("parse date "+tt.arg, func(t *testing.T) {
			gotT, err := parseDate(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if tt.wantErr {
				return
			}
			res, err := time.Parse("02/01/2006", tt.want)
			if err != nil {
				panic(err)
			}
			if !reflect.DeepEqual(gotT, res) {
				t.Errorf("parseDate() = %v, want %v", gotT, res)
			}
		})
	}
}

func Test_parseValue(t *testing.T) {
	tests := []struct {
		arg     string
		want    int64
		wantErr bool
	}{
		{"3,00 €", 300, false},
		{"\n \t 180,02\t €\t", 18002, false},
		{"\n \t−\n 2 199 980,02\t €\t", -219998002, false},
		{"1 180 100,02", 0, true},
		{"00,02", 0, true},
		{"− 00,02", 0, true},
		{"− 00,02 $", 0, true},
	}
	for _, tt := range tests {
		t.Run("parse value "+tt.arg, func(t *testing.T) {
			gotT, err := parseValue(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if tt.wantErr {
				return
			}
			m := money.Money{Code: money.EUR, Amount: tt.want}
			if !reflect.DeepEqual(gotT, m) {
				t.Errorf("parseDate() = %v, want %v", gotT, tt.want)
			}
		})
	}
}
