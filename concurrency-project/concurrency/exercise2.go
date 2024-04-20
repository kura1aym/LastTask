package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

const maxWorkers = 5

func downloadFile(url string, workerPool chan string) {
	defer func() {
		<-workerPool
	}()

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error downloading %s: %s\n", url, err)
		return
	}
	defer response.Body.Close()

	filename := getFileName(url)

	out, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file %s: %s\n", filename, err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, response.Body)
	if err != nil {
		fmt.Printf("Error writing file %s: %s\n", filename, err)
		return
	}

	fmt.Printf("%s downloaded successfully\n", url)
}

func getFileName(urlstr string) string {
	u, err := url.Parse(urlstr)
	if err != nil {
		log.Fatal("Error parsing URL: ", err)
	}

	ex, _ := url.QueryUnescape(u.EscapedPath())
	return filepath.Base(ex)
}

func readURLsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		if url != "" {
			urls = append(urls, url)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

func main() {
	filename := "urls.txt"
	urls, err := readURLsFromFile(filename)
	if err != nil {
		log.Fatalf("Error reading URLs from file %s: %s", filename, err)
	}

	workerPool := make(chan string, maxWorkers)

	for _, url := range urls {
		workerPool <- ""
		go downloadFile(url, workerPool)
	}

	for i := 0; i < cap(workerPool); i++ {
		<-workerPool
	}
}
