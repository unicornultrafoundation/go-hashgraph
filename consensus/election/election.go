package election

import (
	"github.com/unicornultrafoundation/go-hashgraph/consensus/helpers"
	"github.com/unicornultrafoundation/go-hashgraph/hash"
	"github.com/unicornultrafoundation/go-hashgraph/native/idx"
	"github.com/unicornultrafoundation/go-hashgraph/state"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

type (
	// Election is cached data of election algorithm.
	Election struct {
		state *state.State

		// election state
		decidedRoots map[idx.ValidatorID]voteValue // decided roots at "frameToDecide"
		votes        map[voteID]voteValue

		// external world
		observe       ForklessCauseFn
		getFrameRoots GetFrameRootsFn
	}

	// ForklessCauseFn returns true if event A is forkless caused by event B
	ForklessCauseFn func(a hash.Event, b hash.Event) bool
	// GetFrameRootsFn returns all the roots in the specified frame
	GetFrameRootsFn func(f idx.Frame) []types.RootAndSlot
)

type voteID struct {
	fromRoot     types.RootAndSlot
	forValidator idx.ValidatorID
}
type voteValue struct {
	decided      bool
	yes          bool
	observedRoot hash.Event
}

// Res defines the final election result, i.e. decided frame
type Res struct {
	Frame idx.Frame
	Event hash.Event
}

// New election context
func New(
	st *state.State,
	forklessCauseFn ForklessCauseFn,
) *Election {
	el := &Election{
		state:   st,
		observe: forklessCauseFn,
	}
	return el
}

// return root slots which are not within el.decidedRoots
func (el *Election) notDecidedRoots() []idx.ValidatorID {
	activeIndices, _ := helpers.ActiveValidatorIndices(el.state, el.state.Epoch())
	valCount := len(activeIndices)
	notDecidedRoots := make([]idx.ValidatorID, 0, valCount)

	for _, validator := range activeIndices {
		if _, ok := el.decidedRoots[validator]; !ok {
			notDecidedRoots = append(notDecidedRoots, validator)
		}
	}
	if len(notDecidedRoots)+len(el.decidedRoots) != valCount { // sanity check
		panic("Mismatch of roots")
	}
	return notDecidedRoots
}

// observedRoots returns all the roots at the specified frame which do forkless cause the specified root.
func (el *Election) observedRoots(root hash.Event, frame idx.Frame) []types.RootAndSlot {
	valCount, _ := helpers.ActiveValidatorCount(el.state, el.state.Epoch())
	observedRoots := make([]types.RootAndSlot, 0, valCount)

	frameRoots := el.getFrameRoots(frame)
	for _, frameRoot := range frameRoots {
		if el.observe(root, frameRoot.Hash) {
			observedRoots = append(observedRoots, frameRoot)
		}
	}
	return observedRoots
}

func (el *Election) observedRootsMap(root hash.Event, frame idx.Frame) map[idx.ValidatorID]types.RootAndSlot {
	valCount, _ := helpers.ActiveValidatorCount(el.state, el.state.Epoch())
	observedRootsMap := make(map[idx.ValidatorID]types.RootAndSlot, valCount)

	frameRoots := el.getFrameRoots(frame)
	for _, frameRoot := range frameRoots {
		if el.observe(root, frameRoot.Hash) {
			observedRootsMap[frameRoot.Slot.Validator] = frameRoot
		}
	}
	return observedRootsMap
}
