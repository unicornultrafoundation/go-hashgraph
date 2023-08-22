package state

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	ptypes "github.com/unicornultrafoundation/go-hashgraph/proto/u2u/types"
	"github.com/unicornultrafoundation/go-hashgraph/state/stateutil"
)

// SetValidators updates the list of active validators in the U2U chain state.
// This function acquires a lock to ensure thread safety while modifying validator-related data in the State.
// It takes a slice of Validator pointers as input and updates the internal validator list and index map accordingly.
// Ensure that you use this function within a context that handles concurrent access appropriately.
func (s *State) SetValidators(vals []*ptypes.Validator) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.validators = vals
	s.valMap = stateutil.ValidatorIndexMap(vals)
}

// AppendValidator adds a new validator to the list of active validators.
// This function acquires a lock to ensure thread safety while modifying the validator list and index map in the State.
func (s *State) AppendValidator(val *ptypes.Validator) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.validators = append(s.validators, val)
	s.valMap[common.BytesToAddress(val.Address)] = uint64(len(s.validators) - 1)
}

// ApplyToEveryValidator applies the given function to each validator in the state's validator list.
// The provided function takes the index of the validator and the validator itself as arguments.
// It returns a boolean indicating whether the validator was changed, the updated validator, and any encountered error.
// The function locks the state before accessing the validators and releases the lock after updating them.
func (s *State) ApplyToEveryValidator(f func(idx int, val *ptypes.Validator) (bool, *ptypes.Validator, error)) error {
	s.lock.Lock()
	v := s.validators
	s.lock.Unlock()

	for i, val := range v {
		changed, newVal, err := f(i, val)
		if err != nil {
			return err
		}
		if changed {
			v[i] = newVal
		}
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	s.validators = v
	return nil
}

// UpdateValidatorAtIndex updates the details of a validator at the specified index.
// This function acquires a lock to ensure thread safety while modifying the validator list in the State.
func (s *State) UpdateValidatorAtIndex(idx uint64, val *ptypes.Validator) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.validators[idx] = val
}

// SetDelegators updates the list of delegators in the U2U chain state.
// This function acquires a lock to ensure thread safety while modifying delegator-related data in the State.
func (s *State) SetDelegators(dels []*ptypes.Delegator) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.delegators = dels
	s.delMap = stateutil.DelegatorIndexMap(dels)
}

// AppendDelegator adds a new delegator to the list of delegators.
// This function acquires a lock to ensure thread safety while modifying the delegator list and index map in the State.
func (s *State) AppendDelegator(del *ptypes.Delegator) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.delegators = append(s.delegators, del)
	s.valMap[common.BytesToAddress(del.Address)] = uint64(len(s.delegators) - 1)
}

// UpdateDelegatorAtIndex updates the details of a delegator at the specified index.
// This function acquires a lock to ensure thread safety while modifying the delegator list in the State.
func (s *State) UpdateDelegatorAtIndex(idx uint64, del *ptypes.Delegator) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.delegators[idx] = del
}

// SetDelegations updates the list of delegations in the U2U chain state.
// This function acquires a lock to ensure thread safety while modifying delegation-related data in the State.
func (s *State) SetDelegations(delegations []*ptypes.Delegation) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.delegations = delegations
}

// AppendDelegation adds a new delegation to the list of delegations.
// This function acquires a lock to ensure thread safety while modifying the delegation list and index map in the State.
func (s *State) AppendDelegation(del *ptypes.Delegation) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.delegations = append(s.delegations, del)

	key := fmt.Sprintf("%d:%d", del.ValidatorIndex, del.DelegatorIndex)
	s.delValMap[key] = uint64(len(s.delegations) - 1)
}

// UpdateDelegationAtIndex updates the details of a delegation at the specified index.
// This function acquires a lock to ensure thread safety while modifying the delegation list in the State.
func (s *State) UpdateDelegationAtIndex(idx uint64, delegation *ptypes.Delegation) {
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

// SetWithdrawals updates the withdrawal requests list in the U2U chain state.
// This function acquires a lock to ensure thread safety while modifying withdrawal data in the State.
func (s *State) SetWithdrawals(withdrawals []*ptypes.Withdrawal) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.withdrawals = withdrawals
}

// AppendWithdrawal adds a new withdrawal request to the list.
// This function acquires a lock to ensure thread safety while modifying the withdrawal list in the State.
func (s *State) AppendWithdrawal(withdrawal *ptypes.Withdrawal) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.withdrawals = append(s.withdrawals, withdrawal)
}

// SetWithdrawalRewards sets the list of withdrawal rewards in the U2U chain's state.
// This function acquires a lock to ensure thread safety while updating withdrawal reward data.
// It takes a slice of WithdrawalReward pointers as input and updates the internal withdrawal rewards list accordingly.
// Make sure to use this function within a context that handles concurrent access appropriately.
func (s *State) SetWithdrawalRewards(withdrawalRewards []*ptypes.WithdrawalReward) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.withdrawalRewards = withdrawalRewards
}

// AppendWithdrawalReward appends a withdrawal reward to the list of withdrawal rewards in the U2U chain's state.
// This function acquires a lock to ensure thread safety while appending withdrawal reward data.
// It takes a WithdrawalReward pointer as input and adds it to the end of the internal withdrawal rewards list.
// Make sure to use this function within a context that handles concurrent access appropriately.
func (s *State) AppendWithdrawalReward(withdrawalReward *ptypes.WithdrawalReward) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.withdrawalRewards = append(s.withdrawalRewards, withdrawalReward)
}
