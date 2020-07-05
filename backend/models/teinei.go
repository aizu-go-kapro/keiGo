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

func (t *Teinei) Convert(body string) string {
	kagome := Kagome{}
	tokens := kagome.MorphologicalAnalysis(body)

	var convertedBody = ""

	endOfSentenceTokenIndex := len(tokens) - 1
	for {
		if tokens[endOfSentenceTokenIndex].Features()[0] == "記号" {
			endOfSentenceTokenIndex--
		} else {
			break
		}
	}
	if teineiToken := tokens[endOfSentenceTokenIndex]; teineiToken.Surface != "です" && teineiToken.Surface != "ます" {
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
				if token.Surface == "だ" && token.Features()[0] == "助動詞" {
					token.Surface = "です"
				} else if token.Features()[0] == "助動詞" {
					// 動詞の連用形＋助動詞＋です、ます
					word := []rune(token.Surface)
					//過去形の "った(ex: 帰った)" と "た(ex: 伝えた)"の時
					if token.Surface == "た" && tokens[i-1].Features()[0] == "動詞" {
						doushi := []rune(tokens[i-1].Surface)
						tail := doushi[len(doushi)-1]
						if string(tail) == "っ" {
							basicFormDoushi := []rune(tokens[i-1].Features()[6])
							doushiHead, doushiTail := basicFormDoushi[0:len(basicFormDoushi)-1], basicFormDoushi[len(basicFormDoushi)-1]
							runeConvertedBody := []rune(convertedBody)
							convertedBody = string(runeConvertedBody[0:len(runeConvertedBody)-2]) + string(doushiHead) + utoi[string(doushiTail)] + "まし"
						} else if tokens[i-1].Features()[0] == "動詞" {
							token.Surface = "ました"
						}
					} else if !(tokens[i-1].Surface == "でし" || tokens[i-1].Surface == "まし") {
						head, tail := word[0:len(word)-1], word[len(word)-1]
						_, isConverted := utoi[string(tail)]
						if isConverted == true {
							token.Surface = string(head) + "ます"
						} else {
							token.Surface = token.Surface + "です"
						}
					}
				}
				if token.Features()[0] == "名詞" || token.Features()[0] == "形容詞" {
					token.Surface = token.Surface + "です"
				}
			}

			convertedBody += token.Surface
		}
	}

	return convertedBody
}
