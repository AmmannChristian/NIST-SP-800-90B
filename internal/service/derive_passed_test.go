package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	pb "github.com/AmmannChristian/nist-800-90b/pkg/pb"
)

func TestDerivePassedFromEstimators(t *testing.T) {
	passed := func(name string) *pb.Sp80090BEstimatorResult {
		return &pb.Sp80090BEstimatorResult{Name: name, Passed: true}
	}
	failed := func(name string) *pb.Sp80090BEstimatorResult {
		return &pb.Sp80090BEstimatorResult{Name: name, Passed: false}
	}

	tests := []struct {
		name      string
		iid       []*pb.Sp80090BEstimatorResult
		nonIID    []*pb.Sp80090BEstimatorResult
		expectAll bool
	}{
		{
			name:      "all passed",
			iid:       []*pb.Sp80090BEstimatorResult{passed("MCV"), passed("Chi2")},
			nonIID:    []*pb.Sp80090BEstimatorResult{passed("MCV"), passed("Collision")},
			expectAll: true,
		},
		{
			name:      "iid failure propagates",
			iid:       []*pb.Sp80090BEstimatorResult{passed("MCV"), failed("Chi2")},
			nonIID:    []*pb.Sp80090BEstimatorResult{passed("MCV")},
			expectAll: false,
		},
		{
			name:      "non-iid failure propagates",
			iid:       []*pb.Sp80090BEstimatorResult{passed("MCV")},
			nonIID:    []*pb.Sp80090BEstimatorResult{passed("MCV"), failed("Collision")},
			expectAll: false,
		},
		{
			name:      "empty results returns true",
			iid:       nil,
			nonIID:    nil,
			expectAll: true,
		},
		{
			name:      "only iid with failure",
			iid:       []*pb.Sp80090BEstimatorResult{failed("MCV")},
			nonIID:    nil,
			expectAll: false,
		},
		{
			name:      "only non-iid all passed",
			iid:       nil,
			nonIID:    []*pb.Sp80090BEstimatorResult{passed("MCV"), passed("Collision")},
			expectAll: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := derivePassedFromEstimators(tt.iid, tt.nonIID)
			assert.Equal(t, tt.expectAll, result)
		})
	}
}
