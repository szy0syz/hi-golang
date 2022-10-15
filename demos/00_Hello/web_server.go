package hello

import (
	"fmt"
	"log"
	"net/http"
)

func handler_web(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, there, I love %s", r.URL.Path[1:])
}

func web() {
	http.HandleFunc("/", handler_web)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
