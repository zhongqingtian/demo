package helpers

import (
	"crypto/sha1"
	"fmt"
	"io"
	_ "net/http/pprof"
	"regexp"
	_ "runtime/pprof"
)

const regular = "^((13[0-9]) | 14[57] |(15[^4,\\D]) |17[0-9] |(18[0-9]) ) \\d{8}$"

func CryptoSha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func ValidateMobile(mobileNum string) bool {
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
