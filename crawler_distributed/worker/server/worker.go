package main

import (
	"crawler/crawler_distributed/config"
	"crawler/crawler_distributed/rpcsupport"
	"crawler/crawler_distributed/worker"
	"fmt"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort), worker.CrawlService{}))
}
