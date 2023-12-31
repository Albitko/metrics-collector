package app

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-resty/resty/v2"

	"github.com/Albitko/metrics-collector/internal/collector/balance_collector"
	"github.com/Albitko/metrics-collector/internal/collector/price_collector"
	"github.com/Albitko/metrics-collector/internal/collector/strategy_collector"
	"github.com/Albitko/metrics-collector/internal/collector/vault_collector"
	"github.com/Albitko/metrics-collector/internal/entity"
	"github.com/Albitko/metrics-collector/internal/repo"
	"github.com/Albitko/metrics-collector/internal/utils"
)

type collector interface {
	Collect()
}

func mustScheduleJob(s *gocron.Scheduler, job interface{}) {
	var err error
	// at every 5th minute.
	_, err = s.Cron("*/5 * * * *").Do(job)
	if err != nil {
		log.Fatalln("error scheduling job", err)
	}
}

func Run(contractsCfg entity.ContractsSettings) {
	var price, strategy, vault, balance collector

	httpClient := resty.New()
	rpc := utils.NewRPC()
	defer rpc.Close()
	db := repo.New(nil, "postgres://postgres@localhost:5432/postgres")
	defer db.Close()

	price = price_collector.New(contractsCfg, httpClient)
	strategy = strategy_collector.New(contractsCfg, httpClient, rpc)
	vault = vault_collector.New(contractsCfg, rpc)
	balance = balance_collector.New(contractsCfg, httpClient)

	s := gocron.NewScheduler(time.UTC)
	s.SingletonModeAll()
	mustScheduleJob(s, price.Collect)
	mustScheduleJob(s, strategy.Collect)
	mustScheduleJob(s, vault.Collect)
	mustScheduleJob(s, balance.Collect)

	s.StartBlocking()

}
