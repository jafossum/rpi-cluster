package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	dir = "/home/pi/pixels/"
)

type singlePix struct {
	fi string
}

func main() {
	// Init rand generator
	rand.Seed(time.Now().UnixNano())

	log.Println("Starting")
	sp := singlePix{}

	http.HandleFunc("/blinkt", sp.handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func (sp *singlePix) handler(w http.ResponseWriter, r *http.Request) {
	sp.createFile()
	defer sp.deleteFile()
	fmt.Fprintf(w, "Hello! File: "+sp.fi)
	time.Sleep(time.Second)
}

func (sp *singlePix) deleteFile() {
	log.Println("- Run Clean Up - Delete File")
	_ = os.Remove(dir + sp.fi)
	log.Println("- Good bye!")
}

func (sp *singlePix) createFile() {
	log.Println("- Create File")

	i := rand.Intn(8)
	log.Println(i)

	if _, err := os.Stat(dir + strconv.Itoa(i)); os.IsNotExist(err) {
		dp := []byte(fmt.Sprintf("%d:%d:%d", rand.Intn(255), rand.Intn(255), rand.Intn(255)))
		sp.fi = strconv.Itoa(i)
		if err := ioutil.WriteFile(dir+sp.fi, dp, 0666); err != nil {
			log.Println(err)
			os.Exit(3)
		}
	}
}

