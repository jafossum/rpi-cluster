package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const (
	dir = "/home/pi/pixels/"
)

func main() {
	// Init rand generator
	rand.Seed(time.Now().UnixNano())

	log.Println("Starting")

	handleJob()
}

func handleJob() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	fi := createFile()
	defer deleteFile(fi)

	select {
	case <-c:
		fmt.Println("\r- Ctrl+C pressed in Terminal")
	case <-time.After(10 * time.Second):
		fmt.Println("\r- Job Done :)")
	}
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
