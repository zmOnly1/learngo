package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"learngo2/crawal/engine"
	"learngo2/crawal/model"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url: "http://album.zhenai.com/u/1736961045",
		Id:  "1736961045",
		Payload: model.Profile{
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
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"

	err = Save(client, index, expected)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index(index).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	//source := resp.Source
	//fmt.Println(string(source))

	var actual engine.Item
	json.Unmarshal(resp.Source, &actual)
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
