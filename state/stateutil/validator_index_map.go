package stateutil

import (
	"github.com/ethereum/go-ethereum/common"
	ptypes "github.com/unicornultrafoundation/go-hashgraph/proto/u2u/types"
)

// ValidatorIndexMap constructs a map that associates validator addresses with their corresponding indices.
// This function takes a slice of Validator pointers as input and returns a map of common.Address to uint64.
// If the input slice is empty, an empty map is returned.
// The map is intended to help quickly locate the index of a validator in a list.
func ValidatorIndexMap(vals []*ptypes.Validator) map[common.Address]uint64 {
	m := make(map[common.Address]uint64, len(vals))
	if len(vals) == 0 {
		return m
	}
	for idx, record := range vals {
		if record == nil {
			continue
		}
		m[common.BytesToAddress(record.Address)] = uint64(idx)
	}
	return m
}
