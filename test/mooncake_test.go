package test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mooncake/executor"
	"testing"
)

func TestBasicRules(t *testing.T) {
	content, err := ioutil.ReadFile("../test/rules_v1.0.mck")

	if err != nil {
		panic(err)
	}

	rules := string(content)
	jsonFile := "../sample/sample_item.json"

	// sample context struct
	ctx := struct {
		item1 string
		item2 string
		threshold int
	}{
		"comingFromContextA",
		"comingFromContextB",
		10,
	}

	result := executor.Execute(rules, jsonFile, ctx)

	prettyResult, err := json.MarshalIndent(result, "", "  ");

	log.Printf("\n\n====Result===== \n\n %s\n\n", prettyResult)
}