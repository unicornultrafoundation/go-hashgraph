package consensus

import (
	"github.com/unicornultrafoundation/go-hashgraph/consensus/dagidx"
	"github.com/unicornultrafoundation/go-hashgraph/hash"
	"github.com/unicornultrafoundation/go-hashgraph/internal/dag"
	"github.com/unicornultrafoundation/go-hashgraph/internal/idx"
	"github.com/unicornultrafoundation/go-hashgraph/internal/pos"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

var _ types.Consensus = (*Consensus)(nil)

type DagIndex interface {
	dagidx.VectorClock
	dagidx.ForklessCause
}

// Hashgraph performs events ordering and detects cheaters
// It's a wrapper around Orderer, which adds features which might potentially be application-specific:
// confirmed events traversal, cheaters detection.
// Use this structure if need a general-purpose consensus. Instead, use lower-level consensus.Orderer.
type Consensus struct {
	*Orderer
	dagIndex      DagIndex
	uniqueDirtyID uniqueID
	callback      types.ConsensusCallbacks
}

// NewConsensus creates Consensus instance.
func NewConsensus(store *Store, input EventSource, dagIndex DagIndex, crit func(error), config Config) *Consensus {
	p := &Consensus{
		Orderer:  NewOrderer(store, input, dagIndex, crit, config),
		dagIndex: dagIndex,
	}

	return p
}

func (p *Consensus) confirmEvents(frame idx.Frame, event hash.Event, onEventConfirmed func(dag.Event)) error {
	err := p.dfsSubgraph(event, func(e dag.Event) bool {
		decidedFrame := p.store.GetEventConfirmedOn(e.ID())
		if decidedFrame != 0 {
			return false
		}
		// mark all the walked events as confirmed
		p.store.SetEventConfirmedOn(e.ID(), frame)
		if onEventConfirmed != nil {
			onEventConfirmed(e)
		}
		return true
	})
	return err
}

func (p *Consensus) applyEvent(decidedFrame idx.Frame, event hash.Event) *pos.Validators {
	eventVecClock := p.dagIndex.GetMergedHighestBefore(event)

	validators := p.store.GetValidators()
	// cheaters are ordered deterministically
	cheaters := make([]idx.ValidatorID, 0, validators.Len())
	for creatorIdx, creator := range validators.SortedIDs() {
		if eventVecClock.Get(idx.Validator(creatorIdx)).IsForkDetected() {
			cheaters = append(cheaters, creator)
		}
	}

	if p.callback.BeginBlock == nil {
		return nil
	}
	blockCallback := p.callback.BeginBlock(&types.Block{
		Event:    event,
		Cheaters: cheaters,
	})

	// traverse newly confirmed events
	err := p.confirmEvents(decidedFrame, event, blockCallback.ApplyEvent)
	if err != nil {
		p.crit(err)
	}

	if blockCallback.EndBlock != nil {
		return blockCallback.EndBlock()
	}
	return nil
}

func (p *Consensus) Bootstrap(callback types.ConsensusCallbacks) error {
	return p.BootstrapWithOrderer(callback, p.OrdererCallbacks())
}

func (p *Consensus) BootstrapWithOrderer(callback types.ConsensusCallbacks, ordererCallbacks OrdererCallbacks) error {
	err := p.Orderer.Bootstrap(ordererCallbacks)
	if err != nil {
		return err
	}
	p.callback = callback
	return nil
}

func (p *Consensus) OrdererCallbacks() OrdererCallbacks {
	return OrdererCallbacks{
		ApplyEvent: p.applyEvent,
	}
}
