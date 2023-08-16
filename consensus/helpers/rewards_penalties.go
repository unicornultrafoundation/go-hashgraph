package helpers

import (
	"github.com/unicornultrafoundation/go-hashgraph/native/idx"
	"github.com/unicornultrafoundation/go-hashgraph/state"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

// TotalActiveStakedBalance calculates the total active staked balance in the network for the given epoch.
// It iterates through every validator using the ReadFromEveryValidator function and checks if each validator is active for the given epoch.
// If a validator is active, their effective staked balance is added to the total.
func TotalActiveStakedBalance(s *state.State) uint64 {
	total := uint64(0)
	epoch := s.Epoch()
	_ = s.ReadFromEveryValidator(func(idx int, val *types.Validator) error {
		if IsActiveValidator(val, epoch) {
			total += val.EffectiveStakedBalance
		}
		return nil
	})
	return total
}

// IsActiveValidator checks if a validator is active for the given epoch.
// A validator is considered active if their activation epoch is less than or equal to the given epoch and the given epoch is less than their exit epoch.
func IsActiveValidator(val *types.Validator, epoch idx.Epoch) bool {
	return val.ActivationEpoch <= epoch && epoch < val.ExitEpoch
}
