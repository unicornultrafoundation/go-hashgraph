package config

type U2UChainConfig struct {
	MinStakeAmount uint64

	BalanceIncrement uint64

	// Reward and penalty quotients constants.
	BaseRewardPerSecond uint64
	CommissionQuotient  uint64
}
