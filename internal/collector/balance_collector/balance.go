package balance_collector

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-resty/resty/v2"

	"github.com/Albitko/metrics-collector/internal/entity"
)

type balanceCollector struct {
	tokensToCollect []string
	httpClient      *resty.Client
}
type coinGeckoResp map[string]map[string]float64

const coinGeckoAPI = "https://api.coingecko.com/api/v3/"

func getPrice(tokenID string, wg *sync.WaitGroup, c *resty.Client) {
	resp, _ := c.R().SetResult(coinGeckoResp{}).SetContext(context.Background()).Get(
		coinGeckoAPI + "simple/price?ids=" + tokenID + "&vs_currencies=usd",
	)
	result := *resp.Result().(*coinGeckoResp)
	fmt.Println(result[tokenID]["usd"])
	wg.Done()
}

func (bc *balanceCollector) Collect(job gocron.Job) {
	fmt.Println("start collecting balances at: ", time.Now().String())

	var wg sync.WaitGroup
	fmt.Println("TOKENSLENGTH ", len(bc.tokensToCollect))
	for _, token := range bc.tokensToCollect {
		wg.Add(1)
		go getPrice(token, &wg, bc.httpClient)
	}

	wg.Wait()
	// TODO recieve data from chanel and write to DB
	fmt.Println("Всё готово fast")

}

func New(contractsCfg entity.ContractsSettings, client *resty.Client) *balanceCollector {
	var tokens []string
	for _, vault := range contractsCfg.Vaults {
		tokens = append(tokens, vault.WantTokenName)
	}
	return &balanceCollector{
		tokensToCollect: tokens,
		httpClient:      client,
	}
}

//fmt.Println("Application run")
//task_fast := func(in string, job gocron.Job) {
//	fmt.Println(in)
//	fmt.Println("Current date and time is: ", time.Now().String())
//	var wg sync.WaitGroup
//	const n = 5
//
//	for i := 0; i < n; i++ {
//		wg.Add(1) // инкрементируем счётчик, сколько горутин нужно подождать
//
//		go func(i int) {
//			time.Sleep(200 * time.Millisecond)
//			fmt.Printf("fast hi %d\n", i)
//			// уменьшаем счётчик, когда горутина завершает работу
//			wg.Done()
//		}(i)
//	}
//
//	wg.Wait() // ждём все горутины
//	fmt.Println("Всё готово fast")
//}
//task_slow := func(in string, job gocron.Job) {
//	fmt.Println(in)
//	fmt.Println("Current date and time is: ", time.Now().String())
//	time.Sleep(3 * time.Second)
//	var wg sync.WaitGroup
//	const n = 5
//
//	for i := 0; i < n; i++ {
//		wg.Add(1) // инкрементируем счётчик, сколько горутин нужно подождать
//
//		go func(i int) {
//			time.Sleep(100 * time.Millisecond)
//			fmt.Printf("slow hi %d\n", i)
//			// уменьшаем счётчик, когда горутина завершает работу
//			wg.Done()
//		}(i)
//	}
//
//	wg.Wait() // ждём все горутины
//	fmt.Println("Всё готово slow")
//}
//
//var err error
