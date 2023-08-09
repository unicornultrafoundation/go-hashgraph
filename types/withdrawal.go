package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/unicornultrafoundation/go-hashgraph/native/idx"
)

type Withdrawal struct {
	ValidatorIndex    uint64
	Address           common.Address
	Amount            uint64
	WithdrawableEpoch idx.Epoch
}
