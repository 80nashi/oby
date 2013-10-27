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
        url, err := user.LoginURLFederated(c, "/", "https://www.google.com/accounts/o8/id")
        if err != nil {
            fmt.Fprintf(w, "hoge", err)
        }
        fmt.Fprintf(w, `<a href="%s">Sign in or register using twitter</a>`, url)
        return
    }
    url, _ := user.LogoutURL(c, "/")
    fmt.Fprintf(w, `Welcome, %s! %s@%s (<a href="%s">sign out</a>)`, u.Email, u.AuthDomain, u.FederatedIdentity, url)
}
