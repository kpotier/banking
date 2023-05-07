package queue

import (
	"context"
	"errors"
	"sync"
	"time"

	"go.uber.org/atomic"
)

var ErrUnknownBank = errors.New("bank doesn't exist")

type SimpleQueue struct {
	TransactionsUpTo time.Duration
	JobTimeout       time.Duration

	mux   sync.Mutex
	Queue []*SimpleQueueJob

	muxStart   sync.RWMutex
	startTime  time.Time
	end, start int32
	count      atomic.Int32
}

func NewSimpleQueue(jobTimeout time.Duration,
	transactionsDoneUpTo time.Duration, cap int32) Queue {
	s := &SimpleQueue{
		TransactionsUpTo: transactionsDoneUpTo,
		JobTimeout:       jobTimeout,
		Queue:            make([]*SimpleQueueJob, cap),
	}
	go s.Process()
	return s
}

// Set a job in queue. It returns false if the queue is full. This method is
// thread safe.
func (s *SimpleQueue) Set(ctx context.Context, bank string) (Job, bool) {
	s.mux.Lock()
	defer s.mux.Unlock()
	if int(s.count.Load()) == len(s.Queue) {
		return nil, false
	}
	j := &SimpleQueueJob{
		sq:      s,
		id:      s.end,
		setTime: time.Now(),
		ctx:     ctx,
		bank:    bank,
		start:   make(chan struct{}),
	}
	s.Queue[s.end] = j
	s.end = (s.end + 1) % int32(len(s.Queue))
	s.count.Add(1)
	return j, true
}

// Position returns the position of the job in queue. If it is equal to zero,
// the job is being or has been processed.
func (s *SimpleQueue) Position(j *SimpleQueueJob) int32 {
	s.muxStart.RLock()
	defer s.muxStart.RUnlock()
	if s.startTime.After(j.setTime) {
		return 0
	}
	if j.id >= s.start {
		return j.id - s.start + 1
	}
	return int32(len(s.Queue)) - s.start + j.id + 1
}

// Process get the first element in queue and run it. This method blocks
// forever and there is no way to unblock it.
func (s *SimpleQueue) Process() {
	for {
		j, ok := s.get()
		if !ok {
			continue
		}
		j.run()
	}
}

// Get the first element in queue. It returns false if the queue is empty. Be
// careful, unlike Set, this method is not thread safe.
func (s *SimpleQueue) get() (*SimpleQueueJob, bool) {
	if s.count.Load() == 0 {
		return nil, false
	}
	j := s.Queue[s.start]
	s.muxStart.Lock()
	s.startTime = j.setTime
	s.start = (s.start + 1) % int32(len(s.Queue))
	s.muxStart.Unlock()
	s.count.Add(-1)
	return j, true
}
