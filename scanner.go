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

/* Takes in an IP address and increments it by one */
func incrAddress(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] != 0 {
			break
		}
	}
}

/* Takes in host, port, timeout and channel
Checks connection for the port on the host
Reads from the channel */
func checkConnection(host string, port string, timeout int, c chan int) {
	dest := net.JoinHostPort(host, port) // combine host and port into string
	// try to get a connection for the port on the host, check if the channel is open
	connection, err := net.DialTimeout("tcp", dest, time.Duration(timeout)*time.Millisecond)
	if err == nil {
		// no error, channel is open
		fmt.Println(dest, " open")
		connection.Close()
	} else {
		// handle error, channel is closed
		fmt.Println(dest, " closed")
	}
	<-c // read from the input channel
}

/* This is the concurrent port scanner */
func main() {

	var (
		ports string
		hosts []string
	)
	// read console input from user
	flag.StringVar(&ports, "p", "", "write ports with a comma as seperator")
	flag.Parse()
	hosts = flag.Args()

	// terminate if no input was found
	if ports == "" || len(hosts) == 0 {
		log.Fatal("please supply ports and hosts")
	}

	portArr := strings.Split(ports, ",")

	var allPorts []string

	// add given ranges of ports one by one, if any, to the array
	for i := 0; i < len(portArr); i++ {
		if strings.Contains(portArr[i], "-") {
			res := strings.Split(portArr[i], "-")
			int1, err1 := strconv.Atoi(res[0])
			int2, err2 := strconv.Atoi(res[1])
			if err1 != nil || err2 != nil || int1 > int2 || int2 > 65536 { // 65536 is the largest port possible
				log.Fatal("incorrect format of ports")
			}
			for j := int1; j <= int2; j++ {
				numStr := strconv.Itoa(j) // convert int to string
				allPorts = append(allPorts, numStr)
			}
		} else {
			_, err := strconv.Atoi(portArr[i]) // check if port is int
			if err != nil {
				log.Fatal("incorrect format of ports")
			}
			allPorts = append(allPorts, portArr[i])
		}

	}

	var allHosts []string

	// scan all ip addresses in ranges, if any, to get all hosts
	for _, host := range hosts {
		var hostArr []net.IP
		if ip, ipnet, err := net.ParseCIDR(host); err == nil {
			first_ip := ip.Mask(ipnet.Mask)
			for ; ipnet.Contains(first_ip); incrAddress(first_ip) { // go through all ip addresses in range
				hostArr = append(hostArr, first_ip)
			}
			hostArr = hostArr[1 : len(hostArr)-1] // skip the first and last addresses in range since they are reserved
			for _, h := range hostArr {
				allHosts = append(allHosts, h.String())
			}
		} else {
			allHosts = append(allHosts, host) // if host is not an ip address range add straight away
		}
	}

	var timeout int = 500     // 500 ms timeout
	c := make(chan int, 1024) // max 1024 goroutines to safely be able to read the ports

	// check if connection is open for every port with every host
	for _, host := range allHosts {
		for _, port := range allPorts {
			c <- 1 // try to write to the channel
			go checkConnection(host, port, timeout, c)
		}
	}
	// write to the channel so the goroutines can start and finish before the program terminates
	for i := 0; i < 1024; i++ {
		c <- 1
	}
}
