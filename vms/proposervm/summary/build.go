// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package summary

import (
	"fmt"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/common"
)

// emptySummary has an ad-hoc construction in that it must have an empty summary ID
func BuildEmptyProposerSummary(coreSummary common.Summary) (ProposerSummaryIntf, error) {
	return &ProposerSummary{
		StatelessSummary: StatelessSummary{
			ProBlkID:             ids.Empty,
			InnerSummary:         coreSummary.Bytes(),
			ProposerSummaryBytes: nil,
			ProposerSummaryID:    ids.Empty,
		},
		SummaryHeight: coreSummary.Height(),
	}, nil
}

func BuildProposerSummary(proBlkID ids.ID, coreSummary common.Summary) (ProposerSummaryIntf, error) {
	statelessSummary := StatelessSummary{
		ProBlkID:     proBlkID,
		InnerSummary: coreSummary.Bytes(),
	}

	proSummaryBytes, err := cdc.Marshal(codecVersion, &statelessSummary)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal proposer summary due to: %w", err)
	}
	if err := statelessSummary.initialize(proSummaryBytes); err != nil {
		return nil, err
	}

	return &ProposerSummary{
		StatelessSummary: statelessSummary,
		SummaryHeight:    coreSummary.Height(),
	}, nil
}