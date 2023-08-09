package consensus

import (
	"github.com/unicornultrafoundation/go-hashgraph/state"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

// ProcessUndelegates processes a list of UnDelegateDto objects to perform undelegation actions.
// It iterates through each UnDelegateDto in the list and invokes ProcessUndelegate for each.
// If an error occurs during processing, the function returns an error.
// Otherwise, it returns the updated state after processing all undelegate actions.
func ProcessUndelegates(s *state.State, undelegateDtos []*types.UnDelegateDto) (*state.State, error) {
	var err error
	for _, undelegate := range undelegateDtos {
		s, err = ProcessUndelegate(s, undelegate)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

// ProcessUndelegate performs an undelegation action for a specified validator and delegator.
// It checks if the validator and delegator are valid and if there is an existing delegation.
// If a valid delegation is found, it calculates the amount to undelegate and updates the delegation's amount.
// Then, it appends a withdrawal request to the state with the undelegated amount.
// The function returns the updated state after processing the undelegation action.
func ProcessUndelegate(s *state.State, undelegateDto *types.UnDelegateDto) (*state.State, error) {
	var err error

	validatorIdx, ok := s.ValidatorIndexByAddress(undelegateDto.ValidatorAddress)
	if !ok {
		return s, nil
	}

	delegatorIdx, ok := s.DelegatorIndexByAddress(undelegateDto.DelegatorAddress)
	if !ok {
		return s, nil
	}

	delegationIdx, ok := s.DelegationIndexByDelegator(delegatorIdx, validatorIdx)
	if !ok {
		return s, nil
	}

	s, err = ProcessWithdrawReward(s, &types.WithdrawalRewardDto{
		ValidatorAddress: undelegateDto.ValidatorAddress,
		DelegatorAddress: undelegateDto.DelegatorAddress,
	})

	if err != nil {
		return nil, err
	}

	delegation := s.DelegationAtIndex(delegationIdx)
	var undelegateAmount uint64 = 0
	if delegation.Amount < undelegateDto.Amount {
		undelegateAmount = delegation.Amount
	} else {
		undelegateAmount = undelegateDto.Amount
	}

	delegation.Amount -= undelegateAmount

	s.UpdateDelegationAtIndex(delegationIdx, delegation)

	s.AppendWithdrawal(&types.Withdrawal{
		ValidatorIndex: validatorIdx,
		DelegatorIndex: delegatorIdx,
		Amount:         undelegateAmount,
		Epoch:          s.Epoch(),
	})

	return s, nil
}
