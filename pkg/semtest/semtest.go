package semtest

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/go-kit/log"
	"google.golang.org/grpc"
	// "github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/grafana/dskit/grpcclient"
	"github.com/grafana/dskit/services"
	// "github.com/grafana/loki/v3/pkg/iter"
	"github.com/grafana/loki/v3/pkg/logproto"
	// "github.com/grafana/loki/v3/pkg/logql"
	// "github.com/grafana/loki/v3/pkg/logqlmodel/stats"
)

// Config for a semtester.
type Config struct {
	Address          string            `yaml:"address"`
	GRPCClientConfig grpcclient.Config `yaml:"grpc_client_config"`
	Interval         time.Duration     `yaml:"interval"`
}

func (cfg *Config) RegisterFlags(f *flag.FlagSet) {
	cfg.Address = "127.0.0.1:9096"
	cfg.GRPCClientConfig.RegisterFlagsWithPrefix("semtest.grpc-client-config", f)
	cfg.Interval = time.Minute
}

type SemanticTester struct {
	services.Service
	cfg       Config
	logger    log.Logger
	scheduler Scheduler
}

func New(cfg Config, logger log.Logger, _ prometheus.Registerer) (*SemanticTester, error) {
	opts, err := cfg.GRPCClientConfig.DialOption(nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC dial options: %w", err)
	}

	// TODO: Support polling DNS records over time like the query-frontend
	conn, err := grpc.NewClient(cfg.Address, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC client: %w", err)
	}

	querier := logproto.NewQuerierClient(conn)
	_ = querier
	t := SemanticTester{
		cfg:       cfg,
		logger:    logger,
		scheduler: NewPeriodicScheduler(cfg.Interval, logger),
	}
	t.Service = services.NewBasicService(t.starting, t.running, t.stopping)
	return &t, nil
}

func (t *SemanticTester) starting(ctx context.Context) (err error) {
	return nil
}

func (t *SemanticTester) running(ctx context.Context) error {
	t.scheduler.Start()
// 	ticker := time.NewTicker(time.Second)
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return nil
// 		case <-ticker.C:
// 			fmt.Println("tick")
// 			stream, err := t.querier.Query(context.Background(), &logproto.QueryRequest{
// 				Start:    time.Now().Add(-time.Hour),
// 				End:      time.Now(),
// 				Selector: `{foo="bar"}`,
// 			})
// 			if err != nil {
// 				fmt.Println(err)
// 			} else {
// 				resp, err := stream.Recv()
// 				if err != nil {
// 					fmt.Println(err)
// 					continue
// 				}
// 				fmt.Printf("%#v", resp.Stats)
// 				it := iter.NewQueryClientIterator(stream, logproto.FORWARD)
// 				for it.Next() {
// 					fmt.Println(it.At().Line)
// 				}
// 			}
// 		}
// 	}

	return nil
}

func (t *SemanticTester) stopping(_ error) error {
	t.scheduler.Stop()
	return nil
}
