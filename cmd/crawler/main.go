package main

import (
	"github.com/kiyonlin/newworld/crawler/engine"
	"github.com/kiyonlin/newworld/crawler/scheduler"
	"github.com/kiyonlin/newworld/crawler/zhenai/parser"
)

func main() {
	//http://www.zhenai.com/zhenghun
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{URL: "http://www.zhenai.com/zhenghun", ParserFunc: parser.ParseCityList})
}
