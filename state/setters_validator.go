package state

import (
	"github.com/unicornultrafoundation/go-hashgraph/state/stateutil"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

// SetValidators updates the list of active validators in the U2U chain state.
// This function acquires a lock to ensure thread safety while modifying validator-related data in the State.
// It takes a slice of Validator pointers as input and updates the internal validator list and index map accordingly.
// Ensure that you use this function within a context that handles concurrent access appropriately.
func (s *State) SetValidators(vals []*types.Validator) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.validators = vals
	s.valMap = stateutil.ValidatorIndexMap(vals)
}
