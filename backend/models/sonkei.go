package models

import (
	"fmt"
	"strings"

	"github.com/aizu-go-kapro/keiGo/backend/utils"
	"github.com/ikawaha/kagome/tokenizer"
)

type Sonkei struct{}

func (s *Sonkei) Convert(body string) string {
	kagome := Kagome{}
	tokens := kagome.MorphologicalAnalysis(body)

	utils := utils.Utils{}
	conversionRules := utils.JsonDecoder("sonkei.json")

	var convertedBody = ""

	endOfSentenceTokenIndex := len(tokens) - 1
	for {
		if tokens[endOfSentenceTokenIndex].Features()[0] == "記号" {
			endOfSentenceTokenIndex--
		} else {
			break
		}
	}

	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			// BOS: Begin Of Sentence, EOS: End Of Sentence.
			fmt.Printf("%s\n", token.Surface)
			continue
		}
		features := strings.Join(token.Features(), ",")
		fmt.Printf("%s\t%v\n", token.Surface, features)

		// そのtokenの基本形が変換ルールに存在するか判定し、変化後の文字列を取得する
		token.Surface = utils.FindConvertedFromConversionRule(conversionRules, token.Features()[6])

		convertedBody += token.Surface
	}

	return convertedBody
}
