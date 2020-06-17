package models

import (
	"fmt"
	"github.com/ikawaha/kagome/tokenizer"
	"strings"
)

type (
	KeigoRequest struct {
		Body string `json:"body" binding:"required,omitempty"`
	}
	KeigoResponse struct {
		ConvertedBody string `json:"converted_body" binding:"required,omitempty"`
	}
	Keigo struct{}
)

func (k *Keigo) Convert(tokens []tokenizer.Token) string {
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
