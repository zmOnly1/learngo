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

const ageRe = `<div data-[\w-]+="" class="m-btn purple">([^<]+)</div>`
const data = `<div data-v-8b1eac0c="" class="purple-btns"><div data-v-8b1eac0c="" class="m-btn purple">未婚</div><div data-v-8b1eac0c="" class="m-btn purple">32岁</div><div data-v-8b1eac0c="" class="m-btn purple">魔羯座(12.22-01.19)</div><div data-v-8b1eac0c="" class="m-btn purple">158cm</div><div data-v-8b1eac0c="" class="m-btn purple">工作地:广州天河区</div><div data-v-8b1eac0c="" class="m-btn purple">月收入:2-5万</div><div data-v-8b1eac0c="" class="m-btn purple">生物工程</div><div data-v-8b1eac0c="" class="m-btn purple">大学本科</div></div>`

func ParseProfile(contents []byte, name string) engine.ParseResult {
	node := selectNode(contents)
	re := regexp.MustCompile(ageRe)
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
