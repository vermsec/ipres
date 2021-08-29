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

func ipResolve(host string, verbose bool, mapBool bool, outDir string) {
	addr, err := net.LookupHost(host)
	if verbose {
		fmt.Print(newline(addr))
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	} else {
		if err == nil {
			fmt.Print(newline(addr))
		}
	}
	if mapBool {
		file, err := os.OpenFile(outDir, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			fmt.Printf("Error while creating file: %v", err)
		}
		defer file.Close()
		fmt.Fprintf(file, "%s : %v\n", host, addr)
	}

}

func main() {
	var wg sync.WaitGroup

	version := flag.Bool("version", false, "current version")
	help := flag.Bool("help", false, "usage info")
	mapBool := flag.Bool("map", false, "generates IPmap")
	v := flag.Bool("v", false, "verbose")
	outDir := flag.String("o", "ipres.map", "output")
	flag.Parse()

	if *version {
		fmt.Println("0.1.1")
		os.Exit(0)
	}
	if *help {
		fmt.Println("IPres - IP Resolver (By @vermsec)")
		fmt.Printf("Usage: echo hosts.txt |ipres\n")
		os.Exit(0)
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		host := s.Text()

		wg.Add(1)
		go func() {
			ipResolve(host, *v, *mapBool, *outDir)
			wg.Done()
		}()
		wg.Wait()
	}
}
