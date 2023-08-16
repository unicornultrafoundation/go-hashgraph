package epoch

import (
	"github.com/unicornultrafoundation/go-hashgraph/config"
	"github.com/unicornultrafoundation/go-hashgraph/consensus/helpers"
	"github.com/unicornultrafoundation/go-hashgraph/math"
	"github.com/unicornultrafoundation/go-hashgraph/state"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

func ProcessEpochReward(s *state.State) (*state.State, error) {
	validatorFunc := func(idx int, val *types.Validator) (bool, *types.Validator, error) {
		return false, val, nil
	}

	if err := s.ApplyToEveryValidator(validatorFunc); err != nil {
		return nil, err
	}

	return s, nil
}

// BaseReward calculates the base reward for a validator based on their effective staked balance and the total active staked balance in the network.
// It uses the BaseRewardWithTotalStakedBalance function and the TotalActiveStakedBalance helper function to calculate the base reward.
func BaseReward(s *state.State, validatorIdx uint64) uint64 {
	totalStakedBalance := helpers.TotalActiveStakedBalance(s)
	return BaseRewardWithTotalStakedBalance(s, validatorIdx, totalStakedBalance)
}

// BaseRewardWithTotalStakedBalance calculates the base reward for a validator using their effective staked balance and the total active staked balance in the network.
// The function calculates the number of increments in the validator's effective staked balance and multiplies it by the base reward per increment.
// The base reward per increment is calculated using the BaseRewardPerIncrement function.
func BaseRewardWithTotalStakedBalance(s *state.State, validatorIdx uint64, totalStakedBalance uint64) uint64 {
	val := s.ValidatorAtIndex(validatorIdx)
	cfg := config.U2UConfig()
	increments := val.EffectiveStakedBalance / cfg.EffectiveStakedBalanceIncrement
	baseRewardPerInc := BaseRewardPerIncrement(totalStakedBalance)
	return increments * baseRewardPerInc
}

// BaseRewardPerIncrement calculates the base reward per increment based on the given staked balance.
// It uses the U2UConfig to get the effective staked balance increment and base reward factor,
// and calculates the base reward using the formula: (effective staked balance increment * base reward factor) / square root of staked balance.
func BaseRewardPerIncrement(stakedBalance uint64) uint64 {
	cfg := config.U2UConfig()
	return cfg.EffectiveStakedBalanceIncrement * cfg.BaseRewardFactor / math.CachedSquareRoot(stakedBalance)
}
