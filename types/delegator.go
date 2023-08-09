package types

import "github.com/ethereum/go-ethereum/common"

type Delegation struct {
	ValidatorIndex                uint64
	DelegatorIndex                uint64
	Amount                        uint64
	LastAccumulatedRewardPerToken uint64
}

type Delegator struct {
	Address common.Address
}

type DelegateDto struct {
	ValidatorAddress common.Address
	DelegatorAddress common.Address
	Amount           uint64
}

type UnDelegateDto struct {
	ValidatorAddress common.Address
	DelegatorAddress common.Address
	Amount           uint64
}
