package parser

import (
	"learngo2/crawal/model"
	"os"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := os.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "Mandy")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}
	profile := result.Items[0].(model.Profile)
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

	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
