package consensus

import (
	"github.com/unicornultrafoundation/go-hashgraph/consensus/epoch"
	"github.com/unicornultrafoundation/go-hashgraph/hash"
	"github.com/unicornultrafoundation/go-hashgraph/native/idx"
	"github.com/unicornultrafoundation/go-hashgraph/native/pos"
)

// onFrameDecided moves LastDecidedFrameN to frame.
// It includes: moving current decided frame, txs ordering and execution, epoch sealing.
func (p *Orderer) onFrameDecided(frame idx.Frame, event hash.Event) (bool, error) {
	// // new checkpoint
	// var newValidators *pos.Validators
	// if p.callback.ApplyEvent != nil {
	// 	newValidators = p.callback.ApplyEvent(frame, event)
	// }

	// lastDecidedFrame := p.state.LastDecidedFrame()
	// if newValidators != nil {
	// 	lastDecidedFrame = FirstFrame - 1
	// 	err := p.sealEpoch(newValidators)
	// 	if err != nil {
	// 		return true, err
	// 	}
	// 	p.election.Reset(newValidators, FirstFrame)
	// } else {
	// 	lastDecidedFrame = frame
	// 	p.election.Reset(p.store.GetValidators(), frame+1)
	// }
	// p.state.SetLastDecidedState(lastDecidedFrame)
	// return newValidators != nil, nil
	return true, nil
}

func (p *Orderer) sealEpoch(newValidators *pos.Validators) error {
	_, err := epoch.ProcessFinalUpdates(p.state)
	if err != nil {
		return err
	}
	// store state
	return nil
}
