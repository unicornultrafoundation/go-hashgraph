package types

import "github.com/ethereum/go-ethereum/common"

type Delegator struct {
	Address      common.Address
	ValidatorIds []uint64
}

type Delegation struct {
	Amount                        uint64
	LastAccumulatedRewardPerToken uint64
}
