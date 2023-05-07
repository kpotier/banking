// Package bank defines an interface for connecting and fetching data from
// online banks. This package is designed with the hope that many different bank
// back-ends can conform to the interface.
package bank

import (
	"context"
	"errors"
	"time"

	"github.com/kpotier/banking/pkg/money"
)

// A Bank is the main interface or connecting and fetching data from online
// banks.
type Bank interface {
	// Login connects to the bank. Credentials may have been given when
	// instancing the bank. Questions like 2FA code may be asked through the
	// prompt channel (this method will block until it gets an answer).
	Login(username, pwd []byte, ctx context.Context, q chan<- string, a <-chan string) error
	// Logout disconnects from the bank. It may returns ErrNotLoggedIn if the
	// user is already disconnected.
	Logout() error

	// Accounts returns the user's accounts.
	Accounts() ([]*Account, error)
	// Transactions returns the list of transactions of a specified account that
	// occured after a defined time.
	Transactions(account *Account, debitUpTo time.Time) ([]*Transaction, error)
}

// An Account is a financial account maintained by a bank.
type Account struct {
	ID      string
	Name    string
	Type    AccountType
	Balance money.Money
}

// AccountType is the type of an account. For instance, it can be a stock
// account.
type AccountType int

const (
	AccountChecking AccountType = iota
	AccountSavings
	AccountStocks
)

// A Transaction is a record of money that has moved in or out an account.
type Transaction struct {
	ID        string
	Pending   bool
	DateDebit time.Time
	DateDone  time.Time
	RawName   string
	Name      string
	Card      string
	Type      TransactionType
	Category  string
	Value     money.Money
}

type TransactionType int

const (
	TransactionNone TransactionType = iota
	TransactionCard
	TransactionTrsf
	TransactionTrsfSEPA
	TransactionDDebitSEPA
	TransactionTrsfINST
	TransactionCredit
)

var (
	// ErrBadPwd must be used when the password does not match the specifications or is incorrect.
	ErrBadPwd error = errors.New("invalid password")
	// ErrBadLoginPwd must be used when the login OR password does not match the
	// specifications or are incorrect.
	ErrBadLoginPwd error = errors.New("invalid login or password")
	// ErrNotLoggedIn is returned when the user is not authenticated.
	ErrNotLoggedIn error = errors.New("not logged in")
)

// An UnrecognizedError is an unhandled error returned by the bank.
type UnrecognizedError string

func (e UnrecognizedError) Error() string { return string(e) }
