package entity

type Vault struct {
	Address           string
	Chain             string
	TotalSupply       float64
	PricePerShare     float64
	WantTokenDecimals int32
}
