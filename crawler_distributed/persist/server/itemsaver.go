package main

import (
	"crawler/crawler_distributed/config"
	"crawler/crawler_distributed/persist"
	"crawler/crawler_distributed/rpcsupport"
	"fmt"
	"github.com/olivere/elastic"
	"log"
)

func main() {
	log.Fatal(serveRPC(fmt.Sprintf(":%d",
		config.ItemSaverPort),
		config.ElasticIndex))
}

func serveRPC(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})

}
