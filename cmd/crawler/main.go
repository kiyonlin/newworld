package main

import (
	"github.com/kiyonlin/newworld/crawler/engine"
	"github.com/kiyonlin/newworld/crawler/zhenai/parser"
)

func main() {
	//http://www.zhenai.com/zhenghun
	engine.Run(engine.Request{URL: "http://www.zhenai.com/zhenghun", ParserFunc: parser.ParseCityList})
}
