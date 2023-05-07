package queue

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/kpotier/banking/pkg/bank"
)

const (
	okUsername = "myusername"
	okPwd      = "mypwd"
	okQuestion = "foo"
	okMsg      = "bar"

	timeout = 15 * time.Second
	upTo    = 15 * time.Second
)

var (
	errBadIdentifiers = errors.New("wrong identifiers")
	errAccount        = errors.New("account error")
	errTr             = errors.New("transactions error")
)

type FakeBank struct {
	account      bool
	transactions bool
}

func NewFakeBank() bank.Bank {
	return &FakeBank{}
}

func (f *FakeBank) Login(username, pwd []byte, ctx context.Context, q chan<- string, a <-chan string) error {
	if string(username) != okUsername || string(pwd) != okPwd {
		return errBadIdentifiers
	}

	// Send msg
	select {
	case q <- okQuestion:
	case <-ctx.Done():
		return ctx.Err()
	}

	// Wait for answer
	select {
	case msg := <-a:
		switch msg {
		case errAccount.Error():
			f.account = true
		case errTr.Error():
			f.transactions = true
		case okMsg:
		default:
			panic("bad message")
		}
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func (f *FakeBank) Logout() error {
	return nil
}

func (f *FakeBank) Accounts() ([]*bank.Account, error) {
	if f.account {
		return nil, errAccount
	}
	return []*bank.Account{{}}, nil
}

func (f *FakeBank) Transactions(account *bank.Account, after time.Time) ([]*bank.Transaction, error) {
	if float64(time.Until(after).Abs()) >= float64(upTo)*1.1 {
		return nil, errTr
	}
	if f.transactions {
		return nil, errTr
	}
	return nil, nil
}

func init() {
	bank.SetBank("fake", NewFakeBank)
}

func TestSimpleQueue_Set(t *testing.T) {
	for i := 0; i < 100; i++ {
		cap := 2000
		howManyFull := 500
		var (
			full atomic.Int32
			wg   sync.WaitGroup
		)
		q := (Queue)(&SimpleQueue{Queue: make([]*SimpleQueueJob, cap)})
		for i := 0; i < cap+howManyFull; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				_, ok := q.Set(context.Background(), "")
				if !ok {
					full.Add(1)
				}
			}(i)
		}
		wg.Wait()
		if full.Load() != int32(howManyFull) {
			t.Fatalf("queue should have been full for %v item(s), got %v", howManyFull, full.Load())
		}
	}
}

func TestSimpleQueue_get(t *testing.T) {
	cap := 2000
	q := &SimpleQueue{Queue: make([]*SimpleQueueJob, cap)}
	jobs := make([]Job, 0, cap)
	for i := 0; i < cap; i++ {
		j, ok := q.Set(context.Background(), "")
		if !ok {
			t.Fatalf("queue should not be full for item %v", i)
		}
		jobs = append(jobs, j)
	}
	for i := 0; i < cap+cap+1; i++ {
		j, ok := q.get()
		if i >= cap && ok {
			t.Fatalf("should return empty queue for item %v", i)
		} else if i < cap && j != jobs[i] {
			t.Errorf("incorrect order for item %v: want %v, got %v", i, jobs[i], j)
		}
	}
}

func TestSimpleQueueJob_Position(t *testing.T) {
	cap := 55
	q := &SimpleQueue{Queue: make([]*SimpleQueueJob, cap)}

	jobs := make([]Job, cap)
	for i := 0; i < cap; i++ {
		j, ok := q.Set(context.Background(), "nil")
		if !ok {
			t.Fatal("queue should not be full")
		}
		jobs[i] = j
	}
	for i, j := range jobs {
		pos := j.Position()
		if pos-1 != i {
			t.Errorf("wrong position for item %v: want %v, got %v", i, i, pos-1)
		}
	}
	read := cap / 2
	for i := 0; i < read; i++ {
		if _, ok := q.get(); !ok {
			t.Fatal("queue should not be empty")
		}
		j, ok := q.Set(context.Background(), "nil")
		if !ok {
			t.Fatal("queue should not be full")
		}
		jobs = append(jobs, j)
	}
	jobs = jobs[read:]
	for i, j := range jobs {
		pos := j.Position()
		if pos-1 != i {
			t.Errorf("wrong position for item %v: want %v, got %v", i, i, pos-1)
		}
	}

	// old job position
	if got := q.Position(&SimpleQueueJob{id: 3, setTime: time.Time{}}); got != 0 {
		t.Errorf("wrong position for old item: want %v, got %v", 0, got)
	}
}

func TestSimpleQueue(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		initQueue(t, context.Background(), true, true, true, okMsg, nil)
	})

	t.Run("bad bank", func(t *testing.T) {
		initQueue(t, context.Background(), false, false, false, "", ErrUnknownBank)
	})

	t.Run("bad login", func(t *testing.T) {
		initQueue(t, context.Background(), true, false, false, "", errBadIdentifiers)
	})

	t.Run("cancel ctx during login", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		start := time.Now()
		initQueue(t, ctx, true, true, true, okMsg, context.Canceled)
		if time.Until(start) >= timeout/2 {
			t.Error("job was waiting for messages and did not cancel")
		}
	})

	t.Run("account error", func(t *testing.T) {
		initQueue(t, context.Background(), true, true, true, errAccount.Error(), errAccount)
	})

	t.Run("tr error", func(t *testing.T) {
		initQueue(t, context.Background(), true, true, true, errTr.Error(), errTr)
	})
}

func initQueue(t *testing.T, ctx context.Context,
	okBank, okCredentials bool,
	doMsg bool, toSend string,
	wantErr error) {

	var bank string
	if okBank {
		bank = "fake"
	}

	var username, pwd string
	if okCredentials {
		username = okUsername
		pwd = okPwd
	}

	cap := 2
	q := NewSimpleQueue(timeout, upTo, int32(cap))

	j, ok := q.Set(ctx, bank)
	if !ok {
		t.Fatal("queue should not be full")
	}
	<-j.Started()
	answers := make(chan string)
	questions := j.Credentials([]byte(username), []byte(pwd), answers)
	if doMsg {
		go func() {
			defer close(answers)
			select {
			case question, ok := <-questions:
				if ok && question == okQuestion {
					select {
					case answers <- toSend:
					case <-j.Done():
					}
				}
			case <-j.Done():
			}
		}()
	}
	_, err := j.Results()

	if err != wantErr {
		t.Fatalf("j.Done() error = %v, wantErr %v", err, wantErr)
	}
}
