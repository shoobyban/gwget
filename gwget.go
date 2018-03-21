package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
)

func get(url, filename string) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic in %s\n%v\n", url, r)
		}
	}()

	var err error

	myClient := &http.Client{}
	var res *http.Response
	done := make(chan bool)
	go func(done chan bool) {
		res, err = myClient.Get(url)
		done <- true
	}(done)
	select {
	case <-done:
		//all OK
	case <-time.After(10 * time.Second):
		err = errors.New("Timeout")
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	file, _ := os.Create(filename)
	defer file.Close()
	file.Write(body)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("gwget file url")
		fmt.Println("   url     prefix for the download")
		fmt.Println("   file    directory list as argument (dir/dir/file format per line, from find)")
		os.Exit(0)
	}

	url := os.Args[2]

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Panic(err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Panic(err)
	}

	lines := strings.Split(string(b), "\n")
	for _, fullpath := range lines {
		if _, err := os.Stat(fullpath); os.IsNotExist(err) {
			dir := filepath.Dir(fullpath)
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				err := os.MkdirAll(dir, 0777)
				if err != nil {
					panic(err)
				}
				fmt.Println("Created", dir)
			}
			get(url+fullpath, fullpath)
			fmt.Println("Downloaded", fullpath)
		}
	}
}
