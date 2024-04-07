package client

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
)

func get(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("Response status: ", res.Status)

	scanner := bufio.NewScanner(res.Body)

	content, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
