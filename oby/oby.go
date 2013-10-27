package oby

import (
    _ "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", welcome)
    http.HandleFunc("/protected", protected)
    http.HandleFunc("/_ah/login_required", openIdHandler)
}
