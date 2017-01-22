package main

import (
	"log"
	"net/http"
	"sync"
)

const (
  port = ":1234"
)

var calls = 0
var mutex = &sync.Mutex{}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	calls++
	mutex.Unlock()
	log.Printf("request from %v\ncalls: %d\n", r.RemoteAddr, calls)
	w.Write([]byte("howdy\n"))
}

func main() {
	log.Printf("Server at http://localhost%v.\n", port)
	http.HandleFunc("/", helloWorld)
	log.Fatal(http.ListenAndServe(port, nil))
}
