package persist

import (
	"context"
	"encoding/json"
	"testing"

	"gopkg.in/olivere/elastic.v5"

	"github.com/kiyonlin/newworld/crawler/engine"
	"github.com/kiyonlin/newworld/crawler/model"
)

func TestSaver(t *testing.T) {
	expected := engine.Item{
		URL:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		ID:   "108906739",
		Payload: model.Profile{
			Name:          "HiSiri",
			Gender:        "女",
			Age:           28,
			Height:        163,
			Weight:        100,
			Income:        "3001-5000元",
			Marriage:      "未婚",
			Hukou:         "内蒙古赤峰",
			Constellation: "金牛座",
			House:         "自住",
			Car:           "未购车",
		},
	}

	// TODO: Try to start up elastic search
	// here using docker go client
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)

	if err != nil {
		panic(err)
	}

	err = save(client, "dating_profile", expected)
	if err != nil {
		panic(err)
	}

	// fetch expected
	resp, err := client.Get().Index("dating_profile").Type(expected.Type).Id(expected.ID).Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%+v, %v", resp, resp.Source)
	var actual engine.Item
	err = json.Unmarshal([]byte(*resp.Source), &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, err := model.FromJSONObj(actual.Payload)
	actual.Payload = actualProfile

	// verify result
	if actual != expected {
		t.Errorf("Got %v; expected %v", actual, expected)
	}
}
