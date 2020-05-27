package models

// jodoshi 変換用 struct
type Word struct {
	Katsuyou string
	Pre      string
	Teinei   string
}

type Words []Word
