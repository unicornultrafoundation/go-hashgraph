package state

import (
	"github.com/unicornultrafoundation/go-hashgraph/consensus/election"
	"github.com/unicornultrafoundation/go-hashgraph/native/idx"
	ptypes "github.com/unicornultrafoundation/go-hashgraph/proto/u2u/types"
)

// SetEpoch updates the current epoch information in the U2U chain's state.
// This function acquires a lock to ensure thread safety while modifying the epoch data.
// It takes an epoch object as input and updates the internal epoch value accordingly.
func (s *State) SetEpoch(epoch idx.Epoch) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.epoch = epoch
}

// SetTime sets the current time in the state.
func (s *State) SetTime(time uint64) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	s.time = time
}

// SetPrevTime sets the previous time in the state.
func (s *State) SetPrevTime(prevTime uint64) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	s.prevTime = prevTime
}

// SetLatestBlock sets the latest block context in the state.
func (s *State) SetLatestBlock(b *ptypes.Block) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	s.latestBlock = b
}

// SetRoots sets the provided slice of election roots and slots in the U2U chain's state.
// It acquires a read lock to ensure thread-safe access to the roots data, sets the roots,
// and then releases the lock.
func (s *State) SetRoots(roots []*election.RootAndSlot) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	s.roots = roots
}

// AppendRoot appends the provided election root and slot to the existing roots in the U2U chain's state.
// It acquires a read lock to ensure thread-safe access to the roots data, appends the root,
// and then releases the lock.
func (s *State) AppendRoot(root *election.RootAndSlot) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	s.roots = append(s.roots, root)
}
