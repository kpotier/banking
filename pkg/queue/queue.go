package queue

import (
	"context"
	"time"

	"github.com/kpotier/banking/pkg/bank"
)

type Queue interface {
	Set(ctx context.Context, bank string) (Job, bool)
}

type Job interface {
	// Position of the job in queue. If it is equal to zero then the job is
	// being or has been processed.
	Position() int
	Started() <-chan struct{}
	Credentials(login, pwd []byte, a <-chan string) (q <-chan string)
	Done() <-chan struct{}
	Results() ([]*Account, error)
}

type Account struct {
	Account          *bank.Account
	TransactionsUpTo time.Time
	Transactions     []*bank.Transaction
}
