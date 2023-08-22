package precompute

type Validator struct {
	IsSlashed        bool
	IsActive         bool
	StakedBalance    uint64
	Uptime           uint64
	MissedBlocks     uint64
	MissedPeriod     uint64
	BaseRewardWeight uint64
	TxRewardWeight   uint64
}

type Balance struct {
	Active           uint64
	BaseRewardWeight uint64
	TxRewardWeght    uint64
	Reward           uint64
	EpochFee         uint64
}
