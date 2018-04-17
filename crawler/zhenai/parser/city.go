package parser

import (
	"regexp"

	"github.com/kiyonlin/newworld/crawler/engine"
)

const cityReg = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

// ParseCity parse the city into parse result
func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityReg)
	amtch := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range amtch {
		userName := string(m[2])
		result.Items = append(result.Items, "User: "+userName)
		result.Requests = append(result.Requests, engine.Request{
			URL: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, userName)
			},
		})
	}
	return result
}
