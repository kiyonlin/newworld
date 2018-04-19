package parser

import (
	"io/ioutil"
	"testing"

	"github.com/kiyonlin/newworld/crawler/engine"
	"github.com/kiyonlin/newworld/crawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "HiSiri", "http://album.zhenai.com/u/108906739")
	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 element; bus has %v", result.Items)
	}

	actual := result.Items[0]

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
		}}

	if actual != expected {
		t.Errorf("expected %v; bus was %v", expected, actual)
	}
}
