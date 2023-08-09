package state

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unicornultrafoundation/go-hashgraph/native/idx"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

// State defines a struct that encapsulates various utilities for managing the U2U Chain state.
type State struct {
	epoch          idx.Epoch           // Current epoch information.
	validators     []*types.Validator  // List of active validators.
	stakedBalances []uint64            // Staked balances of delegators.
	slashing       []uint64            // Slashing information of validators.
	withdrawals    []*types.Withdrawal // Withdrawal requests from delegators.

	lock   sync.RWMutex              // Mutex for thread-safe access.
	valMap map[common.Address]uint64 // Map to store validator addresses and associated data.
}
