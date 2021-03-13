package goutil

import (
	"sync"
	"testing"
	"unicode/utf8"
)

func TestString2Float(t *testing.T) {
	v := String2Float("3.1415926", -1)
	if v != 3.1415926 {
		t.Fail()
	}
}

func TestString2Int(t *testing.T) {
	v := String2Int("123", -1)
	if v != 123 {
		t.Fail()
	}
}

func TestString2Int64(t *testing.T) {
	v := String2Int64("123456", -1)
	if v != 123456 {
		t.Fail()
	}
}

func TestString2Utf8(t *testing.T) {
	v := String2Utf8("你好")
	if !utf8.Valid([]byte(v)) {
		t.Fail()
	}
}

func TestIsSpace(t *testing.T) {
	if IsSpace('a') {
		t.Fail()
	}
	if !IsSpace(' ') {
		t.Fail()
	}
	if !IsSpace('\n') {
		t.Fail()
	}
	if !IsSpace('\t') {
		t.Fail()
	}
	if !IsSpace('\r') {
		t.Fail()
	}
}

func TestIsBlank(t *testing.T) {
	if !IsBlank(" ") {
		t.Fail()
	}
	if !IsBlank("\t\n\r") {
		t.Fail()
	}
	if IsBlank("abc") {
		t.Fail()
	}
}

func TestUnixNano(t *testing.T) {
	t.Log(UnixNano())
}

func TestUnixMsSec(t *testing.T) {
	t.Log(UnixMsSec(0))
}

func TestGetTime(t *testing.T) {
	t.Log(GetTime(FORMAT_DATE))
	t.Log(GetTime(FORMAT_TIME))
	t.Log(GetTime(FORMAT_DATE_TIME))
}

func TestParseTime(t *testing.T) {
	a := ParseTime(FORMAT_DATE, "2006-01-02")
	b := ParseTime(FORMAT_DATE_TIME, "2006-01-02 15:04:05")
	t.Log(a.Unix())
	t.Log(b.Unix())
}

func TestGetGuid(t *testing.T) {
	s := GetGuid()
	t.Log("my guid", len(s), s)
}

func TestSessionId(t *testing.T) {
	/*ma := make(map[int]str)
	ma[1] = "1"
	ll, ok := ma[1]
	kk := ma[2]
	t.Log(ll, ok)
	if kk != "" {
		t.Log(kk)
	}*/

	t.Log(10 / 3)
	t.Log(10 * 1.0 / 3)
	t.Log(10 / (3 * 1.0))
	//t.Log(SessionId())
}

func TestCheckMake(t *testing.T) {
	rwMutexAddRuleRes := &sync.RWMutex{}
	t.Log(CheckMake(rwMutexAddRuleRes) == nil)

	t.Log(rwMutexAddRuleRes == nil)
	//rwMutexAddRuleRes.Lock()
	t.Log("lock sucess")
}

func TestHashKey(t *testing.T) {
	t.Log(HashKey(MD5("jjkk/kdj/llj")))
}

func TestCreateTs(t *testing.T) {
	CreateTs()
}
