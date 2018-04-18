package parser

import (
	"regexp"

	"github.com/kiyonlin/newworld/crawler/engine"
)

const citylistReg = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// ParseCityList parse city list from the contents
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(citylistReg)
	match := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range match {
		// result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			URL:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
