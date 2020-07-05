package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type (
	Utils struct{}

	ConversionRule struct {
		Original  string `json:"original"`
		Converted string `json:"converted"`
	}
)

func (u *Utils) JsonDecoder(filename string) []ConversionRule {
	// File openは実行時パスから見た相対パスになる
	bytes, err := ioutil.ReadFile("./utils/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	// JSONデコード
	var conversionRules []ConversionRule
	if err := json.Unmarshal(bytes, &conversionRules); err != nil {
		log.Fatal(err)
	}

	return conversionRules
}

func (u *Utils) FindConvertedFromConversionRule(conversionRules []ConversionRule, original string) string {
	for _, rule := range conversionRules {
		if rule.Original == original {
			return rule.Converted
		}
	}
	return original
}
