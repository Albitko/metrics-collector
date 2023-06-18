package price_collector

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/Albitko/metrics-collector/internal/entity"
)

const coinGeckoAPI = "https://api.coingecko.com/api/v3/"

type coinGeckoResp map[string]map[string]float64
type priceCollector struct {
	tokensToCollect []string
	httpClient      *resty.Client
}

func (pc *priceCollector) Collect() {
	fmt.Println("start collecting prices at: ", time.Now().String())

	var wg sync.WaitGroup
	for _, token := range pc.tokensToCollect {
		wg.Add(1)
		go getPrice(token, &wg)
	}

	wg.Wait()
	// TODO recieve data from chanel and write to DB
	fmt.Println("finish collecting prices at: ", time.Now().String())

}

func getPrice(tokenID string, wg *sync.WaitGroup) {
	client := resty.New()
	resp, _ := client.R().SetResult(coinGeckoResp{}).SetContext(context.Background()).Get(
		coinGeckoAPI + "simple/price?ids=" + tokenID + "&vs_currencies=usd",
	)
	result := *resp.Result().(*coinGeckoResp)
	fmt.Println(result[tokenID]["usd"])
	wg.Done()
}

func New(contractsCfg entity.ContractsSettings, client *resty.Client) *priceCollector {
	var tokens []string
	for _, vault := range contractsCfg.Vaults {
		tokens = append(tokens, vault.WantTokenName)
	}
	return &priceCollector{
		tokensToCollect: tokens,
		httpClient:      client,
	}
}
