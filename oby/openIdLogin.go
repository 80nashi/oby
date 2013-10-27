package oby

import (
    "fmt"
    "net/http"

    "appengine"
    "appengine/user"
)

type Provider struct {
    ProviderName string
    ProviderId string
}

func openIdLogin(w http.ResponseWriter, r *http.Request) {
    var directProviders = []Provider{
        { "Google", "https://www.google.com/accounts/o8/id" },
        { "Yahoo",   "yahoo.com" },
        { "MySpace", "myspace.com" },
        { "AOL",     "aol.com" },
        { "MyOpenID","myopenid.com" },
    }
    var usernameProviders = []Provider{
        { "Blogger", "%s.blogspot.com" },
    }
    w.Header().Set("Content-type", "text/html; charset=utf-8")
    c := appengine.NewContext(r)
    u := user.Current(c)
    if u == nil {
        fmt.Fprintf(w, "| Sign in using ")
        for _, p := range directProviders {
            url, err := user.LoginURLFederated(c, "/openIdLogin", p.ProviderId)
            if err != nil {
                fmt.Fprintf(w, "hoge", err)
            }
            fmt.Fprintf(w, `<a href="%s">%s</a> | `, url, p.ProviderName)
        }
        for _, p := range usernameProviders {
            pid := fmt.Sprintf(p.ProviderId, "nashi")
            url, err := user.LoginURLFederated(c, "/openIdLogin", pid)
            if err != nil {
                fmt.Fprintf(w, "hoge2", err)
            }
            fmt.Fprintf(w, `<a href="%s">%s</a> | `, url, p.ProviderName)
        }
        return
    }
    url, _ := user.LogoutURL(c, "/openIdLogin")
    fmt.Fprintf(w, `Welcome, %s! Provider:%s Id:%s (<a href="%s">sign out</a>)`,
            u.Email, u.AuthDomain, u.FederatedIdentity, url)
}
