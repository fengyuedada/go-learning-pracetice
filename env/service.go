package env

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func Service() string {
	appName := os.Args[0]
	pos := strings.LastIndex(appName, "/")
	if pos >= 0 {
		appName = appName[pos+1:]
	}
	return appName
}

// HostIP 获取IP
func HostIP() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return "127.0.0.1"
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println(err)
			return "127.0.0.1"
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// process IP address
			if ip != nil {
				return ip.String()
			}
		}
	}
	return ""
}

func PodName() string {
	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		hostname, _ := os.Hostname()
		if hostname != "" {
			os.Setenv("HOSTNAME", hostname)
		}
	}
	return os.Getenv("HOSTNAME")
}
