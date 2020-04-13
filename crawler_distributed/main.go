package main

import (
	"crawler/crawler_distributed/config"
	itemSaver "crawler/crawler_distributed/persist/client"
	worker "crawler/crawler_distributed/worker/client"
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	"fmt"
)

func main() {
	itemChan, err := itemSaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "https://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})

}
