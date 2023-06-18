package entity

type ContractsSettings struct {
	Vaults []struct {
		Name              string `yaml:"name"`
		Chain             string `yaml:"chain"`
		Address           string `yaml:"address"`
		Decimals          int32  `yaml:"decimals"`
		WantTokenName     string `yaml:"want-token-name"`
		WantTokenAddress  string `yaml:"want-token-address"`
		WantTokenDecimals int32  `yaml:"want-token-decimals"`
		Strategies        []struct {
			Name    string `yaml:"name"`
			Address string `yaml:"address"`
			Pool    string `yaml:"pool"`
		} `yaml:"strategies"`
	} `yaml:"vaults"`
}
