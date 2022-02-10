package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

// Increment IP address
func incrAddress(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] != 0 {
			break
		}
	}
}

func checkConnection(host string, port string, timeout int, c chan int) {
	dest := net.JoinHostPort(host, port) // combine host and port into string
	connection, err := net.DialTimeout("tcp", dest, time.Duration(timeout)*time.Millisecond)
	if err == nil {
		fmt.Println(dest, " open")
		connection.Close()
	} else {
		// handle error
		fmt.Println(dest, " closed")
	}
	<-c
}

func main() {
	var (
		ports string
		hosts []string
	)

	flag.StringVar(&ports, "p", "", "write ports with a comma as seperator")
	flag.Parse()
	hosts = flag.Args()

	if ports == "" || len(hosts) == 0 {
		log.Fatal("please supply ports and hosts")
	}
	portArr := strings.Split(ports, ",")
	var portsArr2 []string
	for i := 0; i < len(portArr); i++ {
		if strings.Contains(portArr[i], "-") {
			res := strings.Split(portArr[i], "-")
			int1, err1 := strconv.Atoi(res[0])
			int2, err2 := strconv.Atoi(res[1])
			if err1 != nil || err2 != nil || int1 > int2 || int2 > 65536 {
				log.Fatal("incorrect format of ports")
			}
			for j := int1; j <= int2; j++ {
				numStr := strconv.Itoa(j)
				portsArr2 = append(portsArr2, numStr)
			}
		} else {
			//todo check if port is int
			_, err := strconv.Atoi(portArr[i])
			if err != nil {
				log.Fatal("incorrect format of ports")
			}
			portsArr2 = append(portsArr2, portArr[i])
		}

	}
	var allHosts []string
	for _, host := range hosts {
		var hostArr []net.IP
		if ip, ipnet, err := net.ParseCIDR(host); err == nil {
			first_ip := ip.Mask(ipnet.Mask)
			for ; ipnet.Contains(first_ip); incrAddress(first_ip) {
				hostArr = append(hostArr, first_ip)
			}
			hostArr = hostArr[1 : len(hostArr)-1]
			for _, h := range hostArr {
				allHosts = append(allHosts, h.String())
			}
		} else {
			allHosts = append(allHosts, host)
		}
	}

	var timeout int = 500     // 0.5 second timeout
	c := make(chan int, 1024) // max 1024 goroutines otherwise we fill the buffer

	//check if connection is open for every port with every host
	for _, host := range allHosts {
		for _, port := range portsArr2 {
			c <- 1
			go checkConnection(host, port, timeout, c)
		}
	}
	for i := 0; i < 1024; i++ {
		c <- 1
	}
}
