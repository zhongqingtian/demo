package emoji

import (
	"fmt"
	"github.com/go-xman/go.emoji"
	"strconv"
	"testing"
	"unicode/utf8"
)

func TestEmoji(t *testing.T) {
	s := "👩‍👩‍👦🇨🇳"
	_ = emoji.ReplaceEmoji(s, func(emoji string) string {
		hexDump(emoji)
		return ""
	})
}

func hexDump(s string)  {
	fmt.Printf("字符串：%s，占用字节数：%d, 字符数：%d\n======================\n", s, len(s), len([]rune(s)))
	for index, item := range []rune(s) {
		hex := strconv.FormatInt(int64(item), 16)
		fmt.Printf("序号%d：Unicode 码点：U+%s，字节数：%d\n", index, hex, utf8.RuneLen(item))
	}
}