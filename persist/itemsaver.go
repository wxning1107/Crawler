package persist

import (
	"context"
	"crawler/engine"
	"errors"
	"github.com/olivere/elastic"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	// Must turn off sniff in docker
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver: got item "+
				"%d: %v", itemCount, item)
			itemCount++
			err := save(client, item, index)
			if err != nil {
				log.Printf("Item saver: error saving item %v: %v", item, err)
			}
		}
	}()

	return out, nil
}

func save(client *elastic.Client, item engine.Item, index string) error {

	if item.Type == "" {
		return errors.New("must supply type")
	}
	// index save data, index is database,type is table
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
