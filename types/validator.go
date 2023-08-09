package types

import "github.com/ethereum/go-ethereum/common"

type Validator struct {
	Address         common.Address
	Slashed         bool
	ActivationEpoch uint64
	ExitEpoch       uint64
}
