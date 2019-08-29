package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"net/http"
)

func main() {
	var addr string
	kingpin.Flag("addr", "listening addr").Default(":8090").StringVar(&addr)
	kingpin.Parse()
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "helloworld")
	}))
	log.Fatal(http.ListenAndServe(addr, nil))
}
