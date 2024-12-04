package semtest

import (
	"github.com/grafana/loki/v3/pkg/logqlmodel/stats"
)

type StatsDiffer interface {
	Diff(stats.Result, stats.Result) (stats.Result, error)
}

type DistanceStatsDiffer struct{}

func (s *DistanceStatsDiffer) Diff(a, b stats.Result) (stats.Result, error) {
	return stats.Result{
		Summary:  s.diffSummaries(a.Summary, b.Summary),
		Querier:  s.diffQuerier(a.Querier, b.Querier),
		Ingester: s.diffIngester(a.Ingester, b.Ingester),
		Caches:   s.diffCaches(a.Caches, b.Caches),
		Index:    s.diffIndex(a.Index, b.Index),
	}, nil
}

func (s *DistanceStatsDiffer) diffSummaries(a, b stats.Summary) stats.Summary {
	return stats.Summary{
		BytesProcessedPerSecond:               a.BytesProcessedPerSecond - b.BytesProcessedPerSecond,
		LinesProcessedPerSecond:               a.LinesProcessedPerSecond - b.LinesProcessedPerSecond,
		TotalBytesProcessed:                   a.TotalBytesProcessed - b.TotalBytesProcessed,
		TotalLinesProcessed:                   a.TotalLinesProcessed - b.TotalLinesProcessed,
		ExecTime:                              a.ExecTime - b.ExecTime,
		QueueTime:                             a.QueueTime - b.QueueTime,
		Subqueries:                            a.Subqueries - b.Subqueries,
		TotalEntriesReturned:                  a.TotalEntriesReturned - b.TotalEntriesReturned,
		Splits:                                a.Splits - b.Splits,
		Shards:                                a.Shards - b.Shards,
		TotalPostFilterLines:                  a.TotalPostFilterLines - b.TotalPostFilterLines,
		TotalStructuredMetadataBytesProcessed: a.TotalStructuredMetadataBytesProcessed - b.TotalStructuredMetadataBytesProcessed,
	}
}

func (s *DistanceStatsDiffer) diffQuerier(a, b stats.Querier) stats.Querier {
	return stats.Querier{
		Store: s.diffStore(a.Store, b.Store),
	}
}

func (s *DistanceStatsDiffer) diffIngester(a, b stats.Ingester) stats.Ingester {
	return stats.Ingester{
		TotalReached:       a.TotalReached - b.TotalReached,
		TotalChunksMatched: a.TotalChunksMatched - b.TotalChunksMatched,
		TotalBatches:       a.TotalBatches - b.TotalBatches,
		TotalLinesSent:     a.TotalLinesSent - b.TotalLinesSent,
		Store:              s.diffStore(a.Store, b.Store),
	}
}

func (s *DistanceStatsDiffer) diffStore(a, b stats.Store) stats.Store {
	return stats.Store{
		TotalChunksRef:        a.TotalChunksRef - b.TotalChunksRef,
		TotalChunksDownloaded: a.TotalChunksDownloaded - b.TotalChunksDownloaded,
		ChunksDownloadTime:    a.ChunksDownloadTime - b.ChunksDownloadTime,
		// QueryReferencedStructured:    a.QueryReferencedStructured - b.QueryReferencedStructured, // bool
		ChunkRefsFetchTime:           a.ChunkRefsFetchTime - b.ChunkRefsFetchTime,
		CongestionControlLatency:     a.CongestionControlLatency - b.CongestionControlLatency,
		PipelineWrapperFilteredLines: a.PipelineWrapperFilteredLines - b.PipelineWrapperFilteredLines,
	}
}

func (s *DistanceStatsDiffer) diffCaches(a, b stats.Caches) stats.Caches {
	return stats.Caches{
		Chunk:               s.diffCache(a.Chunk, b.Chunk),
		Index:               s.diffCache(a.Index, b.Index),
		Result:              s.diffCache(a.Result, b.Result),
		StatsResult:         s.diffCache(a.StatsResult, b.StatsResult),
		VolumeResult:        s.diffCache(a.VolumeResult, b.VolumeResult),
		SeriesResult:        s.diffCache(a.SeriesResult, b.SeriesResult),
		LabelResult:         s.diffCache(a.LabelResult, b.LabelResult),
		InstantMetricResult: s.diffCache(a.InstantMetricResult, b.InstantMetricResult),
	}
}

func (s *DistanceStatsDiffer) diffCache(a, b stats.Cache) stats.Cache {
	return stats.Cache{
		EntriesFound:      a.EntriesFound - b.EntriesFound,
		EntriesRequested:  a.EntriesRequested - b.EntriesRequested,
		EntriesStored:     a.EntriesStored - b.EntriesStored,
		BytesReceived:     a.BytesReceived - b.BytesReceived,
		BytesSent:         a.BytesSent - b.BytesSent,
		Requests:          a.Requests - b.Requests,
		DownloadTime:      a.DownloadTime - b.DownloadTime,
		QueryLengthServed: a.QueryLengthServed - b.QueryLengthServed,
	}
}

func (s *DistanceStatsDiffer) diffIndex(a, b stats.Index) stats.Index {
	return stats.Index{
		TotalChunks:      a.TotalChunks - b.TotalChunks,
		PostFilterChunks: a.PostFilterChunks - b.PostFilterChunks,
		ShardsDuration:   a.ShardsDuration - b.ShardsDuration,
		// UsedBloomFilters: a.UsedBloomFilters - b.UsedBloomFilters, // bool
	}
}
