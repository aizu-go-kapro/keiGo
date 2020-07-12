package converter

import (
	"encoding/json"
	"io/ioutil"
)

type (
	Utils struct {
		Path string // json ファイルのパス
	}

	ConversionRule struct {
		Original  string `json:"original"`
		Converted string `json:"converted"`
	}
)

func (u *Utils) JsonDecoder(filename string) ([]ConversionRule, error) {
	// File openは実行時パスから見た相対パスになる
	bytes, err := ioutil.ReadFile("./utils/" + filename)
	if err != nil {
		return nil, err
	}

	// JSONデコード
	var conversionRules []ConversionRule
	if err := json.Unmarshal(bytes, &conversionRules); err != nil {
		return nil, err
	}

	return conversionRules, nil
}

func (u *Utils) FindConvertedFromConversionRule(conversionRules []ConversionRule, original string) string {
	for _, rule := range conversionRules {
		if rule.Original == original {
			return rule.Converted
		}
	}
	return original
}
