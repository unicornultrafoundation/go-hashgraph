package consensus

import (
	"github.com/unicornultrafoundation/go-hashgraph/state"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

// ProcessCreateValidators processes a list of CreateValidatorDto objects to create validators and perform delegations.
// It iterates through each CreateValidatorDto in the list and invokes ProcessCreateValidator for each.
// If an error occurs during processing, the function returns an error.
// Otherwise, it returns the updated state after processing all create validator actions.
func ProcessCreateValidators(s *state.State, createValidatorDtos []*types.CreateValidatorDto) (*state.State, error) {
	var err error
	for _, createValidatorDto := range createValidatorDtos {
		s, err = ProcessCreateValidator(s, createValidatorDto)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

// ProcessCreateValidator creates a new validator and delegates a specified amount to it.
// It appends a new validator to the state using the information from the CreateValidatorDto.
// Then, it invokes ProcessDelegate to perform the delegation using the provided CreateValidatorDto.
// The function returns the updated state after processing the create validator action.
func ProcessCreateValidator(s *state.State, createValidatorDto *types.CreateValidatorDto) (*state.State, error) {
	s.AppendValidator(&types.Validator{
		Address: createValidatorDto.Address,
		Slashed: false,
	})

	s.AppendStakedBalance(0)

	return ProcessDelegate(s, &types.DelegateDto{
		ValidatorAddress: createValidatorDto.Address,
		DelegatorAddress: createValidatorDto.Address,
		Amount:           createValidatorDto.Amount,
	})
}
