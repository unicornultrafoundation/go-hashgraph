package state

import "github.com/unicornultrafoundation/go-hashgraph/native/idx"

// Epoch retrieves the epoch value of the current U2U chain state.
// This function acquires a read lock to ensure thread safety while accessing the epoch in the State.
// It returns an idx.Epoch representing the current epoch of the U2U chain.
// Ensure that you use this function within a context that handles concurrent access appropriately.
func (s *State) Epoch() idx.Epoch {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.epoch
}
