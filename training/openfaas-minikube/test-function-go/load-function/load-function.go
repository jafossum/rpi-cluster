package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	defaultNum = 20000
	defaultIP = "127.0.0.1"
	defaultPort = 8080
)

func main() {
	nt := flag.Int("n", defaultNum, "Number of invocations")
	ip := flag.String("i", defaultIP, "IP Address")
	po := flag.Int("p", defaultPort, "Port")

	flag.Parse()

	for i := 0; i < *nt; i++ {
		makeRequest("http://" + *ip + ":" + strconv.Itoa(*po) + "/function/test-function-go")
	}
}

func makeRequest(addr string) error {
	resp, err := http.Get(addr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Print(string(body))
	return nil
}