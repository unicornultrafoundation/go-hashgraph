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

// Slashings retrieves information about validator slashing events in the U2U network.
func (s *State) Slashings() []uint64 {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.slashings
}

// Withdrawals retrieves the list of withdrawal requests from delegators on the U2U network.
func (s *State) Withdrawals() []*types.Withdrawal {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.withdrawals
}

// Delegations retrieves the list of delegations from the U2U chain's state.
// This function acquires a lock to ensure thread safety while accessing the delegation data.
// It returns a slice of pointers to Delegation objects representing delegation details.
// To maintain data integrity, make sure to use this function within a context that handles concurrent access properly.
func (s *State) Delegations() []*types.Delegation {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.delegations
}
