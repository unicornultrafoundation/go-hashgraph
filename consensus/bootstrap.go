package consensus

import (
	"errors"

	"github.com/unicornultrafoundation/go-hashgraph/consensus/election"
	"github.com/unicornultrafoundation/go-hashgraph/native/idx"
)

const (
	FirstFrame = idx.Frame(1)
	FirstEpoch = idx.Epoch(1)
)

// Bootstrap restores consensus's state from store.
func (p *Orderer) Bootstrap(callback OrdererCallbacks) error {
	if p.election != nil {
		return errors.New("already bootstrapped")
	}
	p.callback = callback
	p.election = election.New(p.state, p.dagIndex.ForklessCause)
	_, err := p.bootstrapElection()
	return err
}
