package entity

type Strategy struct {
	VaultAddress         string
	StrategyAddress      string
	DeFiLamaPool         string
	Chain                string
	Apy                  float64
	EstimatedTotalAssets float64
	WantTokenDecimals    int32
}
