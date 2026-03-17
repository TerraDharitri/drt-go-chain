package testscommon

import (
	"github.com/TerraDharitri/drt-go-chain/config"
)

// GetEconomicsConfig returns the common configuration used for testing
func GetEconomicsConfig() config.EconomicsConfig {
	return config.EconomicsConfig{
		GlobalSettings: config.GlobalSettings{
			GenesisTotalSupply: "2000000000000000000000",
			MinimumInflation:   0,
			YearSettings: []*config.YearSetting{
				{
					Year:             0,
					MaximumInflation: 0.01,
				},
			},
		},
		RewardsSettings: config.RewardsSettings{
			RewardsConfigByEpoch: []config.EpochRewardSettings{
				{
					LeaderPercentage:                 0.1,
					ProtocolSustainabilityPercentage: 0.1,
					DeveloperPercentage:              0.1,
					ProtocolSustainabilityAddress:    "drt1j25xk97yf820rgdp3mj5scavhjkn6tjyn0t63pmv5qyjj7wxlcfqa9aqd2",
					TopUpGradientPoint:               "100000",
					EcosystemGrowthPercentage:        0.0,
					EcosystemGrowthAddress:           "drt1932eft30w753xyvme8d49qejgkjc09n5e49w4mwdjtm0neld797spn6u9l",
					GrowthDividendPercentage:         0.0,
					GrowthDividendAddress:            "drt1932eft30w753xyvme8d49qejgkjc09n5e49w4mwdjtm0neld797spn6u9l",
				},
			},
		},
		FeeSettings: config.FeeSettings{
			GasLimitSettings: []config.GasLimitSetting{
				{
					EnableEpoch:                 0,
					MaxGasLimitPerBlock:         "1500000000",
					MaxGasLimitPerMiniBlock:     "1500000000",
					MaxGasLimitPerMetaBlock:     "1500000000",
					MaxGasLimitPerMetaMiniBlock: "1500000000",
					MaxGasLimitPerTx:            "600000000",
					MinGasLimit:                 "50000",
					ExtraGasLimitGuardedTx:      "50000",
					MaxGasHigherFactorAccepted:  "10",
				},
			},
			MinGasPrice:            "1000000000",
			GasPerDataByte:         "1500",
			GasPriceModifier:       0.01,
			MaxGasPriceSetGuardian: "2000000000",
		},
	}
}
