package oby

import (
    "fmt"
    "net/http"
)

func protected(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "hello, world");
}
