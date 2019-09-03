package goutil

import "testing"

func TestGetIntranetIp(t *testing.T) {
	t.Log(GetIntranetIp())
}

func TestGetIntranetIpWithPrefix(t *testing.T) {
	t.Log(GetIntranetIpWithPrefix("10"))
	t.Log(GetIntranetIpWithPrefix("172"))
	t.Log(GetIntranetIpWithPrefix("172.15"))
	t.Log(GetIntranetIpWithPrefix("172.16"))
	t.Log(GetIntranetIpWithPrefix("172.18"))
	t.Log(GetIntranetIpWithPrefix("192"))
	t.Log(GetIntranetIpWithPrefix("192.168"))
	t.Log(GetIntranetIpWithPrefix("192.169"))
}
