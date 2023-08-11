package epoch

import (
	"github.com/unicornultrafoundation/go-hashgraph/state"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

// ProcessFinalUpdates performs final updates on the U2U chain's state before finalizing.
// It includes processing effective staked balances for validators using ProcessEffectiveStakedBalances function.
// Returns the updated state after processing and any encountered error.
func ProcessFinalUpdates(s *state.State) (*state.State, error) {
	var err error

	s, err = ProcessEffectiveStakedBalances(s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// ProcessEffectiveStakedBalances updates the effective staked balances for validators in the state.
// It retrieves the staked balances from the state and compares them to the current effective staked balances of validators.
// If there's a difference, it creates a new validator object with the updated effective staked balance and replaces the old validator object.
// Returns the updated state after processing and any encountered error.
func ProcessEffectiveStakedBalances(s *state.State) (*state.State, error) {
	bals := s.StakedBalances()

	validatorFunc := func(idx int, val *types.Validator) (bool, *types.Validator, error) {
		bal := bals[idx]
		if bal != val.EffectiveStakedBalance {
			newVal := val.Clone()
			newVal.EffectiveStakedBalance = bal
			return true, newVal, nil
		}
		return false, val, nil
	}

	if err := s.ApplyToEveryValidator(validatorFunc); err != nil {
		return nil, err
	}

	return s, nil
}
