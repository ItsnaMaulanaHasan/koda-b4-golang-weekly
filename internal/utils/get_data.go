package utils

import (
	"fmt"
	"io"
	"net/http"
)

func GetData(url string) []byte {
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Failed to get data")
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Failed ro read body")
	}

	return body
}
