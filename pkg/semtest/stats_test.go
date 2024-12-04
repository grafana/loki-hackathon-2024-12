package semtest

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/grafana/loki/v3/pkg/logqlmodel/stats"
)

var (
	result1 = stats.Result{
		Summary: stats.Summary{
			BytesProcessedPerSecond: 1024,
		},
	}
	result2 = stats.Result{
		Summary: stats.Summary{
			BytesProcessedPerSecond: 512,
		},
	}
)

func TestDistanceStatsDiffer(t *testing.T) {
	s := DistanceStatsDiffer{}
	diff, err := s.Diff(result1, result2)
	require.NoError(t, err)
	t.Log(diff)
}
