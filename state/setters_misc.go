package state

import "github.com/unicornultrafoundation/go-hashgraph/native/idx"

// SetEpoch updates the current epoch information in the U2U chain's state.
// This function acquires a lock to ensure thread safety while modifying the epoch data.
// It takes an epoch object as input and updates the internal epoch value accordingly.
func (s *State) SetEpoch(epoch idx.Epoch) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.epoch = epoch
}
