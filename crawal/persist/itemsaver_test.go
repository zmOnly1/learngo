package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"learngo2/crawal/model"
	"testing"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
		Name:       "Mandy",
		Marriage:   "未婚",
		Age:        "32岁",
		Height:     "158cm",
		Xinzuo:     "魔羯座(12.22-01.19)",
		WorkPlace:  "工作地:广州天河区",
		Income:     "月收入:2-5万",
		Occupation: "生物工程",
		Education:  "大学本科",
	}
	id, err := save(expected)
	if err != nil {
		panic(err)
	}
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().
		Index("dating_profile").
		Id(id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	//source := resp.Source
	//fmt.Println(string(source))

	var actual model.Profile
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
