package goutil

import (
	"net"
	"strings"
)

func GetIntranetIp() (ips []string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip4 := ipnet.IP.To4(); ip4 != nil {
				if ip4[0] == 10 || // 10.x.x.x
					(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.x.x - 172.31.x.x
					(ip4[0] == 192 && ip4[1] == 168) { // 192.168.x.x
					ips = append(ips, ipnet.IP.String())
				}
			}
		}
	}
	return
}

func GetIntranetIpWithPrefix(prefix string) ([]string, error) {
	ips, err := GetIntranetIp()
	if err != nil {
		return ips, err
	}

	res := make([]string, 0)
	for _, ip := range ips {
		if strings.HasPrefix(ip, prefix) {
			res = append(res, ip)
		}
	}
	return res, nil
}
