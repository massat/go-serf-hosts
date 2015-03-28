package main

import (
	"bufio"
	"github.com/massat/go-serf-hosts"
	"log"
	"os"
)

func main() {
	hostsFile := "/etc/hosts" // 引数からとりたい
	event := os.Getenv("SERF_EVENT")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		data := scanner.Text()
		goSerfHosts.NewSerfHosts(hostsFile).HandleEvent(event, data)

		log.Println(data, event)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
