package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/unicornultrafoundation/go-hashgraph/native/idx"
)

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

type ValidatorEpochMetric struct {
	Missed idx.Block
	Uptime uint64
}
