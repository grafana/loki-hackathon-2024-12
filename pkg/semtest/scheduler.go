package semtest

import (
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type Test interface {
	Run() error
}

type Scheduler interface {
	// Start the scheduler in the current goroutine.
	Start()

	// Stop stops the scheduler.
	Stop()
}

type PeriodicScheduler struct {
	tests  []Test
	intv   time.Duration
	ticker *time.Ticker
	done   chan struct{}
	stop   chan struct{}
	logger log.Logger
}

// NewPeriodicScheduler returns a new scheduler that runs at a fixed periodic interval.
func NewPeriodicScheduler(intv time.Duration, logger log.Logger) *PeriodicScheduler {
	return &PeriodicScheduler{
		intv:   intv,
		ticker: time.NewTicker(intv),
		logger: logger,
	}
}

func (r *PeriodicScheduler) Start() {
	r.done = make(chan struct{})
	r.stop = make(chan struct{})
	go r.run()
}

func (r *PeriodicScheduler) run() {
	defer close(r.done)
	for {
		select {
		case <-r.ticker.C:
			r.runTests()
		case <-r.stop:
			return
		}
	}
}

func (r *PeriodicScheduler) runTests() {
	for _, test := range r.tests {
		if err := test.Run(); err != nil {
			level.Error(r.logger).Log("msg", "failed to run test: %w", err)
		}
	}
}

func (r *PeriodicScheduler) Stop() {
	close(r.stop)
	<-r.done
}

// StatsTester runs queries for a fixed time range and compares the stats against
// those from previous runs.
// type StatsTester struct {
// 	querier logproto.QuerierClient
// 	stats   StatsStore
// }

// func NewStatsTester(querier logproto.QuerierClient, stats StatsStore) *StatsTester {
// 	return &StatsTester{
// 		querier: querier,
// 		stats:   stats,
// 	}
// }

// func (t *StatsTester) CompareLogs(ctx context.Context, r *logproto.QueryRequest, stats stats.Result) error {
// 	client, err := t.querier.Query(ctx, r, nil)
// 	if err != nil {
// 		return err
// 	}
// 	_ = client
// 	fmt.Println(client)
// 	return nil
// }

// func (t *StatsTester) CompareSamples(ctx context.Context, r *logproto.SampleQueryRequest, stats stats.Result) error {
// 	return nil
// }

// type StatsStore interface {
// 	GetSample(ctx context.Context, sampleID string) (stats.Result, error)
// }

// type Result struct {
// }
