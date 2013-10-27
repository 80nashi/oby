package oby

import (
    "fmt"
    "net/http"

    "appengine"
    "appengine/user"
)

func welcome(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "text/html; charset=utf-8")
    c := appengine.NewContext(r)
    u := user.Current(c)
    if u == nil {
        url, _ := user.LoginURLFederated(c, "/", "twitter.com")
        fmt.Fprintf(w, `<a href="%s">Sign in or register using twitter</a>`, url)
        return
    }
    url, _ := user.LogoutURL(c, "/")
    fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)
}
