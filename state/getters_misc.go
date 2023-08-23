package state

import (
	"github.com/unicornultrafoundation/go-hashgraph/hash"
	"github.com/unicornultrafoundation/go-hashgraph/native/idx"
	ptypes "github.com/unicornultrafoundation/go-hashgraph/proto/u2u/types"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

// Epoch retrieves the epoch value of the current U2U chain state.
// This function acquires a read lock to ensure thread safety while accessing the epoch in the State.
// It returns an idx.Epoch representing the current epoch of the U2U chain.
// Ensure that you use this function within a context that handles concurrent access appropriately.
func (s *State) Epoch() idx.Epoch {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.epoch
}

// Time returns the current time in the state.
func (s *State) Time() uint64 {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.time
}

// PrevTime returns the previous time in the state.
func (s *State) PrevTime() uint64 {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.prevTime
}

// LatestBlock returns the latest block context in the state.
func (s *State) LatestBlock() *ptypes.Block {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.latestBlock
}

// Roots returns a slice of election roots and slots stored in the U2U chain's state.
// It acquires a read lock to ensure thread-safe access to the roots data,
// retrieves the roots from the state, and then releases the lock.
// Returns the slice of election roots and slots.
func (s *State) Roots() []*types.RootAndSlot {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.roots
}

// ConfirmedEventByHash retrieves a confirmed event by its hash from the state.
// It acquires a read lock to ensure thread-safe access to the confirmed events map.
// It looks up the index of the event hash in the map and if found, returns the corresponding confirmed event.
// If the event hash is not present in the map, it returns nil.
func (s *State) ConfirmedEventByHash(hash hash.Event) *types.ConfirmedEvent {
	s.lock.RLock()
	defer s.lock.RUnlock()

	idx, ok := s.ceMap[hash]
	if !ok {
		return nil
	}
	return s.confirmedEvents[idx]
}

func (s *State) LastDecidedFrame() idx.Frame {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.lastDecidedFrame
}
