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

func main() {
	// Init rand generator
	rand.Seed(time.Now().UnixNano())

	log.Println("Starting")

	http.HandleFunc("/blinkt", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fi := createFile()
	defer deleteFile(fi)
	fmt.Fprintf(w, "Hello! File: "+fi)
	time.Sleep(time.Second)
}

func deleteFile(fi string) {
	log.Println("- Run Clean Up - Delete File")
	_ = os.Remove(dir + fi)
	log.Println("- Good bye!")
}

func createFile() string {
	log.Println("- Create File")

	fi := ""
	i := rand.Intn(8)
	log.Println(i)

	if _, err := os.Stat(dir + strconv.Itoa(i)); os.IsNotExist(err) {
		dp := []byte(fmt.Sprintf("%d:%d:%d", rand.Intn(255), rand.Intn(255), rand.Intn(255)))
		fi = strconv.Itoa(i)
		if err := ioutil.WriteFile(dir+fi, dp, 0666); err != nil {
			log.Println(err)
			os.Exit(3)
		}
	}
	return fi
}
