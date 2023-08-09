package stateutil

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

// DelegatorIndexMap constructs a map that associates delegator addresses with their corresponding indices.
// This function takes a slice of Delegator pointers as input and returns a map of common.Address to uint64.
// If the input slice is empty, an empty map is returned.
// The map is intended to help quickly locate the index of a delegator in a list.
func DelegatorIndexMap(dels []*types.Delegator) map[common.Address]uint64 {
	m := make(map[common.Address]uint64, len(dels))
	if len(dels) == 0 {
		return m
	}
	for idx, record := range dels {
		if record == nil {
			continue
		}
		m[record.Address] = uint64(idx)
	}
	return m
}
