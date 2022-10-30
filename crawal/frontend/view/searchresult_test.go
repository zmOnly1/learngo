package view

import (
	"learngo2/crawal/engine"
	"learngo2/crawal/frontend/model"
	common "learngo2/crawal/model"
	"os"
	"testing"
)

func TestSearchResultView_render(t *testing.T) {
	view := CreateSearchResultView("template.html")
	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url: "http://album.zhenai.com/u/1736961045",
		Id:  "1736961045",
		Payload: common.Profile{
			Name:       "Mandy",
			Marriage:   "未婚",
			Age:        "32岁",
			Height:     "158cm",
			Xinzuo:     "魔羯座(12.22-01.19)",
			WorkPlace:  "工作地:广州天河区",
			Income:     "月收入:2-5万",
			Occupation: "生物工程",
			Education:  "大学本科",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	out, err := os.Create("template.test.html")
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
