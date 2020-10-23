package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	fmt.Println("Scraping your site!")
	// Machine Learning
	disco := httpRequest(os.Args[1], "index.html")
	if disco != nil {
		panic(disco)
	}
}

func httpRequest(url string, filename string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, byteArray, 0755)

}
