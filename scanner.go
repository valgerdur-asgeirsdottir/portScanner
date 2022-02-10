package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// taka strenginn úr command prompt
// brjóta strenginn niður í ports og hosts
// scanna hvert port á hverjum host og prenta hvort það sé opið eða lokað
func incrAddress(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] != 0 {
			break
		}
	}
}
func main() {
	var (
		inputArgsstr string
		ports        string
		hosts        []string
	)

	flag.StringVar(&ports, "ports", "1-5,6,7", "write ports with a comma as seperator")
	flag.Parse()
	hosts = flag.Args()

	portArr := strings.Split(ports, ",")
	var portsArr2 []string
	for i := 0; i < len(portArr); i++ {
		if strings.Contains(portArr[i], "-") {
			res := strings.Split(inputArgsstr, "-")
			int1, err1 := strconv.Atoi(res[0])
			int2, err2 := strconv.Atoi(res[1])
			if err1 != nil || err2 != nil {
				log.Fatal("Incorrect format of ports")
			}
			for j := int1; j < int2; j++ {
				numStr := strconv.Itoa(j)
				portsArr2 = append(portsArr2, numStr)
			}
		} else {
			//todo check if port is int
			portsArr2 = append(portsArr2, portArr[i])
		}

	}
	var allHosts [][]net.IP
	for i := 0; i < len(hosts); i++ {
		var hostArr []net.IP
		if ip, ipnet, err := net.ParseCIDR(hosts[i]); err == nil {
			first_ip := ip.Mask(ipnet.Mask)
			for ; ipnet.Contains(first_ip); incrAddress(first_ip) {
				hostArr = append(hostArr, first_ip)
			}
		}
		hostArr = hostArr[1 : len(hostArr)-1]
		allHosts = append(allHosts, hostArr)
	}

	//check if connection is open for every port with every host
	var (
		dest    string
		timeout int = 50
	)

	for i := 0; i < len(allHosts); i++ {
		for j := 0; j < int(allHosts[i]); j++ {

		}
		var address []string
		address = append(address, host)
		for j := 0; j < len(portArr); j++ {
			port := portsArr2[j]
			address = append(address, port)
			dest = strings.Join(address, ":")
			connection, err := net.DialTimeout("tcp", dest, time.Duration(timeout)*time.Millisecond)
			if err != nil {
				// handle error
			}

			fmt.Print(hosts[i], ":", portArr[j], " open")
		}
	}

	log.Default()
	sort.Strings(hosts)
	strings.Contains(ports, "a")
	type SafeCounter struct {
		mu sync.Mutex
		v  map[string]int
	}

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	panic("unimplemented")
}

// Scan represents the scan parameters
/*type Scan struct {
	minPort int
	maxPort int
	target  string
}

// NewScan creats scan object
func NewScan(args string, cfg cli.Config) (Scan, error) {
	var (
		scan Scan
		flag map[string]interface{}
		err  error
	)

	args, flag = cli.Flag(args)
	// help
	if _, ok := flag["help"]; ok || args == "" {
		help(cfg)
		return scan, fmt.Errorf("")
	}

	pRange := cli.SetFlag(flag, "p", cfg.Scan.Port).(string)

	re := regexp.MustCompile(`(\d+)(\-{0,1}(\d*))`)
	f := re.FindStringSubmatch(pRange)

	if len(f) != 4 {
		return scan, fmt.Errorf("error! please try scan help")
	}

	scan.target = args
	if len(f) == 4 && f[2] != "" {
		scan.minPort, err = strconv.Atoi(f[1])
		scan.maxPort, err = strconv.Atoi(f[3])
	} else {
		scan.minPort, err = strconv.Atoi(f[1])
		scan.maxPort, err = strconv.Atoi(f[1])
	}

	if err != nil {
		return scan, err
	}

	if !scan.IsCIDR() {
		ipAddr, err := net.ResolveIPAddr("ip", scan.target)
		if err != nil {
			return scan, err
		}
		scan.target = ipAddr.String()
	} else {
		return scan, fmt.Errorf("it doesn't support CIDR")
	}

	return scan, nil
}

// IsCIDR checks the target if it's CIDR
func (s Scan) IsCIDR() bool {
	_, _, err := net.ParseCIDR(s.target)
	if err != nil {
		return false
	}
	return true
}

// Run tries to scan wide range ports (TCP)
func (s Scan) Run() {
	if !s.IsCIDR() {
		host(s.target, s.minPort, s.maxPort)
	}
}
*/
// host tries to scan a single host
/*func host(ipAddr string, minPort, maxPort int) {
	var (
		wg      sync.WaitGroup
		tStart  = time.Now()
		counter int
	)

	var ports []int
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Protocol", "Port", "Status"})

	for i := minPort; i <= maxPort; i++ {
		wg.Add(1)
		go func(i int) {
			for {
				conn, err := net.DialTimeout("tcp", net.JoinHostPort(ipAddr, fmt.Sprintf("%d", i)), 2*time.Second)
				if err != nil {
					if strings.Contains(err.Error(), "too many open files") {
						// random back-off
						time.Sleep(time.Duration(10+rand.Int31n(30)) * time.Millisecond)
						continue
					}
					wg.Done()
					return
				}
				conn.Close()
				break
			}
			counter++
			ports = append(ports, i)
			wg.Done()
		}(i)
	}

	wg.Wait()

	sort.Ints(ports)
	for i := range ports {
		table.Append([]string{"TCP", fmt.Sprintf("%d", ports[i]), "Open"})
	}
	table.Render()

	if counter == 0 {
		println("there isn't any opened port")
	} else {
		elapsed := fmt.Sprintf("%.3f seconds", time.Since(tStart).Seconds())
		println("Scan done:", counter, "opened port(s) found in", elapsed)
	}
}*/

// help represents guide to user
/*func help(cfg cli.Config) {
	fmt.Printf(`
    usage:
          scan ip/host [option]
    options:
          -p port-range or port number      specified range or port number (default is %s)
    example:
          scan 8.8.8.8 -p 53
          scan www.google.com -p 1-500
	`,
		cfg.Scan.Port)
}*/
