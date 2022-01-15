# [go.emoji](https://github.com/Andrew-M-C/go.emoji)

[![GoDoc](https://godoc.org/github.com/Andrew-M-C/go.emoji?status.svg)](https://godoc.org/github.com/Andrew-M-C/go.emoji)
[![](https://goreportcard.com/badge/github.com/Andrew-M-C/go.emoji)](https://goreportcard.com/report/github.com/Andrew-M-C/go.emoji)

This Package `emoji` is designed to recognize and parse every individual Unicode Emoji characters from a string.

## Example

replace emoji
```go
func main() {
	printf := fmt.Printf

	s := "👩‍👩‍👦🇨🇳"
	i := 0

	final := emoji.ReplaceAllEmojiFunc(s, func(emoji string) string {
		i++
		printf("%02d - %s - len %d\n", i, emoji, len(emoji))
		return fmt.Sprintf("%d-", i)
	})

	printf("final: <%s>\n", final)
}

// Output:
// 01 - 👩‍👩‍👦 - len 18
// 02 - 🇨🇳 - len 8
// final: <1-2->
```
check emoji
```go
emoji.HasEmoji("👩‍👩‍👦")

// Output:
// true
```
filter emoji
```go
emoji.FilterEmoji("1⃣️23")

// Output:
// 23
```
human read length
```go
emoji.HumanReadLen("👩‍👩‍👦23")

// Output:
// 3
```
dump emoji
```go
emoji.Dump("👨‍👩‍👧‍👦")
string：👨‍👩‍👧‍👦，byte count：25, character count：7
======================
seq 0：Unicode code：U+1f468，byte count：4
seq 1：Unicode code：U+200d，byte count：3
seq 2：Unicode code：U+1f469，byte count：4
seq 3：Unicode code：U+200d，byte count：3
seq 4：Unicode code：U+1f467，byte count：4
seq 5：Unicode code：U+200d，byte count：3
seq 6：Unicode code：U+1f466，byte count：4
```
