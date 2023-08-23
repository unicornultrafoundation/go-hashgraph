package helpers

import (
	"github.com/unicornultrafoundation/go-hashgraph/consensus/cache"
	"github.com/unicornultrafoundation/go-hashgraph/native/idx"
	ptypes "github.com/unicornultrafoundation/go-hashgraph/proto/u2u/types"
	"github.com/unicornultrafoundation/go-hashgraph/state"
)

var (
	validatorCache = cache.NewValidatorCache()
)

// TotalActiveStakedBalance calculates the total active staked balance in the network for the given epoch.
// It iterates through every validator using the ReadFromEveryValidator function and checks if each validator is active for the given epoch.
// If a validator is active, their effective staked balance is added to the total.
func TotalActiveStakedBalance(st *state.State, epoch idx.Epoch) (uint64, error) {
	cached, err := validatorCache.Get(epoch.Bytes())
	if err != nil {
		return 0, err
	}

	if cached != nil {
		return cached.TotalStakedBalance, nil
	}

	if err := UpdateValidatorCache(st, epoch); err != nil {
		return 0, err
	}
	return TotalActiveStakedBalance(st, epoch)
}

// IsActiveValidator checks if a validator is active for the given epoch.
// A validator is considered active if their activation epoch is less than or equal to the given epoch and the given epoch is less than their exit epoch.
func IsActiveValidator(val *ptypes.Validator, epoch idx.Epoch) bool {
	epochU64 := uint64(epoch)
	return val.ActivationEpoch <= epochU64 && epochU64 < val.ExitEpoch
}

// ActiveValidatorIndices returns the indices of active validators for the given epoch.
// It iterates through the validators in the state and checks if they are active for the specified epoch.
// Returns the list of indices and any encountered error.
func ActiveValidatorIndices(st *state.State, epoch idx.Epoch) ([]idx.ValidatorID, error) {
	cached, err := validatorCache.Get(epoch.Bytes())
	if err != nil {
		return nil, err
	}

	if cached != nil {
		return cached.Indices, nil
	}

	if err := UpdateValidatorCache(st, epoch); err != nil {
		return nil, err
	}
	return ActiveValidatorIndices(st, epoch)
}

func ActiveValidatorCount(st *state.State, epoch idx.Epoch) (int, error) {
	cached, err := validatorCache.Get(epoch.Bytes())
	if err != nil {
		return 0, err
	}

	if cached != nil {
		return int(cached.Count), nil
	}

	if err := UpdateValidatorCache(st, epoch); err != nil {
		return 0, err
	}
	return ActiveValidatorCount(st, epoch)
}

func UpdateValidatorCache(st *state.State, epoch idx.Epoch) error {
	indices := make([]idx.ValidatorID, 0)
	totalStakedBalances := uint64(0)
	count := 0
	if err := st.ReadFromEveryValidator(func(valId int, val *ptypes.Validator) error {
		if IsActiveValidator(val, epoch) {
			indices = append(indices, idx.ValidatorID(valId))
			totalStakedBalances += val.StakedBalance
			count++
		}
		return nil
	}); err != nil {
		return err
	}
	validatorCache.Set(&cache.Validators{
		Count:   uint64(count),
		Key:     epoch.Bytes(),
		Indices: indices,
	})
	return nil
}
