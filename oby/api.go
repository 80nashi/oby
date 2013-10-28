package oby

import (
    "encoding/json"
    "fmt"
    _ "io/ioutil"
    "net/http"

    "appengine"
)

type ApiMessage struct {
    Cmd string
    Arg string
    Err string
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "text/javascript; charset=utf-8")
    email := getEmail(appengine.NewContext(r))
    var m ApiMessage
    if email == "" {
        m.Err = "nouser"
        json.NewEncoder(w).Encode(m)
        return
    }
    dec := json.NewDecoder(r.Body)
    if err := dec.Decode(&m); err != nil {
        m.Err = fmt.Sprint(err)
        json.NewEncoder(w).Encode(m)
        return
    }
    m.Arg = fmt.Sprintf("Return %s to %s", m.Arg, email)
    m.Err = ""
    enc := json.NewEncoder(w)
    enc.Encode(m)
}
