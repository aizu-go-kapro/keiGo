package models

import (
	"fmt"
	"github.com/ikawaha/kagome/tokenizer"
	"strings"
)

type Kagome struct{}

func (ur *Kagome) MorphologicalAnalysis(request KeigoRequest) KeigoResponse {
	var response KeigoResponse
	t := tokenizer.New()
	tokens := t.Tokenize(request.Body) // t.Analyze("寿司が食べたい。", tokenizer.Normal)
	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			// BOS: Begin Of Sentence, EOS: End Of Sentence.
			fmt.Printf("%s\n", token.Surface)
			continue
		}
		features := strings.Join(token.Features(), ",")
		fmt.Printf("%s\t%v\n", token.Surface, features)
		response.ConvertedBody += token.Surface + " "
	}
	return response
}
