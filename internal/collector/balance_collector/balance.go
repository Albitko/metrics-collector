package balance_collector

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/Albitko/metrics-collector/internal/entity"
	"github.com/Albitko/metrics-collector/internal/utils"
)

type balanceCollector struct {
	vaults     []entity.Vault
	httpClient *resty.Client
}

const (
	chainBasePI  = "https://api.chainbase.online/v1/"
	etherscanAPI = "https://api.etherscan.io/api"
)

func getVaultBalance(wg *sync.WaitGroup, vault entity.Vault, c *resty.Client) {
	var balances []entity.UserBalance

	holdersResp := &entity.ChainBaseHoldersResponse{}
	c.R().
		SetResult(holdersResp).
		SetHeader("Accept", "application/json").
		SetHeader("x-api-key", "demo").
		SetContext(context.Background()).
		Get(
			chainBasePI + "token/holders?chain_id=1&contract_address=" + vault.Address + "&page=1&limit=100",
		)
	holders := holdersResp.Data
	etherScanResp := &entity.EtherScanResp{}

	for _, h := range holders {
		var userBalance entity.UserBalance

		c.R().
			SetResult(etherScanResp).
			SetContext(context.Background()).
			Get(
				etherscanAPI + "?module=account&action=tokenbalance&contractaddress=" +
					vault.Address + "&address=" + h + "&tag=latest&apikey=",
			)

		amountResp := new(big.Int)
		amountResp, _ = amountResp.SetString(etherScanResp.Result, 10)

		amount := utils.CastNumberWithDecimals(amountResp, vault.WantTokenDecimals)
		userBalance.UserAddress = h
		userBalance.VaultAddress = vault.Address
		userBalance.Chain = vault.Chain
		userBalance.Amount = amount
		balances = append(balances, userBalance)
	}

	for _, v := range balances {
		fmt.Println(v)
	}

	wg.Done()
}

func (bc *balanceCollector) Collect() {
	fmt.Println("start collecting balances at: ", time.Now().String())
	var wg sync.WaitGroup

	for _, v := range bc.vaults {
		wg.Add(1)
		go getVaultBalance(&wg, v, bc.httpClient)
	}
	wg.Wait()
	fmt.Println("finish collecting balances at: ", time.Now().String())

}

func New(contractsCfg entity.ContractsSettings, client *resty.Client) *balanceCollector {
	var vaults []entity.Vault
	for _, v := range contractsCfg.Vaults {
		var vault entity.Vault
		vault.Address = v.Address
		vault.Chain = v.Chain
		vault.WantTokenDecimals = v.WantTokenDecimals
		vaults = append(vaults, vault)
	}
	return &balanceCollector{
		vaults:     vaults,
		httpClient: client,
	}
}
