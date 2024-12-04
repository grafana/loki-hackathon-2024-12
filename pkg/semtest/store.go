package semtest

import (
	"github.com/grafana/loki/v3/pkg/logproto"
	"github.com/grafana/loki/v3/pkg/logqlmodel/stats"
)

type Store interface {
	GetStats(sampleID string) (StatsResult, error)
	GetRandomStats(limit int64) ([]StatsResult, error)
}

type StatsResult struct {
	Request logproto.QueryRequest
	Stats   stat.Result
}