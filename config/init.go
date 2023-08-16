package config

var config *U2UChainConfig

func U2UConfig() *U2UChainConfig {
	return config
}

func init() {
	config = mainnetConfig
}
