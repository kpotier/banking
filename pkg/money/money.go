package money

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

// CurrencyCode is a three-letter currency code defined in ISO 4217.
type CurrencyCode string

// List of ISO 4217 codes.
const (
	EUR CurrencyCode = "EUR"
	GBP CurrencyCode = "GBP"
)

// Money represents an amount of money with its currency type.
type Money struct {
	Code CurrencyCode
	// Amount of money as integer number of minor units. For instance 3.56€ is
	// 356 because EUR has two decimals.
	Amount int64
}

// Scan implements the sql.Scanner interface.
func (m *Money) Scan(value interface{}) error {
	val, ok := value.(string)
	if !ok {
		return fmt.Errorf("wrong type %v", value)
	}
	idx := strings.IndexRune(val, ';')
	if idx == -1 || idx+1 >= len(val) {
		return fmt.Errorf("invalid format")
	}
	amount, err := strconv.ParseInt(val[:idx], 10, 64)
	if err != nil {
		return err
	}
	m.Amount = amount
	m.Code = CurrencyCode(val[idx+1:])
	return nil
}

// Value implements the sql.Valuer interface.
func (m Money) Value() (driver.Value, error) {
	var str string
	str += strconv.FormatInt(m.Amount, 10)
	str += ";"
	str += string(m.Code)
	return str, nil
}
