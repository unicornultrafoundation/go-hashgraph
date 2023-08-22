package precompute

import (
	"github.com/unicornultrafoundation/go-hashgraph/config"
	ptypes "github.com/unicornultrafoundation/go-hashgraph/proto/u2u/types"
	"github.com/unicornultrafoundation/go-hashgraph/state"
)

// ProcessRewardsAndPenaltiesPrecompute calculates and applies rewards, penalties, and updates to validator attributes.

// This function takes the U2U network state, a list of validators, and balance information as input.
// It computes rewards and penalties based on validator attributes, rewards configuration, and balances.
// Then, it applies the calculated rewards and penalties to the validators' attributes.
// Returns the updated state after processing rewards, penalties, and updates.
func ProcessRewardsAndPenaltiesPrecompute(s *state.State, validators []*Validator, bal *Balance) (*state.State, error) {
	delReward, vRewards, penalties := rewardsAndPenalties(s, validators, bal)
	bals := s.StakedBalances()
	validatorFunc := func(idx int, val *ptypes.Validator) (bool, *ptypes.Validator, error) {
		if delReward[idx] > 0 || vRewards[idx] > 0 || penalties[idx] > 0 {
			newVal := val
			newVal.CommissionReward += vRewards[idx]
			newVal.AccumulatedRewardPerToken += delReward[idx]
			bals[idx] = bals[idx] - penalties[idx]
			newVal.PrevTxFees = newVal.TxFees
			return true, newVal, nil
		}
		return false, val, nil
	}

	if err := s.ApplyToEveryValidator(validatorFunc); err != nil {
		return nil, err
	}
	s.SetStakedBalances(bals)
	return s, nil
}

func rewardsAndPenalties(s *state.State, validators []*Validator, bal *Balance) ([]uint64, []uint64, []uint64) {
	cfg := config.U2UConfig()
	dRewards := make([]uint64, len(validators))
	vRewards := make([]uint64, len(validators))
	penalties := make([]uint64, len(validators))
	for idx, val := range validators {
		if val.IsActive {
			reward := bal.Reward * val.BaseRewardWeight / bal.BaseRewardWeight
			reward += bal.EpochFee * val.TxRewardWeight / bal.TxRewardWeght
			commissionReward := reward / cfg.CommissionQuotient
			delegatorsReward := reward - commissionReward
			vRewards[idx] = commissionReward
			dRewards[idx] = (delegatorsReward * cfg.BalanceIncrement) / bal.Active
		}

	}
	return dRewards, vRewards, penalties
}
