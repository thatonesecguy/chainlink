package llo

import (
	"testing"

	"github.com/stretchr/testify/assert"

	llotypes "github.com/smartcontractkit/chainlink-common/pkg/types/llo"
)

func Test_ChannelDefinitionCache(t *testing.T) {
	t.Run("Definitions", func(t *testing.T) {
		// NOTE: this is covered more thoroughly in the integration tests
		dfns := llotypes.ChannelDefinitions(map[llotypes.ChannelID]llotypes.ChannelDefinition{
			1: {
				ReportFormat: llotypes.ReportFormat(43),
				Streams:      []llotypes.Stream{{StreamID: 1, Aggregator: llotypes.AggregatorMedian}, {StreamID: 2, Aggregator: llotypes.AggregatorMode}, {StreamID: 3, Aggregator: llotypes.AggregatorQuote}},
				Opts:         llotypes.ChannelOpts{1, 2, 3},
			},
		})

		cdc := &channelDefinitionCache{definitions: dfns}

		assert.Equal(t, dfns, cdc.Definitions())
	})
}
