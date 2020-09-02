package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	teinei := Teinei{}
	t.Run("文末が「です」のときは変換されずそのままであること", func(t *testing.T) {
		var body string = "私は寿司が食べたいです。"
		var convertedBody string = "私は寿司が食べたいです。"
		actualConvertedBody, err := teinei.Convert(body)
		if err != nil {
			t.Errorf("convertedBody Error: %v", err)
		}
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})

	t.Run("文末が「ます」のときは変換されずそのままであること", func(t *testing.T) {
		var body string = "私は寿司を食べます。"
		var convertedBody string = "私は寿司を食べます。"
		actualConvertedBody, err := teinei.Convert(body)
		if err != nil {
			t.Errorf("convertedBody Error: %v", err)
		}
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})

	t.Run("文末が「動詞の終止形」のときは「動詞の連用形+ます」に変換されること", func(t *testing.T) {
		var body string = "私は寿司を食べる。"
		var convertedBody string = "私は寿司を食べます。"
		actualConvertedBody, err := teinei.Convert(body)
		if err != nil {
			t.Errorf("convertedBody Error: %v", err)
		}
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})

	t.Run("文末が「動詞の終止形/サ変」のときは「動詞の連用形+ます」に変換されること", func(t *testing.T) {
		var body string = "いったん気絶する。"
		var convertedBody string = "いったん気絶します。"
		actualConvertedBody, err := teinei.Convert(body)
		if err != nil {
			t.Errorf("convertedBody Error: %v", err)
		}
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})

	t.Run("文末が「動詞の終止形/カ変」のときは「動詞の連用形+ます」に変換されること", func(t *testing.T) {
		var body string = "外に来る。"
		var convertedBody string = "外に来ます。"
		actualConvertedBody, err := teinei.Convert(body)
		if err != nil {
			t.Errorf("convertedBody Error: %v", err)
		}
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})

	t.Run("文末が「助動詞の断定のだ」のときは「だ→です」に変換されること", func(t *testing.T) {
		var body string = "私は寿司だ。"
		var convertedBody string = "私は寿司です。"
		actualConvertedBody, err := teinei.Convert(body)
		if err != nil {
			t.Errorf("convertedBody Error: %v", err)
		}
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})

	t.Run("文末が「動詞の連用タ接続+助動詞の完了/過去/存続のた」のときは「動詞の連用形+まし+た」に変換されること", func(t *testing.T) {
		var body string = "私は家に着いた。"
		var convertedBody string = "私は家に着きました。"
		actualConvertedBody, err := teinei.Convert(body)
		if err != nil {
			t.Errorf("convertedBody Error: %v", err)
		}
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})

	t.Run("文末が「動詞の連用タ接続(促音便形)+助動詞の完了/過去/存続のた」のときは「動詞の連用形+まし+た」に変換されること", func(t *testing.T) {
		var body string = "私は家に帰った。"
		var convertedBody string = "私は家に帰りました。"
		actualConvertedBody, err := teinei.Convert(body)
		if err != nil {
			t.Errorf("convertedBody Error: %v", err)
		}
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})

	t.Run("文末が「動詞の連用ウ接続+助動詞の推量/意志のう」のときは「動詞の連用形+ましょう」に変換されること", func(t *testing.T) {
		var body string = "本を読もう。"
		var convertedBody string = "本を読みましょう。"
		actualConvertedBody, err := teinei.Convert(body)
		if err != nil {
			t.Errorf("convertedBody Error: %v", err)
		}
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})

	t.Run("文末が「名詞/形容詞詞」のときは「文末+です」に変換されること", func(t *testing.T) {
		var body string = "私は寿司。"
		var convertedBody string = "私は寿司です。"
		actualConvertedBody, err := teinei.Convert(body)
		if err != nil {
			t.Errorf("convertedBody Error: %v", err)
		}
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})

	t.Run("文末が「動詞+ない」のときは「動詞の連用形+ません」に変換されること", func(t *testing.T) {
		var body string = "私は学ばない。"
		var convertedBody string = "私は学びません。"
		actualConvertedBody, err := teinei.Convert(body)
		if err != nil {
			t.Errorf("convertedBody Error: %v", err)
		}
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})

	t.Run("文末が「動詞以外+ない」のときは「ない+です」に変換されること", func(t *testing.T) {
		var body string = "それじゃない。"
		var convertedBody string = "それじゃないです。"
		actualConvertedBody, err := teinei.Convert(body)
		if err != nil {
			t.Errorf("convertedBody Error: %v", err)
		}
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})

	t.Run("文末が「助詞」のときは「その1つ前を文末として」変換されること", func(t *testing.T) {
		var body string = "あっちに行こうよ。"
		var convertedBody string = "あっちに行きましょうよ。"
		actualConvertedBody, err := teinei.Convert(body)
		if err != nil {
			t.Errorf("convertedBody Error: %v", err)
		}
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})

	t.Run("文末が「記号」でないときも動作すること", func(t *testing.T) {
		var body string = "私は寿司が食べたいです"
		var convertedBody string = "私は寿司が食べたいです"
		actualConvertedBody, err := teinei.Convert(body)
		if err != nil {
			t.Errorf("convertedBody Error: %v", err)
		}
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})

	t.Run("kagomeの辞書に存在しない語が含まれていても動作すること", func(t *testing.T) {
		var body string = "隣で酔っててワロタ。"
		var convertedBody string = "隣で酔っててワロタです。"
		actualConvertedBody, err := teinei.Convert(body)
		if err != nil {
			t.Errorf("convertedBody Error: %v", err)
		}
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})
}
