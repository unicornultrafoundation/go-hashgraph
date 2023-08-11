package types

import "github.com/ethereum/go-ethereum/common"

// Validator defines a struct that represents the attributes and status of a validator in the network.
type Validator struct {
	Address                common.Address // Ethereum address of the validator.
	Slashed                bool           // Indicates if the validator has been slashed for misconduct.
	ActivationEpoch        uint64         // Epoch in which the validator became active.
	ExitEpoch              uint64         // Epoch in which the validator intends to exit the network.
	EffectiveStakedBalance uint64         // Effective staked balance of the validator.
}

// Clone creates a deep copy of the Validator object.
func (v *Validator) Clone() *Validator {
	return &Validator{
		Address:                v.Address,
		Slashed:                v.Slashed,
		ActivationEpoch:        v.ActivationEpoch,
		ExitEpoch:              v.ExitEpoch,
		EffectiveStakedBalance: v.EffectiveStakedBalance,
	}
}

// CreateValidatorDto is a data transfer object (DTO) that holds information for creating a new validator.
type CreateValidatorDto struct {
	Address common.Address // Ethereum address of the new validator.
	Amount  uint64         // Amount of tokens being staked for validator creation.
}

// ExitValidatorDto is a data transfer object representing a request to initiate the exit process for a validator.
// It contains the validator's address that is targeted for the exit process.
type ExitValidatorDto struct {
	Address common.Address // The address of the validator to be exited.
}

// SlashingDto represents a data transfer object (DTO) containing information about a slashing event.
type SlashingDto struct {
	ValidatorAddress common.Address // Address of the validator associated with the slashing event.
}
