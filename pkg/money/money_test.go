package money

import (
	"reflect"
	"testing"
)

func TestMoney(t *testing.T) {
	tests := []struct {
		name    string
		m       *Money
		args    interface{}
		wantErr bool
		dontVal bool
	}{
		{
			"ok",
			&Money{GBP, 1234},
			"1234;GBP",
			false,
			false,
		},
		{
			"error format",
			nil,
			"1234",
			true,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var m Money
			if err := (&m).Scan(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("Money.Scan() error = %v, wantErr %v", err, tt.wantErr)
			} else if !tt.wantErr && !reflect.DeepEqual(&m, tt.m) {
				t.Errorf("Money.Value(), error = mismatch, want %v, got %v", tt.m, m)
			}

			if tt.dontVal {
				return
			}
			val, err := tt.m.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Money.Value() error = %v, wantErr %v", err, tt.wantErr)
			} else if !tt.wantErr && val != tt.args {
				t.Errorf("Money.Value(), error = mismatch, want %v, got %v", tt.args, val.(string))
			}
		})
	}
}
