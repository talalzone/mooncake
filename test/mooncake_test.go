package test

import (
	"../executor"
	"io/ioutil"
	"log"
	"testing"
)

func TestBasicRules(t *testing.T) {
	content, err := ioutil.ReadFile("../test/rules_v1.0.mck")

	if err != nil {
		panic(err)
	}

	rules := string(content)
	jsonFile := "../sample/sample_items.json"

	// sample context struct
	ctx := struct {
		item1 string
		item2 string
		item3 int
	}{
		"comingFromContextA",
		"comingFromContextB",
		100000,
	}

	result := executor.Execute(rules, jsonFile, ctx)

	log.Printf("Result - %+v\n", result)
}
