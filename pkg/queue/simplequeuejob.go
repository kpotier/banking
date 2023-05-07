package queue

import (
	"context"
	"time"

	"github.com/kpotier/banking/pkg/bank"
)

type SimpleQueueJob struct {
	sq *SimpleQueue

	id      int32
	setTime time.Time

	ctx  context.Context
	bank string

	start chan struct{}
	done  context.Context
	cred  chan [2][]byte
	q     chan string
	a     <-chan string

	accounts []*Account
	err      error
}

func (s *SimpleQueueJob) Position() int {
	return int(s.sq.Position(s))
}

func (s *SimpleQueueJob) run() {
	ctx, cancel := context.WithTimeout(s.ctx, s.sq.JobTimeout)
	defer cancel()
	s.done = ctx
	s.cred = make(chan [2][]byte)
	close(s.start)

	newb, ok := bank.GetBank(s.bank)
	if !ok {
		s.err = ErrUnknownBank
		return
	}

	// We wait for the credentials.
	var cred [2][]byte
	s.q = make(chan string)
	select {
	case cred, ok = <-s.cred:
		if !ok {
			close(s.q)
			panic("unexpected closed channel")
		}
	case <-ctx.Done():
		close(s.q)
		s.err = ctx.Err()
		return
	}

	b := newb()
	err := b.Login(cred[0], cred[1], ctx, s.q, s.a)
	close(s.q) // no more questions
	if err != nil {
		s.err = err
		return
	}

	acc, err := b.Accounts()
	if err != nil {
		s.err = err
		return
	}

	upTo := time.Now().Add(-s.sq.TransactionsUpTo)
	for _, a := range acc {
		select {
		case <-ctx.Done():
			s.err = ctx.Err()
			return
		default:
		}
		tr, err := b.Transactions(a, upTo)
		if err != nil {
			s.err = err
			return
		}
		s.accounts = append(s.accounts, &Account{
			Account:          a,
			TransactionsUpTo: upTo,
			Transactions:     tr,
		})
	}
}

func (s *SimpleQueueJob) Started() <-chan struct{} {
	return s.start
}

// Credentials passes the username and password It should be called when the
// Started method closes. This method is meant to be called a single time so the
// credentials channel will be closed right after sending the credentials.
func (s *SimpleQueueJob) Credentials(login, pwd []byte, a <-chan string) (q <-chan string) {
	s.a = a
	select {
	case s.cred <- [2][]byte{login, pwd}:
	case <-s.Done():
	}
	close(s.cred)
	return s.q
}

func (s *SimpleQueueJob) Done() <-chan struct{} {
	return s.done.Done()
}

// Results blocks until the job is done.
func (s *SimpleQueueJob) Results() ([]*Account, error) {
	<-s.Done()
	return s.accounts, s.err
}
