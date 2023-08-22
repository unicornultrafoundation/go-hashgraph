package precompute

import (
	"github.com/unicornultrafoundation/go-hashgraph/config"
	"github.com/unicornultrafoundation/go-hashgraph/consensus/helpers"
	ptypes "github.com/unicornultrafoundation/go-hashgraph/proto/u2u/types"
	"github.com/unicornultrafoundation/go-hashgraph/state"
)

// New generates a list of precomputed Validator objects and computes the balance for reward distribution.

// This function takes the U2U network state as input and returns a list of precomputed Validator objects
// along with the balance information used for reward distribution.
func New(s *state.State) ([]*Validator, *Balance, error) {
	cfg := config.U2UConfig()
	pValidators := make([]*Validator, s.NumValidators())
	lastBlock := s.LatestBlock()
	epoch := s.Epoch()
	bal := &Balance{}
	epochDuration := s.Time() - lastBlock.Time
	bal.Reward = epochDuration * cfg.BaseRewardPerSecond
	if err := s.ReadFromEveryValidator(func(idx int, val *ptypes.Validator) error {
		uptime := val.Uptime + lastBlock.Time - val.LastOnlineTime
		pVal := &Validator{
			IsSlashed:     val.Slashed,
			MissedBlocks:  uint64(lastBlock.Id - val.LastBlockId),
			MissedPeriod:  lastBlock.Time - val.LastOnlineTime,
			IsActive:      helpers.IsActiveValidator(val, epoch),
			Uptime:        uptime,
			StakedBalance: val.StakedBalance,
		}

		if pVal.IsActive {
			bal.Active += val.StakedBalance
			pVal.BaseRewardWeight = val.StakedBalance * uptime / epochDuration
			bal.BaseRewardWeight += pVal.BaseRewardWeight

			txFees := val.TxFees - val.PrevTxFees
			pVal.TxRewardWeight = txFees * pVal.Uptime / epochDuration
			bal.TxRewardWeght += pVal.TxRewardWeight
			bal.EpochFee += txFees
		}

		pValidators[idx] = pVal
		return nil
	}); err != nil {
		return nil, nil, err
	}

	return pValidators, bal, nil
}
