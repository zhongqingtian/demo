package ssh

import (
	"crypto/x509"
	"encoding/pem"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"strings"
)

func GetPubKey() (string, error) {
	/*	str := `-----BEGIN RSA PRIVATE KEY-----
		MIIEowIBAAKCAQEA2YmtN/vn5hPV+B9UcS7GggH/qPzP+CZtW+qHW3GvAmBYbOSm
		WHVy5Q/ppyMcGALdaijCUifRU2KWTv/9ziB6W0mSIeG9RqeFO0D1xjFI2hatKc+u
		WqZ46zSHVgyyRx+/eIbVmMRdGFNrbS0kTPyiL7CT3suUwpLfuy1NQJuvjrSsCZsa
		jSiqb8R89QdOTeXgY4N8/mtkNbDAfsyrHZUQP35DnguvZEmK3i8YCqdM24YNqZPB
		+wSM8D5IJIRVH2f19eFIt9+LwnWuvJOULms+uJOqRcnz1ycfaeBDicfLVjM2kXhR
		laSh+PIfrWzWper5jDnbwLrSNXX0I1enNtkItQIDAQABAoIBAAYWbjyE1iaAI5q/
		9bL9Jhz5RVhLSt6c5K6Pz53lfopIe/EALQ+IqeBzI4zD1MNvQRb5cQavipIUzjQD
		weyEyCfJ1UcRdCSIB5hy5u3W7zz7KXlkEHME+ZlLbQYhcNPAtbD/xPZzeE939iel
		/cARydAf1uODOvR++7fXOrQ2JyGNK5H387a7Ei9ITCP4oMhymWkupvpDqBzZGShl
		wCl6EmnWJFQOTiKtxre183FXFTsCsJxeQ2sqRk4rdcpzNn0Uosc+42PYTBxjp68t
		Gn33zdENsGNvH4jmfaEKy+DC2MJtBhIvLpKd0vxI9noBWVh7RIOhxnlT6D4NPXx6
		ggOxqAECgYEA9+cINrMJbE2UV0EdH+uBJ/h4pPq01Cbgqzvlydp33hOebG6GbeNl
		FLvkP7TZYni06xdBoY5a+O5V/pInVaRqwx+tWljNYsfqJpi7uvvYnBD43tlN3UaP
		X9YUlwgA/b7295OKaG8gCM3EgZKp8A9hRXWqG6bG4mCnH81K7dv9X7UCgYEA4KS7
		+PlHmEGqyLOHP9JtZPn4HfVjCbTAK2LhN+m/IW8Tvk/Q4nKqwdxn/z0p9mCYPKp8
		lybplQ/mDzzL+iKdLvHjKXQNa5rSblcZx3XPzc56fRlcIJyX9CBwQaVoyhAlkFde
		T6JqIfKDfvs9RLFAs2d7E7LisC+cXStWWz78pQECgYAd3DA7FnmuF/IkSG2PqLhX
		7i/p/2DEpSIuvfpBxBToJQjSGsYKend0deVaXxhIKi3wm06GfTUsAtlM3qHqSPGr
		KyegWE309guYNpF2RlujANBZV2dvnc4rRwgX4Rjtd5Eb5SOozKnZJZ/j05ywk8mt
		qiyhRAd0wRugUwQRfowYeQKBgQCGJdmadEIsygoSkCLUGfT5t8xTF5Zs1WNshT+5
		A/z/GjXQEBnNIGyYF+aCpSPJ3SUMFUMRgEs8mKHpt4SXetQilI/y0TqyZ2ICQkFC
		uze8zvCWg+5gCyMyxaxDh2NsRvgnmgTXyUVXwFqGGcIcozzf5Yu517lHPSdC+/41
		vwdgAQKBgCbDYen9W+oMZ+YefH9vxarXsG0XVLR3VDyCs0xAPshBDHZKPnxJ/TvU
		wuX59RfY/TMqMsqg3V+lWr+xzotPtE8rHS4p1n9ABybx1mhDwDEf/WLMH6VUlYvS
		//DOOX8AUJ1xO96Z21ImqjzmJanq9vmQjHMILNSaF9SDiAe1c6fl
		-----END RSA PRIVATE KEY-----`
	fmt.Println(str)*/
	str2 := `-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAuMqNkAABPHo2ohE6HVr9W6uinGHgSWg4XwXH1unilYa+NmF5\nVrtPrmRieOypFgpx40wYmQkAb2aytENovYuYI2pDbsQ7fILbxFFuZTquF48Tywj3\npjoWjK4gXk4ddvDBW8vBwGFKxjuo+S4j7Usl1pUfMf54Nk3wIoNZAcHRaXjY63Zg\numAG+qajpyG491Mh6x7ES9InpgyF+nnyGAx1Q+3dM+aBdVrpH3T2lcxjWwPH3xD3\nsmt8mjNf3s4eLk9pPsc2JLlvvjDbFNp3iVh5XgLSbQYmwNmahhWarSya29rB5+rq\nOd5QT1lc+fjorfpQouCgsKIHZwG7z0ERDNHMTQIDAQABAoIBABusCWLk/UpnKsM+\nPSukMoiFvUYZyhJ0Duz72p8uSwYFNXrDQf5g8l+1HKHrHq70RXW6MXzJQk1sXZps\n9cRcoR28jfiKXjYZCnvtjdCO5v8UJ+yhnSRBUpsigCpfHD5xcyMh8hOy/hWNXZJi\n/HJObI4jqpY7+xh2t10TukJvGcLnWmGMpF8UYoj3/DUoAl+dh218XcM6YjnKB4yA\n/HZtlcEZoeygwkTSVaY0lUrseUhcpOUX6dS9Lo/EGgX9HgX5jcfBa7Q5H+PFv+/x\n91XlMMM1HD/5/FFgS/2A4Q17rplWQwyxHRSdQo4a8DkmjG15BWk7e3kgth06jyt7\nwi0b0GkCgYEA8fPpYLjnfQlODEpffCA7e/oY8tLHLZ8P+qcxdctHQYcFOL2CNS2Q\n5V3efrWISNAtwVy8s6iiSEnm+1ZJ2BFq0iVEzMlNQZ6ZFr4A2Wy5NGZQq0eHtWK1\n3VCmeb8GltTQMsOuZ3XHI33bxODpt6RZhPfFfkd22vZYvHXv0kRDxH8CgYEAw4UP\n8yB+FDYKdib3JpLAcb6hq7ZJfbNnJUQjyNleQ+STAQ7eMlzZp0vVS2M4QXCLSOBs\nZYO73fJtNp747S+iFFOYi78mlPrf0mGMKdPiB8LnHrf7b3Q+aiihMptIZNRBzCfv\nq3dJb9JWqp21XXx4Lxk23uLt8gLFk5BVOBU+2TMCgYAiXhAPUhRnmVPBwSqw+e8Z\nHBxBGZl7LHXbl7YZ0kcL0r72IU0GerNzMRMyklyKYy+soJyZNGHgvMaXetlAXLD8\nqpiMMifMI1gBd/ms6JmiDMp2mdF3/U5x4kvTTUQFVrNAErWKMpuXuf/mhkAMesbQ\ncVVXK5WBYq6WDnaDOnkZxwKBgG9gnF5eUxQGWN1HgyD5FiF1e7lKZQDzMr263Dus\n/5LjYY6HvrGnzOENRNrmEYTNrLp+BaGqclmToP0tgBX1PXogyNRKtprFcrzwjpzB\nM5BCSdBo9BVgfbGRpYojhRDTDQxUb6YYlL/bbT1SC5/OqNQDpAtB8FQdmhlebf09\ntQibAoGBAISRPChOKnSsMkiREClTDM6ylZfQ801BEsYu1Z/4sMxYWU2yP5EpBaOk\nWDZb8dONivEBQQeOcwnfShZRtHJqkmkA2WyYxYQZlVStL5V7v2LcBIegWdW9B2E0\nL+uMtVi0u20NwiX7t+7p6hIn8KPOjX8+7BE4ZGwEpZAbbRmLD94R\n-----END RSA PRIVATE KEY-----\n`
	str2 = strings.ReplaceAll(str2, "\\n", "\n")
	r := strings.NewReader(str2)
	pemBytes, err := ioutil.ReadAll(r)
	if err != nil {
	}
	pri, _ := pem.Decode(pemBytes)
	private, err := x509.ParsePKCS1PrivateKey(pri.Bytes)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("get publicKey by privateKey fail")
		return "", err
	}
	publicKeys, err := ssh.NewPublicKey(&private.PublicKey)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("privateKey change publicKey fail")
		return "", err
	}
	fingerprint := ssh.FingerprintLegacyMD5(publicKeys) // 计算指纹
	return fingerprint, nil
}
