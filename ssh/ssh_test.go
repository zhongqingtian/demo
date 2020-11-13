package ssh

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"golang.org/x/crypto/ssh"
	"strings"
	"testing"
)

func TestGetPubKey(t *testing.T) {
	GetPubKey()
}

func TestTestKey(t *testing.T) {
	// GetKey()
	s := "hhhhdd"
	n := strings.ReplaceAll(s, "h", "a")
	t.Log(n)
}

func TestKey(t *testing.T) {
	privateKeyStr := `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA1PTSCscV/pWxBfWyg/CPs/6gS/uGxxm3c2yfhge8eQ8LlC0/
LkWvFPatZxnQhz31RnHFpRAx8FWjVb4Z+TNiinGM5AfEH0kT18mGliwMBgAwGGue
8/WYukWotg/PYkiQydjLAVxfkd8d7oXj20wvZK+MNt2NVWmr3hbhF2gg345jGTWL
1Unp1rwFshAiTqy85D0E871UzgCxcu8VACJihN6qraDkMaxkUhdP0CgzD6X57n8/
ATAQ95HRP0kCt50pmmooyC3xW7ZBFIRBMglbY8ildahI8A7Xk1SXUyIM9q6dJOzX
nQHFoxeg9iPWYjuQ/I9td7ZloI2oCabH69e90wIDAQABAoIBAEKoh368JpRUAt9c
sw49pJ2w3DgseVmlIRlOLPjlPwedwAAX0akIQu+K72hQWkebC0WLo1Qxp5AVHlVe
df7dKMExHeray3cHhuCKwkgLbZALSwK53T176NyVpWOnu0CGelvUMmcwpPGiJ1pc
KqCgTk1z7a3W18CenR+dhcOv4mpXM0ZrAqP16A3+SyzB1qtZHOy3w73VmI4eR9g+
Hp5yOibzLTfOODiX3AoV4YPNsD0ey3SpxhpnrHmzrl0Jd5+0ux0UlzOdYR2rZ1l3
+1yAw+NBi1Do6QEHjG1Ubw8Z++DDpH/LFqSzPH/BR+KYHTpQVDQW+8NRQNU0CJjU
s03MleECgYEA8d4TJ8uvx1tFL4FxBPVV179UU5n+xzGwUfZpZB5gp5f55cYyqfH6
zk2JJJlLkJI997EAiiIdAojc36HZshx+svYRW/7zRRXwzClotCuGYVQFM8Gh+30m
ByN+CbeFVuu4o9T36M3i1JeqxAT1QcGFzvO3pxXddsbrc/rKrRIQ7UsCgYEA4WZI
vScFezYTNmRfejreJGQrmdg3uN4LXO8h5VZz07gwhCyZB/jcO3gDwOtyb31LJkMe
x/zHJXocZXBP1Xh/PhaV/qz4BkKR9FiPLG4DYj0FOlCDEkw94TkX3Wf8HY/Fdq+t
kiP7QsWE4K/9fULfDtOqk6mj+BRAHmIa+UzWRJkCgYEAlb53P5k08AlK8VmDnp1C
E3wKemQKQEs8LdKQd4LzNm+6MQ8fiOYe8KFCel1KUmyoXvTAL8VlDo+rp+0tub84
lhYQs1NKTAiEd/JfN8UGwrVhLyT/45Mc+yh3PQ7ZR/JS82PEROdLO3HXMNsc+SQW
cwWjO5gm8AXB+qFX4AUiY7cCgYBxuRO2K/v+AHiNFvN9TjpKjifZhVgH9DAzgYvO
lGH4TpRb+8xyU0N61cC/E4w3aDYEZhUfvhAUfpnZLGN7R0Kb5vVi+45gSjIt4IjM
V+1pIXSDyZ94GfOWsSGzWgXAPIvaqCSg+XqKY/H5E5KIIrc8WdtUiWDrh/wZ2FHX
usbY8QKBgAVOh56ftwWDE3yJ7UFj0eQ4BZSKIdFML/nZk0Bv+hE7bsR8JP7BjAh+
o2qXoJZ3JaLzFjNNIhkhC15WntgJ/VOjPo7jX49xTRrYM7pq7HCIcXBQCw6LJc+T
+fRR8vsqU0f4NvyZILw7XC7sgX1gaFdYrzdcv2pQKXcWz53UANFD
-----END RSA PRIVATE KEY-----
`
	fmt.Println(privateKeyStr)
	privatePem, _ := pem.Decode([]byte(privateKeyStr))
	fmt.Println(privatePem)
	//var private *crypto.PrivateKey
	// x509.MarshalPKCS1PrivateKey
	privateKey, err := x509.ParsePKCS1PrivateKey(privatePem.Bytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(privateKey)
	// ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDU9NIKxxX+lbEF9bKD8I+z/
	// qBL+4bHGbdzbJ+GB7x5DwuULT8uRa8U9q1nGdCHPfVGccWlEDHwVaNVvhn5M2
	// KKcYzkB8QfSRPXyYaWLAwGADAYa57z9Zi6Rai2D89iSJDJ2MsBXF+R3x3uheP
	// bTC9kr4w23Y1VaaveFuEXaCDfjmMZNYvVSenWvAWyECJOrLzkPQTzvVTOALFy
	// 7xUAImKE3qqtoOQxrGRSF0/QKDMPpfnufz8BMBD3kdE/SQK3nSmaaijILfFbt
	// kEUhEEyCVtjyKV1qEjwDteTVJdTIgz2rp0k7NedAcWjF6D2I9ZiO5D8j213tm
	// WgjagJpsfr173T

	// ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDU9NIKxxX+lbEF9bKD8I+z/
	// qBL+4bHGbdzbJ+GB7x5DwuULT8uRa8U9q1nGdCHPfVGccWlEDHwVaNVvhn5M2
	// KKcYzkB8QfSRPXyYaWLAwGADAYa57z9Zi6Rai2D89iSJDJ2MsBXF+R3x3uheP
	// bTC9kr4w23Y1VaaveFuEXaCDfjmMZNYvVSenWvAWyECJOrLzkPQTzvVTOALFy
	// 7xUAImKE3qqtoOQxrGRSF0/QKDMPpfnufz8BMBD3kdE/SQK3nSmaaijILfFbt
	// kEUhEEyCVtjyKV1qEjwDteTVJdTIgz2rp0k7NedAcWjF6D2I9ZiO5D8j213tm
	// WgjagJpsfr173T

	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	//public := ssh.MarshalAuthorizedKey(publicKey)
	//fmt.Println("公钥:")
	//fmt.Println(string(public))
	f := ssh.FingerprintLegacyMD5(publicKey)
	fmt.Println(f)
}
