package main

import (
	"crawler/crawler_distributed/rpcsupport"
	"crawler/engine"
	"crawler/model"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1107"

	// start itemSaverServer
	go serveRPC(host, "test1")

	time.Sleep(time.Second)

	// start itemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Call save
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Id:   "108906739",
		Type: "zhenai",
		Payload: model.Profile{
			Name:       "安静的雪",
			Gender:     "女",
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Marriage:   "离异",
			Education:  "大学本科",
			Occupation: "人事/行政",
			Hokou:      "山东菏泽",
			Xinzuo:     "牧羊座",
			House:      "已购房",
			Car:        "未购车",
		},
	}

	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %v; err: %v", result, err)
	}
}
