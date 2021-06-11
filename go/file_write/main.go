package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://place.dog/300/200")
	fmt.Println(err)

	defer resp.Body.Close()
	file, err := os.Create("sample.jpeg")
	if err != nil {
		return
	}

	defer file.Close()

	io.Copy(file, resp.Body)
}
