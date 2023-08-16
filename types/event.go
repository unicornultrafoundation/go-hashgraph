package types

import (
	"github.com/unicornultrafoundation/go-hashgraph/hash"
	"github.com/unicornultrafoundation/go-hashgraph/native/idx"
)

type FinalEventDto struct {
	Block  *BlockCtx
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
