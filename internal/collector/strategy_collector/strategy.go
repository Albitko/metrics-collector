package strategy_collector

import (
	"context"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/Albitko/metrics-collector/internal/entity"
	"github.com/Albitko/metrics-collector/internal/utils"
)

const (
	defiLamaAPI = "https://yields.llama.fi/pools"
	abi         = `[
	{"inputs":[],"name":"estimatedTotalAssets","outputs":[{"internalType":"uint256","name":"_wants","type":"uint256"}],"stateMutability":"view","type":"function"}
]`
)

type rpc interface {
	Call(contractABI, address, funcName string) ([]interface{}, error)
}

type strategyCollector struct {
	pools      map[string]*entity.Strategy
	httpClient *resty.Client
	rpcClient  rpc
}

func (sc *strategyCollector) Collect() {
	fmt.Println("start collecting strategies APY at: ", time.Now().String())

	poolsResponse := &entity.DeFiLamaResponse{}
	sc.httpClient.R().SetResult(poolsResponse).SetContext(context.Background()).Get(defiLamaAPI)

	for _, p := range poolsResponse.Data {
		if _, ok := sc.pools[p.Pool]; ok {
			sc.pools[p.Pool].Apy = p.Apy
		}
	}

	for k, p := range sc.pools {
		callRes, _ := sc.rpcClient.Call(abi, p.StrategyAddress, "estimatedTotalAssets")
		totalAssets := utils.CastNumberWithDecimals(callRes[0], p.WantTokenDecimals)
		sc.pools[k].EstimatedTotalAssets = totalAssets
		fmt.Println(sc.pools[k])
	}
	fmt.Println("finish collecting strategies APY lances at: ", time.Now().String())

}

func New(contractsCfg entity.ContractsSettings, client *resty.Client, rpcClient rpc) *strategyCollector {
	pools := make(map[string]*entity.Strategy)
	for _, v := range contractsCfg.Vaults {
		for _, vs := range v.Strategies {
			var s entity.Strategy
			s.VaultAddress = v.Address
			s.Chain = v.Chain
			s.DeFiLamaPool = vs.Pool
			s.StrategyAddress = vs.Address
			s.WantTokenDecimals = v.WantTokenDecimals
			pools[vs.Pool] = &s
		}
	}
	return &strategyCollector{
		httpClient: client,
		pools:      pools,
		rpcClient:  rpcClient,
	}
}
