package util

import (
	"encoding/json"
	"log"
	"net"
)

var (
	localIP = ""
)

func LocalIP() string {
	if localIP == "" {
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			log.Printf("[ERROR]:get InterfaceAddres failed,err:%s \n", err.Error())
			return ""
		}
		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					localIP = ipnet.IP.String()
					log.Printf("InitLocalIp, LocalIp:%s \n", localIP)
					break
				}
			}
		}
	}
	return localIP
}

func ToJsonString(object interface{}) string {
	js, _ := json.Marshal(object)
	return string(js)
}
