package entity

type DeFiLamaResponse struct {
	Status string `json:"status"`
	Data   []struct {
		Chain        string        `json:"chain"`
		Project      string        `json:"project"`
		Symbol       string        `json:"symbol"`
		TvlUsd       int64         `json:"tvlUsd"`
		ApyBase      int           `json:"apyBase"`
		ApyReward    interface{}   `json:"apyReward"`
		Apy          float64       `json:"apy"`
		RewardTokens []interface{} `json:"rewardTokens"`
		Pool         string        `json:"pool"`
		ApyPct1D     int           `json:"apyPct1D"`
		ApyPct7D     int           `json:"apyPct7D"`
		ApyPct30D    interface{}   `json:"apyPct30D"`
		Stablecoin   bool          `json:"stablecoin"`
		IlRisk       string        `json:"ilRisk"`
		Exposure     string        `json:"exposure"`
		Predictions  struct {
			PredictedClass       interface{} `json:"predictedClass"`
			PredictedProbability interface{} `json:"predictedProbability"`
			BinnedConfidence     interface{} `json:"binnedConfidence"`
		} `json:"predictions"`
		PoolMeta         interface{} `json:"poolMeta"`
		Mu               int         `json:"mu"`
		Sigma            int         `json:"sigma"`
		Count            int         `json:"count"`
		Outlier          bool        `json:"outlier"`
		UnderlyingTokens []string    `json:"underlyingTokens"`
		Il7D             interface{} `json:"il7d"`
		ApyBase7D        int         `json:"apyBase7d"`
		ApyMean30D       int         `json:"apyMean30d"`
		VolumeUsd1D      interface{} `json:"volumeUsd1d"`
		VolumeUsd7D      interface{} `json:"volumeUsd7d"`
		ApyBaseInception interface{} `json:"apyBaseInception"`
	} `json:"data"`
}
