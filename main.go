package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
)

type newline []string

func (s newline) String() string {
	var str string
	for _, i := range s {
		str += fmt.Sprintf("%s\n", i)
	}
	return str
}

func ipResolve(host string) {
	addr, err := net.LookupHost(host)
	if err == nil {
		fmt.Print(newline(addr))
	}
}

func main() {
	var wg sync.WaitGroup

	version := flag.Bool("version", false, "current version")
	help := flag.Bool("help", false, "usage info")
	flag.Parse()

	if *version {
		fmt.Println("0.1.0")
		os.Exit(0)
	}
	if *help {
		fmt.Printf("Usage:\n echo hosts.txt |ipres\n")
		os.Exit(0)
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		host := s.Text()

		wg.Add(1)
		go func() {
			ipResolve(host)
			wg.Done()
		}()
		wg.Wait()
	}
}
