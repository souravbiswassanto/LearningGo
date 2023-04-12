/*

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int

var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func echoHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {
	http.HandleFunc("/", echoString)
	http.HandleFunc("/hi", echoHi)
	http.HandleFunc("/increment", incrementCounter)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

*/

package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServe(":8081", nil))
}
