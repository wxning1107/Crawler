package persist

import (
	"context"
	"crawler/engine"
	"crawler/model"
	"encoding/json"
	"github.com/olivere/elastic"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
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

	// TODO: try to start up elastic search
	// here using docker go client
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_profile"
	// Save expected item
	err = Save(client, expected, index)
	if err != nil {
		panic(err)
	}

	// Fetch Save item
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	// Verify result
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}

}
