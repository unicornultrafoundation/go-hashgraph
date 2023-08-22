package types

import (
	"github.com/unicornultrafoundation/go-hashgraph/hash"
	"github.com/unicornultrafoundation/go-hashgraph/native/idx"
	ptypes "github.com/unicornultrafoundation/go-hashgraph/proto/u2u/types"
)

type FinalEventDto struct {
	Block  *ptypes.Block
	Events []*EventInfoDto
}

type EventInfoDto struct {
	ID   hash.Event
	Time uint64
}

type BlockCtx struct {
	Id    idx.Block
	Time  uint64
	Event hash.Event
}
