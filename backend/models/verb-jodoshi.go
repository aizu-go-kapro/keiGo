package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/ikawaha/kagome/tokenizer"
)

// EndOfJodoshi for convert end of jodoshi
type EndOfJodoshi struct {
	Hiragana string `json:"hiragana"`
	Romaji   string `json:"romaji"`
}

func convert() {
	dic := tokenizer.SysDicSimple()
	t := tokenizer.NewWithDic(dic)
	morphs := t.Tokenize("帰った")

	fmt.Println("変換前")
	for _, string := range morphs {
		fmt.Print(string.Surface)
	}

	jodoshiConvertSlice := jodoshiHenkan(morphs)

	fmt.Println("\n\n変換後")
	for _, string := range jodoshiConvertSlice {
		fmt.Println(string)
	}

}

func jodoshiHenkan(tokens []tokenizer.Token) []string {

	var slice []string

	tokens = tokens[1 : len(tokens)-1]
	endOfSentenceTokenIndex := len(tokens) - 1
	for {
		if tokens[endOfSentenceTokenIndex].Features()[0] == "記号" {
			endOfSentenceTokenIndex--
		} else {
			break
		}
	}

	for _, string := range tokens {
		slice = append(slice, string.Surface)
	}
	jodoshi := slice[endOfSentenceTokenIndex]

	var alphabet string
	bytes, err := ioutil.ReadFile("jodoshi.json")
	if err != nil {
		log.Fatal(err)
	}
	// JSONデコード
	var endOfJodoshi []EndOfJodoshi
	if err := json.Unmarshal(bytes, &endOfJodoshi); err != nil {
		log.Fatal(err)
	}

	for _, e := range endOfJodoshi {
		if strings.HasSuffix(jodoshi, e.Hiragana) {
			alphabet = e.Romaji
			break
		} else {
			alphabet = ""
		}
	}
	if alphabet != "" {
		if strings.HasSuffix(alphabet, "i") {
			jodoshi = jodoshi + "です"
		} else if strings.HasSuffix(alphabet, "katta") {
			jodoshi = jodoshi + "です"
		} else if strings.HasSuffix(alphabet, "ta") {
			jodoshi = "ました"
		} else if strings.HasSuffix(alphabet, "u") {
			fmt.Println(len([]rune(jodoshi)))
			jodoshi = jodoshi[:len(jodoshi)-3] + "ます"
		}
	}
	slice[endOfSentenceTokenIndex] = jodoshi

	return slice
}
