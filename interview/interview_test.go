package interview

import (
	"net/url"
	"testing"
)

func TestMapTest(t *testing.T) {
	//t.Log()
	//MapTest()

	// UpSafePoint()
	kk, err := url.Parse("https://clientID:clientSecret@2tfg8nvuw6mgileh8v5.xbase.cloud/v1/auth/token")

	t.Log(kk, err)
}
