package state

import "github.com/unicornultrafoundation/go-hashgraph/types"

// NumValidators returns the current count of validators in the U2U network's validator registry.
func (s *State) NumValidators() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return len(s.validators)
}

// Validators retrieves the list of validators from the U2U network's state.
func (s *State) Validators() []*types.Validator {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.validators
}

// StakedBalances retrieves the staked balance information of delegators in the U2U network.
func (s *State) StakedBalances() []uint64 {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.stakedBalances
}

// Slashing retrieves information about validator slashing events in the U2U network.
func (s *State) Slashing() []uint64 {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.slashing
}

// Withdrawals retrieves the list of withdrawal requests from delegators on the U2U network.
func (s *State) Withdrawals() []*types.Withdrawal {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.withdrawals
}
