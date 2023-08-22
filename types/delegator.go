package types

import "github.com/ethereum/go-ethereum/common"

// DelegateDto is a data transfer object (DTO) that holds information about a delegation action.
type DelegateDto struct {
	ValidatorAddress common.Address // Ethereum address of the validator being delegated to.
	DelegatorAddress common.Address // Ethereum address of the delegator making the delegation.
	Amount           uint64         // Amount of tokens being delegated.
}

// UnDelegateDto is a data transfer object (DTO) for handling undelegation actions.
type UnDelegateDto struct {
	ValidatorAddress common.Address // Ethereum address of the validator being undelegated from.
	DelegatorAddress common.Address // Ethereum address of the delegator making the undelegation.
	Amount           uint64         // Amount of tokens being undelegated.
}

// WithdrawalRewardDto represents a data transfer object for withdrawal rewards.
type WithdrawalRewardDto struct {
	ValidatorAddress common.Address // Address of the validator receiving the reward.
	DelegatorAddress common.Address // Address of the delegator receiving the reward.
}

// StakeDto represents a data transfer object for staking tokens.
type StakeDto struct {
	ValidatorAddress common.Address // Address of the validator to stake tokens with.
	DelegatorAddress common.Address // Address of the delegator staking tokens.
	Amount           uint64         // Amount of tokens to be staked.
}
