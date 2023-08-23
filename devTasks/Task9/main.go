package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args
	Wget(args[1], "index.html")
}

func Wget(url, fileName string) {
	resp := getResponse(url)
	if fileName == "" {
		urlSplit := strings.Split(url, "/")
		fileName = urlSplit[len(urlSplit)-1]
	}
	fileWriter(fileName, resp)
}

func getResponse(url string) *http.Response {
	tr := new(http.Transport)
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	return resp
}

func fileWriter(fileName string, resp *http.Response) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buffer := bufio.NewWriterSize(file, 1024*8)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(buffer, resp.Body)
	if err != nil {
		panic(err)
	}
}
