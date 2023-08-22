package consensus

import (
	ptypes "github.com/unicornultrafoundation/go-hashgraph/proto/u2u/types"
	"github.com/unicornultrafoundation/go-hashgraph/state"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

// ProcessDelegates processes a list of delegate data and updates the state accordingly.
// It iterates through each delegate in the list and invokes ProcessDelegate for each.
// If an error occurs during processing, the function returns an error.
// Otherwise, it returns the updated state after processing all delegates.
func ProcessDelegates(s *state.State, delegates []*types.DelegateDto) (*state.State, error) {
	var err error
	for _, delegate := range delegates {
		s, err = ProcessDelegate(s, delegate)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

// ProcessDelegate processes a delegate and updates the state with delegation information.
// It checks if the validator and delegator exist in the state and processes accordingly.
// If the validator doesn't exist, the function returns the original state.
// If the delegator doesn't exist, it appends a new delegator to the state.
// If a delegation exists, it updates the existing delegation with additional amount.
// Otherwise, it appends a new delegation to the state.
// The function returns the updated state after processing the delegate.
func ProcessDelegate(s *state.State, delegate *types.DelegateDto) (*state.State, error) {
	var err error
	validatorIdx, ok := s.ValidatorIndexByAddress(delegate.ValidatorAddress)
	if !ok {
		return s, nil
	}

	delegatorIdx, ok := s.DelegatorIndexByAddress(delegate.DelegatorAddress)
	if !ok {
		s.AppendDelegator(&ptypes.Delegator{Address: delegate.DelegatorAddress[:]})
	} else {
		delegatorIdx, _ = s.DelegatorIndexByAddress(delegate.DelegatorAddress)
	}

	delegationIdx, ok := s.DelegationIndexByDelegator(delegatorIdx, validatorIdx)

	lastAccumulatedRewardPerToken := s.ValidatorAtIndex(validatorIdx).AccumulatedRewardPerToken

	if !ok {
		s.AppendDelegation(&ptypes.Delegation{
			ValidatorIndex:                validatorIdx,
			DelegatorIndex:                delegatorIdx,
			Amount:                        delegate.Amount,
			LastAccumulatedRewardPerToken: lastAccumulatedRewardPerToken,
		})
	} else {
		s, err = ProcessWithdrawReward(s, &types.WithdrawalRewardDto{
			ValidatorAddress: delegate.ValidatorAddress,
			DelegatorAddress: delegate.DelegatorAddress,
		})

		if err != nil {
			return nil, err
		}

		delegation := s.DelegationAtIndex(delegationIdx)
		delegation.Amount = delegation.Amount + delegate.Amount
		delegation.LastAccumulatedRewardPerToken = lastAccumulatedRewardPerToken
		s.UpdateDelegationAtIndex(delegationIdx, delegation)
	}

	s.UpdateStakedBalanceAtIndex(validatorIdx, s.StakedBalanceAtIndex(validatorIdx)+delegate.Amount)
	return s, nil
}
