package state

import (
	"github.com/unicornultrafoundation/go-hashgraph/state/stateutil"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

// SetValidators updates the list of active validators in the U2U chain state.
// This function acquires a lock to ensure thread safety while modifying validator-related data in the State.
// It takes a slice of Validator pointers as input and updates the internal validator list and index map accordingly.
// Ensure that you use this function within a context that handles concurrent access appropriately.
func (s *State) SetValidators(vals []*types.Validator) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.validators = vals
	s.valMap = stateutil.ValidatorIndexMap(vals)
}

// AppendValidator adds a new validator to the list of active validators.
// This function acquires a lock to ensure thread safety while modifying the validator list and index map in the State.
func (s *State) AppendValidator(val *types.Validator) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.validators = append(s.validators, val)
	s.valMap[val.Address] = uint64(len(s.validators) - 1)
}

// UpdateValidatorAtIndex updates the details of a validator at the specified index.
// This function acquires a lock to ensure thread safety while modifying the validator list in the State.
func (s *State) UpdateValidatorAtIndex(idx uint64, val *types.Validator) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.validators[idx] = val
}

// SetDelegators updates the list of delegators in the U2U chain state.
// This function acquires a lock to ensure thread safety while modifying delegator-related data in the State.
func (s *State) SetDelegators(dels []*types.Delegator) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.delegators = dels
	s.delMap = stateutil.DelegatorIndexMap(dels)
}

// AppendDelegator adds a new delegator to the list of delegators.
// This function acquires a lock to ensure thread safety while modifying the delegator list and index map in the State.
func (s *State) AppendDelegator(del *types.Delegator) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.delegators = append(s.delegators, del)
	s.valMap[del.Address] = uint64(len(s.delegators) - 1)
}

// UpdateDelegatorAtIndex updates the details of a delegator at the specified index.
// This function acquires a lock to ensure thread safety while modifying the delegator list in the State.
func (s *State) UpdateDelegatorAtIndex(idx uint64, del *types.Delegator) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.delegators[idx] = del
}

// SetDelegations updates the list of delegations in the U2U chain state.
// This function acquires a lock to ensure thread safety while modifying delegation-related data in the State.
func (s *State) SetDelegations(delegations []*types.Delegation) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.delegations = delegations
}

// AppendDelegation adds a new delegation to the list of delegations.
// This function acquires a lock to ensure thread safety while modifying the delegation list and index map in the State.
func (s *State) AppendDelegation(del *types.Delegation) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.delegations = append(s.delegations, del)
	s.delValMap[del.ValidatorIndex][del.DelegatorIndex] = uint64(len(s.delegations) - 1)
}

// UpdateDelegationAtIndex updates the details of a delegation at the specified index.
// This function acquires a lock to ensure thread safety while modifying the delegation list in the State.
func (s *State) UpdateDelegationAtIndex(idx uint64, delegation *types.Delegation) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.delegations[idx] = delegation
}

// SetStakedBalances updates the staked balance list in the U2U chain state.
// This function acquires a lock to ensure thread safety while modifying staked balance data in the State.
func (s *State) SetStakedBalances(bals []uint64) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.stakedBalances = bals
}

// AppendStakedBalances adds a new staked balance entry to the list.
// This function acquires a lock to ensure thread safety while modifying the staked balance list in the State.
func (s *State) AppendStakedBalance(bal uint64) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.stakedBalances = append(s.stakedBalances, bal)
}

// UpdateStakedBalanceAtIndex updates the staked balance at the specified index.
// This function acquires a lock to ensure thread safety while modifying the staked balance list in the State.
func (s *State) UpdateStakedBalanceAtIndex(idx uint64, bal uint64) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.stakedBalances[idx] = bal
}

// SetAccumulatedRewards updates the accumulated rewards list in the U2U chain state.
// This function acquires a lock to ensure thread safety while modifying accumulated rewards data in the State.
func (s *State) SetAccumulatedRewards(accumulatedRewards []uint64) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.accumulatedRewards = accumulatedRewards
}

// AppendAccumulatedReward adds a new accumulated reward entry to the list.
// This function acquires a lock to ensure thread safety while modifying the accumulated rewards list in the State.
func (s *State) AppendAccumulatedReward(accumulatedReward uint64) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.accumulatedRewards = append(s.accumulatedRewards, accumulatedReward)
}

// UpdateAccumulatedRewardAtIndex updates the accumulated reward at the specified index.
// This function acquires a lock to ensure thread safety while modifying the accumulated rewards list in the State.
func (s *State) UpdateAccumulatedRewardAtIndex(idx uint64, accumulatedReward uint64) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.accumulatedRewards[idx] = accumulatedReward
}

// SetSlashings updates the slashing information list in the U2U chain state.
// This function acquires a lock to ensure thread safety while modifying slashing data in the State.
func (s *State) SetSlashings(slashings []uint64) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.slashings = slashings
}

// UpdateSlashingAtIndex updates the slashing value at the specified index.
// This function acquires a lock to ensure thread safety while modifying the slashing data in the State.
func (s *State) UpdateSlashingAtIndex(idx uint64, val uint64) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.slashings[idx] = val
}

// SetWithdrawals updates the withdrawal requests list in the U2U chain state.
// This function acquires a lock to ensure thread safety while modifying withdrawal data in the State.
func (s *State) SetWithdrawals(withdrawals []*types.Withdrawal) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.withdrawals = withdrawals
}

// AppendWithdrawal adds a new withdrawal request to the list.
// This function acquires a lock to ensure thread safety while modifying the withdrawal list in the State.
func (s *State) AppendWithdrawal(withdrawal *types.Withdrawal) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.withdrawals = append(s.withdrawals, withdrawal)
}
