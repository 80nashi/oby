package oby

import (
    _ "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", welcome)
    http.HandleFunc("/api", apiHandler)
    http.HandleFunc("/openIdLogin", openIdLogin)
    http.HandleFunc("/_ah/login_required", openIdHandler)
}
