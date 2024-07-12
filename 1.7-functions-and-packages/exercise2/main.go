package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://www.google.com", nil)
	if err != nil {
		panic("The request could not be created")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic("The request could not be executed")
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
	fmt.Println()
}
