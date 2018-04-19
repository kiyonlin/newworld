package parser

import (
	"regexp"

	"github.com/kiyonlin/newworld/crawler/engine"
)

var (
	cityReg    = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityURLReg = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)">`)
)

// ParseCity parse the city into parse result
func ParseCity(contents []byte) engine.ParseResult {
	match := cityReg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range match {
		url := string(m[1])
		userName := string(m[2])
		// result.Items = append(result.Items, "User: "+userName)
		result.Requests = append(result.Requests, engine.Request{
			URL: url,
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, userName, url)
			},
		})
	}

	matches := cityURLReg.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			URL:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
