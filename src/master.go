package main

import (
	"fmt"
	"net/http"
)

const (
	connHost = "localhost"
	connPort = "8080"
	message = "Hello tehre"
)

func serveHTTP(w http.ResponseWriter, r *http.Request) { 
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(message))	
	return
}

// Create master that serves at / with GET & POST options
func master(){
	fmt.Println("Master will use later")	
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", serveHTTP)
	fmt.Printf("Starting %s on port %s", connHost, connPort)
	http.ListenAndServe(connHost+":"+connPort, mux)
	master()
}

