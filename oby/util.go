package oby

import (
    "appengine"
    "appengine/user"
)

func getEmail(c appengine.Context) string {
    u := user.Current(c)
    if u != nil {
        return u.Email
    }
    return ""
}
