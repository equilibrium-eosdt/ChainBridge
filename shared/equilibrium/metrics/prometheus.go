package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

type ChainMetrics struct {
	BlocksProcessed prometheus.Counter
	LatestProcessedBlock prometheus.Gauge
	LatestKnownBlock prometheus.Gauge
	VotesSubmitted  prometheus.Counter
	LatestBlocksRequested prometheus.Counter
}

func NewChainMetrics(chain string) *ChainMetrics {
	metrics := &ChainMetrics{
		BlocksProcessed: prometheus.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_blocks_processed", chain),
			Help: "Number of blocks processed by the chain's listener",
		}),
		LatestProcessedBlock: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: fmt.Sprintf("%s_latest_processed_block", chain),
			Help: "Latest block processed by the listener",
		}),
		LatestKnownBlock: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: fmt.Sprintf("%s_latest_known_block", chain),
			Help: "Latest block the listener has seen",
		}),
		VotesSubmitted: prometheus.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_votes_submitted", chain),
			Help: "Number of votes submitted to chain",
		}),
		LatestBlocksRequested: prometheus.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_latest_blocks_requested", chain),
			Help: "Number of latest block requests performed by the chain's listener",
		}),
	}

	prometheus.MustRegister(metrics.BlocksProcessed)
	prometheus.MustRegister(metrics.LatestProcessedBlock)
	prometheus.MustRegister(metrics.LatestKnownBlock)
	prometheus.MustRegister(metrics.VotesSubmitted)
	prometheus.MustRegister(metrics.LatestBlocksRequested)

	return metrics
}