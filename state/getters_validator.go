package state

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

// NumValidators returns the current count of validators in the U2U network's validator registry.
// This function acquires a lock to ensure thread safety while accessing the validator count.
func (s *State) NumValidators() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return len(s.validators)
}

// Validators retrieves the list of validators from the U2U network's state.
// This function acquires a lock to ensure thread safety while accessing the validator list.
// It returns a slice of pointers to Validator objects representing validator details.
func (s *State) Validators() []*types.Validator {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.validators
}

// ValidatorIndexByAddress returns the index of a validator by its address and a boolean indicating its existence.
// This function acquires a read lock to ensure safe concurrent access to the validator address map.
func (s *State) ValidatorIndexByAddress(valAddr common.Address) (uint64, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	idx, ok := s.valMap[valAddr]
	return idx, ok
}

// ValidatorAtIndex retrieves a validator by its index from the list.
// This function acquires a read lock to ensure thread safety while accessing validator details.
func (s *State) ValidatorAtIndex(idx uint64) *types.Validator {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.validators[idx]
}

// ReadFromEveryValidator iterates through each validator and applies a function to each.
// The provided function 'f' should accept an index and a clone of the validator.
// If 'f' returns an error, the iteration stops and the error is returned.
func (s *State) ReadFromEveryValidator(f func(idx int, val *types.Validator) error) error {
	s.lock.Lock()
	validators := s.validators
	s.lock.Unlock()

	for i, v := range validators {
		if err := f(i, v.Clone()); err != nil {
			return err
		}
	}
	return nil
}

// StakedBalances retrieves the staked balance information of delegators in the U2U network.
// This function acquires a lock to ensure thread safety while accessing staked balance data.
// It returns a slice of uint64 values representing staked balances of delegators.
func (s *State) StakedBalances() []uint64 {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.stakedBalances
}

// StakedBalanceAtIndex retrieves the staked balance of a delegator at the specified index.
// This function acquires a lock to ensure thread safety while accessing the staked balance data.
func (s *State) StakedBalanceAtIndex(idx uint64) uint64 {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.stakedBalances[idx]
}

// Slashings retrieves information about validator slashing events in the U2U network.
// This function acquires a lock to ensure thread safety while accessing slashing data.
// It returns a slice of uint64 values representing slashing events.
func (s *State) Slashings() []uint64 {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.slashings
}

// SlashingAtIndex retrieves the slashing event value at the specified index.
// This function acquires a lock to ensure thread safety while accessing the slashing data.
func (s *State) SlashingAtIndex(idx uint64) uint64 {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.slashings[idx]
}

// Withdrawals retrieves the list of withdrawal requests from delegators on the U2U network.
// This function acquires a lock to ensure thread safety while accessing withdrawal data.
// It returns a slice of pointers to Withdrawal objects representing withdrawal requests.
func (s *State) Withdrawals() []*types.Withdrawal {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.withdrawals
}

// WithdrawalAtIndex retrieves a withdrawal request by its index from the list.
// This function acquires a lock to ensure thread safety while accessing withdrawal details.
func (s *State) WithdrawalAtIndex(idx uint64) *types.Withdrawal {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.withdrawals[idx]
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

// DelegationAtIndex retrieves a delegation by its index from the list.
// This function acquires a lock to ensure thread safety while accessing delegation details.
func (s *State) DelegationAtIndex(idx uint64) *types.Delegation {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.delegations[idx]
}

// DelegationIndexByDelegator retrieves the index of a delegation by its delegator and validator indexes.
// This function acquires a lock to ensure thread safety while accessing delegation data.
// It returns the index and a boolean indicating whether the delegation was found.
func (s *State) DelegationIndexByDelegator(delIdx uint64, valIdx uint64) (uint64, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	idx, ok := s.delValMap[valIdx][delIdx]
	return idx, ok
}

// AccumulatedRewards retrieves the list of accumulated rewards in the U2U network's state.
// This function acquires a lock to ensure thread safety while accessing accumulated rewards data.
// It returns a slice of uint64 values representing accumulated rewards for validators.
func (s *State) AccumulatedRewards() []uint64 {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.accumulatedRewards
}

// AccumulatedRewardByIndex retrieves the accumulated reward value at the specified index.
// This function acquires a lock to ensure thread safety while accessing the accumulated rewards data.
func (s *State) AccumulatedRewardByIndex(idx uint64) uint64 {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.accumulatedRewards[idx]
}

// Delegators retrieves the list of delegators from the U2U chain's state.
// This function acquires a lock to ensure thread safety while accessing the delegator list.
// It returns a slice of pointers to Delegator objects representing delegator details.
func (s *State) Delegators() []*types.Delegator {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.delegators
}

// DelegatorIndexByAddress retrieves the index of a delegator by their address and a boolean indicating its existence.
// This function acquires a lock to ensure thread safety while accessing the delegator address map.
func (s *State) DelegatorIndexByAddress(addr common.Address) (uint64, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	idx, ok := s.delMap[addr]
	return idx, ok
}

// DelegatorAtIndex retrieves a delegator by its index from the list.
// This function acquires a lock to ensure thread safety while accessing delegator details.
func (s *State) DelegatorAtIndex(idx uint64) *types.Delegator {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.delegators[idx]
}

// WithdrawalRewards retrieves the list of withdrawal rewards from the U2U chain's state.
// This function acquires a lock to ensure thread safety while accessing the withdrawal reward data.
// It returns a slice of pointers to WithdrawalReward objects representing withdrawal reward details.
// Make sure to use this function within a context that handles concurrent access properly.
func (s *State) WithdrawalRewards() []*types.WithdrawalReward {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.withdrawalRewards
}

// WithdrawalRewardAtIndex retrieves a specific withdrawal reward by its index in the list from the U2U chain's state.
// This function acquires a lock to ensure thread safety while accessing the withdrawal reward data.
// It takes an index as input and returns the WithdrawalReward pointer at that index.
// Make sure to use this function within a context that handles concurrent access properly.
func (s *State) WithdrawalRewardAtIndex(idx uint64) *types.WithdrawalReward {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.withdrawalRewards[idx]
}
