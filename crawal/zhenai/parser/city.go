package parser

import (
	"learngo2/crawal/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/\d+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		userName := string(m[2])
		result.Items = append(result.Items, "User "+userName)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, userName)
			},
		})
	}
	return result
}