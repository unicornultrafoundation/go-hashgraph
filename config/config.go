package config

type U2UChainConfig struct {
	MinStakeAmount uint64

	EffectiveStakedBalanceIncrement uint64

	// Reward and penalty quotients constants.
	BaseRewardFactor uint64
}
