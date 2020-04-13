package client

import (
	"crawler/crawler_distributed/config"
	"crawler/crawler_distributed/rpcsupport"
	"crawler/crawler_distributed/worker"
	"crawler/engine"
	"fmt"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort))
	if err != nil {
		return nil, err
	}

	return func(req engine.Request) (engine.ParseResult, error) {

		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}

		return worker.DeserializeResult(sResult), nil
	}, nil
}
