package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := "https://agung-setiawan.com/bukuvim/img/captain.png"

	err := download(url)

	if err != nil {
		fmt.Println("download success")
	} else {
		fmt.Println("download failed")
	}
}

func download(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	fileName := path.Base(response.Request.URL.String())
	
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)

	return err
}
