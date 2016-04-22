package main

import (
	"io"
	"net/http"
	"strconv"
)

func serve_the_webpage(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("user-cookie")
	//no user-cookie exist, create one
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Nm:  "user-cookie",
			Val: "0",
		}
	}
	//update existing cookie
	i, _ := strconv.Atoi(cookie.Val)
	i++
	cookie.Val = strconv.Itoa(i)

	http.SetCookie(res, cookie)

	io.WriteString(res, cookie.Val)
}

func throw_the_icon(res http.ResponseWriter, req *http.Request) {
	//do nothing when favicon.ico is requested.
	//does not increment the user-cookie
}

func main() {
	http.HandleFunc("/", serve_the_webpage)
	http.HandleFunc("/favicon.ico", throw_the_icon)
	http.ListenAndServe(":8080", nil)
}
