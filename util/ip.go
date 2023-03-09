package util

import "net"

var LocalIP = net.ParseIP("127.0.0.1")

// SetLocalIPs Set Local IP
func SetLocalIPs() []net.IP {
	ips := GetLocalIPs()
	if len(ips) > 0 {
		LocalIP = ips[0]
	}
	return ips
}

// GetLocalIPs Get Local IP
func GetLocalIPs() (ips []net.IP) {
	interfaceAddress, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}
	for _, address := range interfaceAddress {
		ip, isValidIP := address.(*net.IPNet)
		if isValidIP && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				ips = append(ips, ip.IP)
			}
		}
	}
	return ips
}
