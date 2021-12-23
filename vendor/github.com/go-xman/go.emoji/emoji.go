// Package emoji is designed to recognize and parse
// every indivisual Unicode Emoji characters from a string.
//
// Unicode Emoji Documents: http://www.unicode.org/Public/emoji/13.0/
package emoji

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/go-xman/go.emoji/official"
)

// replaceAllEmojiFunc searches string and find all emojis.
func replaceAllEmojiFunc(s string, f func(emoji string) string) string {
	buff := bytes.Buffer{}
	nextIndex := 0

	for i, r := range s {
		if i < nextIndex {
			continue
		}

		match, length := official.AllSequences.HasEmojiPrefix(s[i:])
		if false == match {
			buff.WriteRune(r)
			continue
		}

		nextIndex = i + length
		if f != nil {
			buff.WriteString(f(s[i : i+length]))
		}
	}

	return buff.String()
}

func HasEmoji(s string) bool {
	for i := range s {
		match, _ := official.AllSequences.HasEmojiPrefix(s[i:])
		if match {
			return true
		}
	}
	return false
}

func FilterEmoji(s string) string {
	return replaceAllEmojiFunc(s, func(emoji string) string {
		return ""
	})
}

func ReplaceEmoji(s string, f func(emoji string) string) string {
	return replaceAllEmojiFunc(s, f)
}

func HumanReadLen(s string) int {
	count := 0
	nextIndex := 0
	for i := range s {
		if i < nextIndex {
			continue
		}

		count++
		match, length := official.AllSequences.HasEmojiPrefix(s[i:])
		if false == match {
			continue
		}

		nextIndex = i + length
	}

	return count
}

func Dump(s string) {
	fmt.Printf("string：%s，byte count：%d, character count：%d\n======================\n", s, len(s), len([]rune(s)))
	for index, item := range []rune(s) {
		hex := strconv.FormatInt(int64(item), 16)
		fmt.Printf("seq %d：Unicode code：U+%s，byte count：%d\n", index, hex, utf8.RuneLen(item))
	}
}
