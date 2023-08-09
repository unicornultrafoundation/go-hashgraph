package state

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unicornultrafoundation/go-hashgraph/native/idx"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

// State defines a struct that encapsulates various utilities for managing the U2U Chain state.
type State struct {
	epoch              idx.Epoch                 // Current epoch information.
	validators         []*types.Validator        // List of active validators.
	stakedBalances     []uint64                  // Staked balances of delegators.
	accumulatedRewards []uint64                  // Accumulated rewards for validators.
	slashings          []uint64                  // Slashings information of validators.
	withdrawals        []*types.Withdrawal       // Withdrawal requests from delegators.
	withdrawalRewards  []*types.WithdrawalReward // Withdrawal rewards associated with delegator withdrawals.
	delegations        []*types.Delegation       // Delegation details.
	delegators         []*types.Delegator        // List of delegators.

	lock      sync.RWMutex                 // Mutex for thread-safe access.
	valMap    map[common.Address]uint64    // Map to store validator addresses and associated data.
	delValMap map[uint64]map[uint64]uint64 // Nested map to store delegation and validator relationship.
	delMap    map[common.Address]uint64    // Map to store delegator addresses and associated data.
}
