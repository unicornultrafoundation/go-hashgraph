package consensus

import (
	"github.com/unicornultrafoundation/go-hashgraph/state"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

// ProcessSlashings processes slashing events in the U2U chain's state.
// It iterates through the provided slashingDtos and processes each slashing event using ProcessSlashing function.
// Returns the updated state after processing all slashings and any encountered error.
func ProcessSlashings(s *state.State, slashingDtos []*types.SlashingDto) (*state.State, error) {
	var err error
	for _, slashingDto := range slashingDtos {
		s, err = ProcessSlashing(s, slashingDto)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

// ProcessSlashing processes an individual slashing event in the U2U chain's state.
// It first checks if the validator associated with the slashingDto exists in the state.
// If the validator does not exist, it returns the state as is.
// Otherwise, it marks the validator as slashed and applies a penalty to their staked balance.
// The penalty quotient is used to determine the amount of penalty applied.
// Finally, it initiates the exit process for the slashed validator using ProcessExitValidator function.
// Returns the updated state after processing the slashing event and any encountered error.
func ProcessSlashing(s *state.State, slashingDto *types.SlashingDto) (*state.State, error) {
	validatorIdx, ok := s.ValidatorIndexByAddress(slashingDto.ValidatorAddress)
	if !ok {
		return s, nil
	}

	validator := s.ValidatorAtIndex(validatorIdx)
	validator.Slashed = true

	// @todo need move to config
	penaltyQuotient := uint64(128)

	stakedBalance := s.StakedBalanceAtIndex(validatorIdx)
	s.UpdateStakedBalanceAtIndex(validatorIdx, stakedBalance-stakedBalance/penaltyQuotient)

	return ProcessExitValidator(s, &types.ExitValidatorDto{Address: slashingDto.ValidatorAddress})
}
