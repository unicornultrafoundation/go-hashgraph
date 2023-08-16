package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/unicornultrafoundation/go-hashgraph/native/idx"
)

// Validator defines a struct that represents the attributes and status of a validator in the network.
type Validator struct {
	Address                   common.Address // Ethereum address of the validator.
	Slashed                   bool           // Indicates if the validator has been slashed for misconduct.
	ActivationEpoch           idx.Epoch      // Epoch in which the validator became active.
	ExitEpoch                 idx.Epoch      // Epoch in which the validator intends to exit the network.
	EffectiveStakedBalance    uint64         // Effective staked balance of the validator.
	AccumulatedRewardPerToken uint64
	CommissionReward          uint64
	LastBlockId               idx.Block
	LastOnlineTime            uint64
	Uptime                    uint64
}

// Clone creates a deep copy of the Validator object.
func (v *Validator) Clone() *Validator {
	return &Validator{
		Address:                v.Address,
		Slashed:                v.Slashed,
		ActivationEpoch:        v.ActivationEpoch,
		ExitEpoch:              v.ExitEpoch,
		EffectiveStakedBalance: v.EffectiveStakedBalance,
		LastBlockId:            v.LastBlockId,
		LastOnlineTime:         v.LastOnlineTime,
		Uptime:                 v.Uptime,
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

type ValidatorEpochMetric struct {
	Missed idx.Block
	Uptime uint64
}
