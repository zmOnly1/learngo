package parser

import (
	"os"
	"testing"
)

func TestPrintCityList(t *testing.T) {
	contents, err := os.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := PrintCityList(contents)
	const resultSize = 470

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d Items; but had %d", resultSize, len(result.Items))
	}

}
