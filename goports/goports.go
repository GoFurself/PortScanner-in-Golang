package goports

// * All there is is here * //

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

const max_routines = 200

func Scan(ip net.IP, port_begin int, port_end int) {

	var wg sync.WaitGroup
	sp := make(chan struct{}, max_routines)
	defer close(sp)

	ip_str := ip.String()
	for port := port_begin; port <= port_end; port++ {
		wg.Add(1)
		sp <- struct{}{}
		go func(port int) {
			defer wg.Done()
			if conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip_str, port)); err == nil {
				defer conn.Close()
				fmt.Printf("Port %d is open\n", port)
			}
			<-sp
		}(port)
	}
	wg.Wait()
}

func ParseIP(ip_str string) (net.IP, error) {

	if ip := net.ParseIP(os.Args[1]); ip != nil {
		return ip, nil
	}
	return nil, errors.New("invalid IP address")
}

func ParsePorts(ports_str string) (int, int, error) {

	if strings.Contains(ports_str, "-") {

		var (
			port_range            = strings.Split(ports_str, "-")
			port_begin, err_begin = strconv.Atoi(port_range[0])
			port_end, err_end     = strconv.Atoi(port_range[1])
		)

		if err_begin != nil || err_end != nil {
			return 0, 0, errors.New("invalid port range")
		}

		return port_begin, port_end, nil
	}

	port, err := strconv.Atoi(ports_str)
	if err != nil {
		return 0, 0, errors.New("invalid port number")
	}

	return port, port, nil
}
