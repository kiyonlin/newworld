package parser

import (
	"regexp"
	"strconv"

	"github.com/kiyonlin/newworld/crawler/engine"
	"github.com/kiyonlin/newworld/crawler/model"
)

var nameRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)</span></td>`)

var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var hukouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(` <td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var constellationRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)

var guessRe = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)

var idURLRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

// ParseProfile parse profile info from the contests
func ParseProfile(contents []byte, name string, url string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Age = extractInt(contents, ageRe)
	profile.Height = extractInt(contents, heightRe)
	profile.Weight = extractInt(contents, weightRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Hukou = extractString(contents, hukouRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)
	profile.Constellation = extractString(contents, constellationRe)
	id := extractString([]byte(url), idURLRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				URL:     url,
				Type:    "zhenai",
				ID:      id,
				Payload: profile,
			},
		},
	}

	matches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		url := string(m[1])
		username := string(m[2])
		result.Requests = append(result.Requests, engine.Request{
			URL:        url,
			ParserFunc: ProfileParser(username),
		})
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}

func extractInt(contents []byte, re *regexp.Regexp) int {
	i, err := strconv.Atoi(extractString(contents, re))
	if err == nil {
		return i
	}

	return 0
}

// ProfileParser parse profile data
func ProfileParser(name string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProfile(c, name, url)
	}
}
