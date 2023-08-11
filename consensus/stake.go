package consensus

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unicornultrafoundation/go-hashgraph/state"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

// ProcessStakes processes stake events in the U2U chain's state.
// It iterates through the provided stakeDtos and processes each stake event using ProcessStake function.
// Returns the updated state after processing all stakes and any encountered error.
func ProcessStakes(s *state.State, stakeDtos []*types.StakeDto) (*state.State, error) {
	var err error
	for _, stakeDto := range stakeDtos {
		s, err = ProcessStake(s, stakeDto)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

// ProcessStake processes an individual stake event in the U2U chain's state.
// It first checks if the delegation is a self-stake by comparing the validator and delegator addresses.
// If it's a self-stake and the validator doesn't exist, it initiates the createValidator process.
// Otherwise, it processes the stake event by either creating a new delegation or updating an existing one.
// Returns the updated state after processing the stake event and any encountered error.
func ProcessStake(s *state.State, stakeDto *types.StakeDto) (*state.State, error) {
	if isSelfStake(stakeDto.ValidatorAddress, stakeDto.DelegatorAddress) {
		_, ok := s.ValidatorIndexByAddress(stakeDto.ValidatorAddress)
		if !ok {
			return ProcessCreateValidator(s, &types.CreateValidatorDto{
				Address: stakeDto.ValidatorAddress,
				Amount:  stakeDto.Amount,
			})
		}
	}
	return ProcessDelegate(s, &types.DelegateDto{
		ValidatorAddress: stakeDto.ValidatorAddress,
		DelegatorAddress: stakeDto.DelegatorAddress,
		Amount:           stakeDto.Amount,
	})
}

// isSelfStake checks if a delegation is a self-stake by comparing the validator and delegator addresses.
// It returns true if they are equal, indicating a self-stake scenario.
func isSelfStake(val common.Address, del common.Address) bool {
	return bytes.Equal(val[:], del[:])
}
