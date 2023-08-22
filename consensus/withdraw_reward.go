package consensus

import (
	ptypes "github.com/unicornultrafoundation/go-hashgraph/proto/u2u/types"
	"github.com/unicornultrafoundation/go-hashgraph/state"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

// ProcessWithdrawRewards processes withdrawal rewards for multiple delegators and validators.
// It takes the current state 's' and a slice of WithdrawalRewardDto pointers 'withdrawalRewardDtos' as input.
// This function iterates through the list of withdrawal reward DTOs and applies the withdrawal reward process
// for each of them using the ProcessWithdrawReward function. If any error occurs during processing, it returns
// nil and the error. Otherwise, it returns the updated state and nil error.
func ProcessWithdrawRewards(s *state.State, withdrawalRewardDtos []*types.WithdrawalRewardDto) (*state.State, error) {
	var err error
	for _, withdrawalRewardDto := range withdrawalRewardDtos {
		s, err = ProcessWithdrawReward(s, withdrawalRewardDto)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

// ProcessWithdrawReward processes withdrawal rewards for a single delegator and validator.
// It takes the current state 's' and a WithdrawalRewardDto pointer 'withdrawalRewardDto' as input.
// This function first retrieves the validator index and delegator index using the provided addresses.
// If the validator or delegator does not exist in the state, it returns the unchanged state.
// Next, it obtains the delegation index using the retrieved indices. If the delegation does not exist,
// it also returns the unchanged state. The function then calculates the reward amount by considering the
// difference in accumulated rewards and the delegation's accumulated reward per token, multiplied by the delegation amount.
// If the calculated reward is zero, the function returns the unchanged state. Otherwise, it updates the delegation's
// last accumulated reward per token, updates the delegation in the state, and appends a new WithdrawalReward to
// the state's withdrawalRewards list. Finally, the function returns the updated state and nil error.
func ProcessWithdrawReward(s *state.State, withdrawalRewardDto *types.WithdrawalRewardDto) (*state.State, error) {
	validatorIdx, ok := s.ValidatorIndexByAddress(withdrawalRewardDto.ValidatorAddress)
	if !ok {
		return s, nil
	}
	delegatorIdx, ok := s.DelegatorIndexByAddress(withdrawalRewardDto.DelegatorAddress)
	if !ok {
		return s, nil
	}

	delegationIdx, ok := s.DelegationIndexByDelegator(delegatorIdx, validatorIdx)
	if !ok {
		return s, nil
	}

	delegation := s.DelegationAtIndex(delegationIdx)
	val := s.ValidatorAtIndex(validatorIdx)
	reward := (val.AccumulatedRewardPerToken - delegation.LastAccumulatedRewardPerToken) * delegation.Amount
	if reward == 0 {
		return s, nil
	}
	delegation.LastAccumulatedRewardPerToken = val.AccumulatedRewardPerToken
	s.UpdateDelegationAtIndex(delegationIdx, delegation)

	s.AppendWithdrawalReward(&ptypes.WithdrawalReward{
		DelegatorIndex: delegationIdx,
		Amount:         reward,
	})
	return s, nil
}
