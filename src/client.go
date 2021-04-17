package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
)

const (
	connectHost = "localhost"
	connectPort = "8080"
	connectType = "tcp"
)

func client() {
	resp, err := http.Get("http://"+connectHost+":"+connectPort)
	if err != nil {
		fmt.Println("An error occured when sending")
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not read")
		os.Exit(1)
	}
	fmt.Printf("Got: %s\n", body)
}

func main() {
	fmt.Println("Hello Client")
	client()
	return
}
