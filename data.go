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
	ctx    context.Context

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

	return &DataLoader{cancel: cancel, ctx: ctx}, ctx
}

// LoaderOpts is a configuration for a DataLoader.
type LoaderOpts struct {
	// Limit is the number of records to load concurrently.
	Limit int
	// Offset is the number of records to skip.
	Offset int
	// OrderBy is the order of the records.
	OrderBy string
	// Order is the order of the records.
	Order string
	// Filter is the filter of the records.
	Filter string
	// Search is the search of the records.
	Search string
}

// NewLoaderOpts returns a new LoaderOpts.
func NewLoaderOpts() *LoaderOpts {
	return &LoaderOpts{}
}

// Configure is a configuration for a DataLoader Options.
func (o *LoaderOpts) Configure(opts ...LoaderOpt) *LoaderOpts {
	for _, opt := range opts {
		opt(o)
	}

	return o
}

// LoaderOpt is an option for a DataLoader.
type LoaderOpt func(*LoaderOpts)

// LoaderFunc is a function that loads data.
//
//	func LoadUsers(opts ...LoaderOpt) LoaderFunc {
//	   	opts := &LoaderOpts{}
//		opts.Configure(opts)
//
//	    return func(ctx context.Context) error {
//	        # Load users
//	        return nil
//		}
//	}
type LoaderFunc func(ctx context.Context) error

// Load loads data.
func (l *DataLoader) Load(f LoaderFunc) {
	if l.sem != nil {
		l.sem <- token{}
	}

	l.wg.Add(1)
	go func() {
		defer l.done()

		if err := f(l.ctx); err != nil {
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
