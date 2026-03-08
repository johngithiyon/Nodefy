package utils

import (
	"log"
	"net/http"
)

func Getsessionid(r *http.Request,cookiename string) (string,error)  {

	sessionid, cokkierr := r.Cookie(cookiename)
	if cokkierr != nil {
	    log.Println("cookie not found err in login")
		return "",cokkierr
	}

	return sessionid.Value,nil 
}