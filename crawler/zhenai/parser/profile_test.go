package parser

import (
	"io/ioutil"
	"testing"

	"github.com/kiyonlin/newworld/crawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "HiSiri")
	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 element; bus has %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
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
	}

	if profile != expected {
		t.Errorf("expected %v; bus was %v", expected, profile)
	}
}
