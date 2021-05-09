package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := "https://agung-setiawan.com/bukuvim/img/captain.png"
	fileName := "ega.png"

	err := download(fileName, url)

	if err != nil {
		fmt.Println("download gagal")
	} else {
		fmt.Println("download berhasil")
	}
}

func download(fileName string, url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)

	return err
}
