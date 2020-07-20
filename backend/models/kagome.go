package models

import (
	"github.com/ikawaha/kagome/tokenizer"
)

type Kagome struct{}

func (k *Kagome) MorphologicalAnalysis(body string) []tokenizer.Token {
	t := tokenizer.New()
	tokens := t.Tokenize(body) // t.Analyze("寿司が食べたい。", tokenizer.Normal)
	// BOSとEOSはリストに含まない
	tokens = tokens[1 : len(tokens)-1]
	return tokens
}
