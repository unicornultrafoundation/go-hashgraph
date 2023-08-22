package state

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unicornultrafoundation/go-hashgraph/consensus/election"
	"github.com/unicornultrafoundation/go-hashgraph/hash"
	"github.com/unicornultrafoundation/go-hashgraph/native/idx"
	pState "github.com/unicornultrafoundation/go-hashgraph/proto/u2u/state"
	ptypes "github.com/unicornultrafoundation/go-hashgraph/proto/u2u/types"
	"github.com/unicornultrafoundation/go-hashgraph/state/stateutil"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

// State defines a struct that encapsulates various utilities for managing the U2U Chain state.
type State struct {
	epoch             idx.Epoch // Current epoch information.
	time              uint64
	prevTime          uint64
	latestBlock       *ptypes.Block
	lastDecidedFrame  idx.Frame
	validators        []*ptypes.Validator        // List of active validators.
	stakedBalances    []uint64                   // Staked balances of delegators.
	withdrawals       []*ptypes.Withdrawal       // Withdrawal requests from delegators.
	withdrawalRewards []*ptypes.WithdrawalReward // Withdrawal rewards associated with delegator withdrawals.
	delegations       []*ptypes.Delegation       // Delegation details.
	delegators        []*ptypes.Delegator        // List of delegators.
	roots             []*election.RootAndSlot
	confirmedEvents   []*types.ConfirmedEvent

	lock      sync.RWMutex              // Mutex for thread-safe access.
	valMap    map[common.Address]uint64 // Map to store validator addresses and associated data.
	delValMap map[string]uint64         // Nested map to store delegation and validator relationship.
	delMap    map[common.Address]uint64 // Map to store delegator addresses and associated data.
	ceMap     map[hash.Event]int
}

func FromProto(st *pState.State) *State {
	s := &State{
		validators:        st.Validators,
		delegators:        st.Delegators,
		stakedBalances:    st.StakedBalances,
		epoch:             idx.Epoch(st.Epoch),
		time:              st.Time,
		prevTime:          st.PrevTime,
		latestBlock:       st.LastBlock,
		withdrawals:       st.Withdrawals,
		withdrawalRewards: st.WithdrawalRewards,
		valMap:            stateutil.ValidatorIndexMap(st.Validators),
		delMap:            stateutil.DelegatorIndexMap(st.Delegators),
		delValMap:         make(map[string]uint64),
		lastDecidedFrame:  idx.Frame(st.LastDecidedFrame),
	}
	s.SetDelegations(st.Delegations)
	return s
}

func (s *State) ToProto() *pState.State {
	return &pState.State{
		Epoch:             uint64(s.epoch),
		Time:              s.time,
		PrevTime:          s.prevTime,
		LastBlock:         s.latestBlock,
		Validators:        s.validators,
		StakedBalances:    s.stakedBalances,
		Delegations:       s.delegations,
		Delegators:        s.delegators,
		Withdrawals:       s.withdrawals,
		WithdrawalRewards: s.withdrawalRewards,
		LastDecidedFrame:  uint64(s.lastDecidedFrame),
	}
}
