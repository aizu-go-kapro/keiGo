package models

import (
	"fmt"
	"strings"

	"github.com/ikawaha/kagome/tokenizer"
)

type Teinei struct{}

var utoi = map[string]string{
	"う": "い",
	"く": "き",
	"す": "し",
	"つ": "ち",
	"ぬ": "に",
	"ふ": "ひ",
	"む": "み",
	"ゆ": "い",
	"る": "り",
}

func (t *Teinei) Convert(tokens []tokenizer.Token) string {
	var convertedBody = ""

	// BOSとEOSはリストに含まない
	tokens = tokens[1 : len(tokens)-1]
	endOfSentenceTokenIndex := len(tokens) - 1
	for {
		if tokens[endOfSentenceTokenIndex].Features()[0] == "記号" {
			endOfSentenceTokenIndex--
		} else {
			break
		}
	}

	for i, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			// BOS: Begin Of Sentence, EOS: End Of Sentence.
			fmt.Printf("%s\n", token.Surface)
			continue
		}
		features := strings.Join(token.Features(), ",")
		fmt.Printf("%s\t%v\n", token.Surface, features)

		if i == endOfSentenceTokenIndex {
			if token.Features()[0] == "動詞" {
				// 動詞の終止形 -> 動詞の連用形 + ます
				word := []rune(token.Surface)
				head, tail := word[0:len(word)-1], word[len(word)-1]
				token.Surface = string(head) + utoi[string(tail)] + "ます"
			}
			if token.Features()[0] == "助動詞" {
				// 動詞の連用形＋助動詞＋です、ます
				word := []rune(token.Surface)
				head, tail := word[0:len(word)-1], word[len(word)-1]
				_, isConverted := utoi[string(tail)]
				if isConverted == true {
					token.Surface = string(head) + "ます"
				} else {
					token.Surface = token.Surface + "です"
				}
			}
			if token.Surface == "だ" && token.Features()[0] == "助動詞" {
				token.Surface = "です"
			}
			if token.Features()[0] == "名詞" || token.Features()[0] == "形容詞" {
				token.Surface = token.Surface + "です"
			}
		}

		convertedBody += token.Surface
	}

	return convertedBody
}
