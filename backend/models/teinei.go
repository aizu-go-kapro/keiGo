package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ikawaha/kagome/tokenizer"
)

// Teinei struct
type Teinei struct{}

// Feature is type of token.Feature
type Feature struct {
	surface       string
	part          string
	class         []string
	variant       string
	conjugated    string
	base          string
	reading       string
	pronunciation string
}

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
	"ぶ": "び",
}

func convertRenyo(feat Feature) (string, error) {
	if feat.part != "動詞" {
		return "", errors.New("not verb")
	}

	variant := []rune(feat.variant)
	if string(variant[0:2]) == "五段" {
		verb := []rune(feat.base)
		l := verb[len(verb)-1:]
		last, ok := utoi[string(l)]
		if !ok {
			return "", errors.New("変換エラー")
		}
		return string(verb[0:len(verb)-1]) + last, nil
	}
	if string(variant[0:2]) == "一段" {
		verb := []rune(feat.base)
		return string(verb[0 : len(verb)-1]), nil
	}
	if string(variant[0:2]) == "カ変" {
		verb := []rune(feat.base)
		ki := verb[0]
		if string(ki) == "来" {
			return string(ki), nil
		}
		return "き", nil
	}
	if string(variant[0:2]) == "サ変" {
		return "し", nil
	}

	return feat.base, nil
}

// Convert is the function convert to teinei-go
func (t *Teinei) Convert(body string) (string, error) {
	kagome := Kagome{}
	tokens := kagome.MorphologicalAnalysis(body)

	endIndex := len(tokens) - 1
	for {
		if tokens[endIndex].Features()[0] == "記号" || tokens[endIndex].Features()[0] == "助詞" {
			endIndex--
		} else {
			break
		}
	}

	features := []Feature{}
	var surfaces []string
	for i, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			// BOS: Begin Of Sentence, EOS: End Of Sentence.
			fmt.Printf("%s\n", token.Surface)
			continue
		}
		fmt.Printf("%s\t%v\n", token.Surface, strings.Join(token.Features(), ","))

		f := token.Features()
		if len(f) < 9 {
			f = append(f, make([]string, 9-len(f))...)
		}
		fs := Feature{surface: token.Surface, part: f[0], class: f[1:4], variant: f[4], conjugated: f[5], base: f[6], reading: f[7], pronunciation: f[8]}

		// 例外処理　あろ+う == ある
		if fs.base == "う" && features[i-1].base == "ある" {
			surfaces[i-1] = "ある"
			continue
		}

		features = append(features, fs)
		surfaces = append(surfaces, token.Surface)
	}

	// すでに敬語の場合
	for _, fs := range features {
		if fs.base == "です" || fs.base == "ます" {
			return strings.Join(surfaces, ""), nil
		}
	}

	// 文末が動詞
	if features[endIndex].part == "動詞" {
		verb, err := convertRenyo(features[endIndex])
		if err != nil {
			return "", err
		}
		surfaces[endIndex] = string(verb) + "ます"
		return strings.Join(surfaces, ""), nil
	}

	if features[endIndex].part == "名詞" {
		surfaces[endIndex] = surfaces[endIndex] + "です"
		return strings.Join(surfaces, ""), nil
	}

	// 助動詞の「だ」が文中にある
	for i, f := range features {
		if f.base == "だ" && f.part == "助動詞" {
			if i != endIndex && surfaces[i] == "で" && features[i+1].base == "ある" {
				surfaces[i] = "でしょ"
				surfaces[i+1] = "う"
			} else if conjufated := []rune(f.conjugated); string(conjufated[0:2]) == "未然" {
				surfaces[i] = "でしょ"
			} else if string(conjufated[0:2]) == "連用" {
				surfaces[i] = "でし"
			} else {
				surfaces[i] = "です"
			}
			return strings.Join(surfaces, ""), nil
		}
	}

	// 文末が「ない」
	if features[endIndex].base == "ない" {
		if features[endIndex-1].part == "動詞" {
			verb, err := convertRenyo(features[endIndex-1])
			if err != nil {
				return "", err
			}
			surfaces[endIndex-1] = string(verb)
			surfaces[endIndex] = "ません"
			return strings.Join(surfaces, ""), nil
		}
		surfaces[endIndex] = surfaces[endIndex] + "です"
		return strings.Join(surfaces, ""), nil
	}

	// 文末が「た」で一つ前が動詞
	if features[endIndex].base == "た" && features[endIndex-1].part == "動詞" {
		verb, err := convertRenyo(features[endIndex-1])
		if err != nil {
			return "", err
		}
		surfaces[endIndex-1] = string(verb)
		surfaces[endIndex] = "ました"
		return strings.Join(surfaces, ""), nil
	}

	// 文末が「う」で一つ前が動詞
	if features[endIndex].base == "う" && features[endIndex-1].part == "動詞" {
		verb, err := convertRenyo(features[endIndex-1])
		if err != nil {
			return "", err
		}
		surfaces[endIndex-1] = string(verb)
		surfaces[endIndex] = "ましょう"
		return strings.Join(surfaces, ""), nil
	}

	// それ以外（文末は助動詞の終止形なはず）
	surfaces[endIndex] = surfaces[endIndex] + "です"
	return strings.Join(surfaces, ""), nil
}
