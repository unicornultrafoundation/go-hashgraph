package state

import (
	"github.com/unicornultrafoundation/go-hashgraph/native/idx"
	ptypes "github.com/unicornultrafoundation/go-hashgraph/proto/u2u/types"
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
