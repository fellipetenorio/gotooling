// simple hello world
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// to check error: errcheck
// install go-wrk
// go-wrk -d 5 http://localhost:8080/campoy@golang.org

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	if strings.HasSuffix(path, "@golang.org") {
		name := strings.TrimSuffix(path, "@golang.org")
		fmt.Fprintf(w, "gopher %s", name)
		return
	}
	fmt.Fprintf(w, "Hello %s", path)
}

// cmd: go run main.go
// a http server will listen 8080 and return the string to browser
