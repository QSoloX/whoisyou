package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		domain := strings.ToLower(sc.Text())
		ips, err := net.LookupIP(strings.Replace(strings.Replace(domain, "http://", "", 1), "https://", "", 1))
		if err != nil {
			continue
		}
		for _, ip := range ips {
			output := domain + " " + ip.String()
			fmt.Println(output)
		}
	}

}
