package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	teinei := Teinei{}
	t.Run("文末が「です」のときは変換されずそのままであること", func(t *testing.T) {
		var body string = "私は寿司が食べたいです。"
		var convertedBody string = "私は寿司が食べたいです。"
		actualConvertedBody := teinei.Convert(body)
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})

	t.Run("文末が「ます」のときは変換されずそのままであること", func(t *testing.T) {
		var body string = "私は寿司を食べます。"
		var convertedBody string = "私は寿司を食べます。"
		actualConvertedBody := teinei.Convert(body)
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})

	t.Run("文末が「動詞の終止形」のときは「動詞の連用形+ます」に変換されること", func(t *testing.T) {
		var body string = "私は寿司を食べる。"
		var convertedBody string = "私は寿司を食べます。"
		actualConvertedBody := teinei.Convert(body)
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})

	t.Run("文末が「助動詞の断定のだ」のときは「だ→です」に変換されること", func(t *testing.T) {
		var body string = "私は寿司だ。"
		var convertedBody string = "私は寿司です。"
		actualConvertedBody := teinei.Convert(body)
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})

	t.Run("文末が「動詞の連用タ接続+助動詞の完了/過去/存続のた」のときは「動詞の連用形+まし+た」に変換されること", func(t *testing.T) {
		var body string = "私は家に着いた。"
		var convertedBody string = "私は家に着きました。"
		actualConvertedBody := teinei.Convert(body)
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})

	t.Run("文末が「動詞の連用タ接続(促音便形)+助動詞の完了/過去/存続のた」のときは「動詞の連用形+まし+た」に変換されること", func(t *testing.T) {
		var body string = "私は家に帰った。"
		var convertedBody string = "私は家に帰りました。"
		actualConvertedBody := teinei.Convert(body)
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})

	t.Run("文末が「名詞/形容詞詞」のときは「文末+です」に変換されること", func(t *testing.T) {
		var body string = "私は寿司。"
		var convertedBody string = "私は寿司です。"
		actualConvertedBody := teinei.Convert(body)
		assert.Equal(t, convertedBody, actualConvertedBody)
		t.Logf("convertedBody: %s", actualConvertedBody)
	})
}
