package oby

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/welcome", welcome)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}
