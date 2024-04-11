package htmx

import (
	"context"
	"fmt"
	"sync"
)

type token struct{}

// DataLoader is creating go routines to load data.
type DataLoader struct {
	cancel func(error)

	wg  sync.WaitGroup
	sem chan token

	errOnce sync.Once
	err     error
}

func (l *DataLoader) done() {
	if l.sem != nil {
		<-l.sem
	}
	l.wg.Done()
}

// WithContext returns a new DataLoader with the given context.
func WithContext(ctx context.Context) (*DataLoader, context.Context) {
	ctx, cancel := context.WithCancelCause(ctx)

	return &DataLoader{cancel: cancel}, ctx
}

// Load loads data.
func (l *DataLoader) Load(f func() error) {
	if l.sem == nil {
		l.sem <- token{}
	}

	l.wg.Add(1)
	go func() {
		defer l.done()

		if err := f(); err != nil {
			l.errOnce.Do(func() {
				l.err = err
				if l.cancel != nil {
					l.cancel(err)
				}
			})
		}
	}()
}

// Wait waits for all data loading to complete.
func (l *DataLoader) Wait() error {
	l.wg.Wait()
	if l.cancel != nil {
		l.cancel(l.err)
	}

	return l.err
}

// SetLimit limits the number of concurrent data loading.
func (l *DataLoader) SetLimit(n int) {
	if n < 0 {
		l.sem = nil
		return
	}

	if len(l.sem) != 0 {
		panic(fmt.Errorf("dataloader: modify limit while %v goroutines in the loader are still active", len(l.sem)))
	}

	l.sem = make(chan token, n)
}
