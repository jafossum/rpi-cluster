package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const (
	dir = "/home/pi/pixels/"
)

type singlePix struct {
	brigtness1 int
	brigtness2 int
	brigtness3 int
	fi         string
}

func main() {
	sp := singlePix{
		brigtness1: 0,
		brigtness2: 128,
		brigtness3: 0,
	}
	if err := sp.getBrightness(); err != nil {
		os.Exit(2)
	}
	log.Println(fmt.Sprintf("Brightness: %d, %d, %d", sp.brigtness1, sp.brigtness2, sp.brigtness3))

	sp.createFile()

	sp.blockUntilClose()
}

func (sp *singlePix) blockUntilClose() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	fmt.Println("\r- Ctrl+C pressed in Terminal")
	sp.deleteFile()
	os.Exit(0)
}

func (sp *singlePix) deleteFile() {
	log.Println("- Run Clean Up - Delete File")
	_ = os.Remove(dir + sp.fi)
	log.Println("- Good bye!")
}

func (sp *singlePix) createFile() {
	log.Println("- Create File")

	for i := 0; i < 8; i++ {
		log.Println(i)
		if i == 8 {
			log.Println("No more spots left")
			break
		}
		if _, err := os.Stat(dir + strconv.Itoa(i)); os.IsNotExist(err) {
			dp := []byte(fmt.Sprintf("%d:%d:%d", sp.brigtness1, sp.brigtness2, sp.brigtness3))
			sp.fi = strconv.Itoa(i)
			if err := ioutil.WriteFile(dir+sp.fi, dp, 0666); err != nil {
				log.Println(err)
				os.Exit(3)
			}
			break
		}
	}
}

func (sp *singlePix) getBrightness() error {
	if ev1 := os.Getenv("BRIGHTNESS_1"); ev1 != "" {
		i, err := strconv.Atoi(ev1)
		if err != nil {
			return err
		}
		sp.brigtness1 = i
	}
	if ev2 := os.Getenv("BRIGHTNESS_2"); ev2 != "" {
		i, err := strconv.Atoi(ev2)
		if err != nil {
			return err
		}
		sp.brigtness2 = i
	}
	if ev3 := os.Getenv("BRIGHTNESS_3"); ev3 != "" {
		i, err := strconv.Atoi(ev3)
		if err != nil {
			return err
		}
		sp.brigtness3 = i
	}
	return nil
}
