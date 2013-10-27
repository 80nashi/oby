package oby

import (
    "fmt"
    "net/http"
)

func openIdHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "hello, world");
}
