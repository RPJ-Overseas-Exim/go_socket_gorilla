package cookie

import (
	"net/http"
	"time"
)

func CreateCookie(name, value string, expire time.Time) *http.Cookie {
    cookie := new(http.Cookie)

    cookie.Name = name
    cookie.Value = value
    cookie.Expires = expire

    return cookie
}

func DeleteCookie(name, value string) *http.Cookie {
    return CreateCookie(name, value, time.Now())
}
