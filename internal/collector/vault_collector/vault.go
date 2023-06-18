package vault_collector

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"

	"github.com/Albitko/metrics-collector/internal/entity"
	"github.com/Albitko/metrics-collector/internal/utils"
)

const abi = `[
	{"stateMutability":"view","type":"function","name":"pricePerShare","inputs":[],"outputs":[{"name":"","type":"uint256"}]},
	{"stateMutability":"view","type":"function","name":"totalSupply","inputs":[],"outputs":[{"name":"","type":"uint256"}]}
]`

type rpc interface {
	Call(contractABI, address, funcName string) ([]interface{}, error)
}
type vaultCollector struct {
	vaults    []entity.Vault
	rpcClient rpc
}

func (vc *vaultCollector) Collect(job gocron.Job) {
	fmt.Println("start collecting vaults data at: ", time.Now().String())

	for i, v := range vc.vaults {
		supplyCallRes, _ := vc.rpcClient.Call(abi, v.Address, "totalSupply")
		fmt.Println(supplyCallRes)
		totalSupply := utils.CastNumberWithDecimals(supplyCallRes[0], v.WantTokenDecimals)
		priceCallRes, _ := vc.rpcClient.Call(abi, v.Address, "pricePerShare")
		fmt.Println(priceCallRes)

		pricePerShare := utils.CastNumberWithDecimals(priceCallRes[0], v.WantTokenDecimals)

		vc.vaults[i].TotalSupply = totalSupply
		vc.vaults[i].PricePerShare = pricePerShare
		fmt.Println(vc.vaults[i])
	}

	fmt.Println("finish collecting vaults data at: ", time.Now().String())
}

func New(contractsCfg entity.ContractsSettings, rpcClient rpc) *vaultCollector {
	var vaults []entity.Vault
	for _, v := range contractsCfg.Vaults {
		var vault entity.Vault
		vault.Chain = v.Chain
		vault.Address = v.Address
		vault.WantTokenDecimals = v.WantTokenDecimals
		vaults = append(vaults, vault)
	}

	return &vaultCollector{
		rpcClient: rpcClient,
		vaults:    vaults,
	}
}
