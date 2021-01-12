package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const (
	endpoint    = "/blinkt"
	defaultIP   = "127.0.0.1"
	defaultPort = 30000
)

func main() {
	// Init rand generator
	rand.Seed(time.Now().UnixNano())

	nt := flag.Int("n", 1, "Number of threads")
	ip := flag.String("i", defaultIP, "IP Address")
	po := flag.Int("p", defaultPort, "Port")

	flag.Parse()

	addr := "http://" + *ip + ":" + strconv.Itoa(*po) + endpoint
	i := 0
	for i < *nt {
		go loopReuest(addr, i)
		i++
	}

	blockUntilClose()
}

func loopReuest(addr string, i int) {
	for {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		// log.Println("Make new reuest - " + strconv.Itoa(i) + "...")
		if err := makeRequest(addr); err != nil {
			log.Println(err)
		}
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
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

	log.Println(string(body))
	return nil
}

func blockUntilClose() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	fmt.Println("\r- Ctrl+C pressed in Terminal")
	os.Exit(0)
}
