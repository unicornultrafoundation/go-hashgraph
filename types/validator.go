package types

import "github.com/ethereum/go-ethereum/common"

// Validator defines a struct that represents the attributes and status of a validator in the network.
type Validator struct {
	Address         common.Address // Ethereum address of the validator.
	Slashed         bool           // Indicates if the validator has been slashed for misconduct.
	ActivationEpoch uint64         // Epoch in which the validator became active.
	ExitEpoch       uint64         // Epoch in which the validator intends to exit the network.
}

// CreateValidatorDto is a data transfer object (DTO) that holds information for creating a new validator.
type CreateValidatorDto struct {
	Address common.Address // Ethereum address of the new validator.
	Amount  uint64         // Amount of tokens being staked for validator creation.
}
