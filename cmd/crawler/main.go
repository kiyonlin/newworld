package main

import (
	"github.com/kiyonlin/newworld/crawler/engine"
	"github.com/kiyonlin/newworld/crawler/persist"
	"github.com/kiyonlin/newworld/crawler/scheduler"
	"github.com/kiyonlin/newworld/crawler/zhenai/parser"
)

func main() {
	//http://www.zhenai.com/zhenghun
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{URL: "http://www.zhenai.com/zhenghun", ParserFunc: parser.ParseCityList})
}
