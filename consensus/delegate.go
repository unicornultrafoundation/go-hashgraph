package consensus

import (
	"github.com/unicornultrafoundation/go-hashgraph/state"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

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

func ProcessDelegate(s *state.State, delegate *types.DelegateDto) (*state.State, error) {
	validatorIdx, ok := s.ValidatorIndexByAddress(delegate.ValidatorAddress)
	if !ok {
		return s, nil
	}

	delegatorIdx, ok := s.DelegatorIndexByAddress(delegate.DelegatorAddress)
	if !ok {
		s.AppendDelegator(&types.Delegator{Address: delegate.DelegatorAddress})
	} else {
		delegatorIdx, _ = s.DelegatorIndexByAddress(delegate.DelegatorAddress)
	}

	delegationIdx, ok := s.DelegationIndexByDelegator(delegatorIdx, validatorIdx)

	validator := s.ValidatorAtIndex(validatorIdx)

	if !ok {
		s.AppendDelegation(&types.Delegation{
			ValidatorIndex:                validatorIdx,
			DelegatorIndex:                delegatorIdx,
			Amount:                        delegate.Amount,
			LastAccumulatedRewardPerToken: validator.AccumulatedRewardPerToken,
		})
	} else {
		// @todo implement function withdraw reward process

		delegation := s.DelegationAtIndex(delegationIdx)
		delegation.Amount = delegation.Amount + delegate.Amount
		delegation.LastAccumulatedRewardPerToken = validator.AccumulatedRewardPerToken
		s.UpdateDelegationAtIndex(delegationIdx, delegation)
	}

	return s, nil
}
