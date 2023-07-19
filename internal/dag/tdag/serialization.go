package tdag

import (
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/unicornultrafoundation/go-hashgraph/hash"
	"github.com/unicornultrafoundation/go-hashgraph/internal/idx"
)

type TestEventMarshaling struct {
	Epoch idx.Epoch
	Seq   idx.Event

	Frame idx.Frame

	Creator idx.ValidatorID

	Parents hash.Events

	Lamport idx.Lamport

	ID   hash.Event
	Name string
}

// EventToBytes serializes events
func (e *TestEvent) Bytes() []byte {
	b, _ := rlp.EncodeToBytes(&TestEventMarshaling{
		Epoch:   e.Epoch(),
		Seq:     e.Seq(),
		Frame:   e.Frame(),
		Creator: e.Creator(),
		Parents: e.Parents(),
		Lamport: e.Lamport(),
		ID:      e.ID(),
		Name:    e.Name,
	})
	return b
}
