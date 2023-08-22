package epoch

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/unicornultrafoundation/go-hashgraph/consensus"
	"github.com/unicornultrafoundation/go-hashgraph/hash"
	pstate "github.com/unicornultrafoundation/go-hashgraph/proto/u2u/state"
	ptypes "github.com/unicornultrafoundation/go-hashgraph/proto/u2u/types"
	"github.com/unicornultrafoundation/go-hashgraph/state"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

func TestProcessFinalUpdates(t *testing.T) {
	var err error
	s := fakeState()

	s, err = consensus.ProcessCreateValidator(s, &types.CreateValidatorDto{
		Address: common.Address{},
		Amount:  100,
	})

	require.NoError(t, err)

	s, err = ProcessFinalUpdates(s)

	require.NoError(t, err)
	require.Equal(t, s.NumValidators(), 1)
	require.Equal(t, len(s.Delegations()), 1)
	require.Equal(t, len(s.Delegators()), 1)
}

func fakeState() *state.State {
	return state.FromProto(&pstate.State{
		Epoch:          1,
		Validators:     make([]*ptypes.Validator, 0),
		Delegations:    make([]*ptypes.Delegation, 0),
		StakedBalances: make([]uint64, 0),
		Delegators:     make([]*ptypes.Delegator, 0),
		Time:           uint64(time.Now().Unix()),
		PrevTime:       uint64(time.Now().Unix()),
		LastBlock: &ptypes.Block{
			Id:   1,
			Time: uint64(time.Now().Unix()),
			Hash: hash.Zero[:],
		},
	})
}
