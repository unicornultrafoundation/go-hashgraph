package consensus

import (
	"github.com/unicornultrafoundation/go-hashgraph/state"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

// ProcessExitValidators processes a batch of exit requests for validators in the U2U chain's state.
// It iterates through the provided list of exitValidatorDtos, each containing the address of a validator
// to initiate the exit process for. It calls the ProcessExitValidator function for each validator and
// updates the state accordingly. Returns the updated state and any encountered error.
func ProcessExitValidators(s *state.State, exitValidatorDtos []*types.ExitValidatorDto) (*state.State, error) {
	var err error
	for _, exitValidatorDto := range exitValidatorDtos {
		s, err = ProcessExitValidator(s, exitValidatorDto)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

// ProcessExitValidator initiates the exit process for a specific validator in the U2U chain's state.
// It takes the state s and an exitValidatorDto containing the address of the validator to exit.
// It retrieves the index of the validator using the ValidatorIndexByAddress function, updates the
// validator's ExitEpoch to the next epoch, and then updates the state using UpdateValidatorAtIndex.
// After initiating the exit process, it also processes the undelegation of the validator's stake and
// performs any necessary reward withdrawal. Returns the updated state and any encountered error.
func ProcessExitValidator(s *state.State, exitValidatorDto *types.ExitValidatorDto) (*state.State, error) {
	var err error

	validatorIdx, ok := s.ValidatorIndexByAddress(exitValidatorDto.Address)
	if !ok {
		return s, nil
	}

	delegatorIdx, _ := s.DelegatorIndexByAddress(exitValidatorDto.Address)
	validator := s.ValidatorAtIndex(validatorIdx)
	validator.ExitEpoch = uint64(s.Epoch()) + 1
	s.UpdateValidatorAtIndex(validatorIdx, validator)

	delegationIdx, _ := s.DelegationIndexByDelegator(delegatorIdx, validatorIdx)
	delegation := s.DelegationAtIndex(delegationIdx)

	s, err = ProcessUndelegate(s, &types.UnDelegateDto{
		ValidatorAddress: exitValidatorDto.Address,
		DelegatorAddress: exitValidatorDto.Address,
		Amount:           delegation.Amount,
	})

	if err != nil {
		return nil, err
	}

	// @todo implement withdrawal of validator commission reward process

	return s, nil
}
