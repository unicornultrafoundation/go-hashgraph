package types

import (
	"github.com/unicornultrafoundation/go-hashgraph/native/idx"
)

type Withdrawal struct {
	ValidatorIndex uint64
	DelegatorIndex uint64
	Amount         uint64
	Epoch          idx.Epoch
}
