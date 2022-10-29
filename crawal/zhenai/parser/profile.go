package parser

import (
	"bytes"
	"github.com/ericchiang/css"
	"golang.org/x/net/html"
	"learngo2/crawal/engine"
	"learngo2/crawal/model"
	"regexp"
	"strings"
)

var (
	re      = regexp.MustCompile(`<div data-[\w-]+="" class="m-btn purple">([^<]+)</div>`)
	guessRe = regexp.MustCompile(`<a data-[\w-]+="" target="_self" class="user f-cl" href="//www.zhenai.com/n/login?channelId=905819&amp;fromurl=https%3A%2F%2Falbum.zhenai.com%2Fu%2F1160833431">`)
)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	node := selectNode(contents)
	matches := re.FindAllSubmatch(node, -1)

	profile := model.Profile{}
	profile.Name = name

	Marriage := matches[0][1]
	Age := matches[1][1]
	Xinzuo := matches[2][1]
	Height := matches[3][1]
	WorkPlace := matches[4][1]
	Income := matches[5][1]
	Occupation := matches[6][1]
	Education := matches[7][1]

	profile.Marriage = string(Marriage)
	profile.Age = string(Age)
	profile.Height = string(Height)
	profile.Xinzuo = string(Xinzuo)
	profile.WorkPlace = string(WorkPlace)
	profile.Income = string(Income)
	profile.Occupation = string(Occupation)
	profile.Education = string(Education)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func selectNode(contents []byte) []byte {
	sel, err := css.Parse(".m-content-box .purple-btns")
	if err != nil {
		panic(err)
	}
	node, err := html.Parse(strings.NewReader(string(contents)))
	if err != nil {
		panic(err)
	}

	for _, ele := range sel.Select(node) {
		b := &bytes.Buffer{}
		html.Render(b, ele)
		return b.Bytes()
	}
	return nil
}
