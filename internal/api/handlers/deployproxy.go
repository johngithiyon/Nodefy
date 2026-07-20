package handlers

import (
	"log"
	"net/http"

	"net/http/httputil"
	"net/url"

	"github.com/johngithiyon/Nodefy/internal/services"
	"github.com/johngithiyon/Nodefy/pkg/response"
)

func Deployproxy(w http.ResponseWriter, r *http.Request) {
     
	 containerip,_,err:= services.Deployproxy(r.Host)

	 if err != nil {
	       response.Response(w,500,"Internal Server Error")
		   return	 
	 }
     
	target,parserr := url.Parse("http://"+containerip+":"+"8090")

	if parserr != nil {
		log.Println("Parse Err",parserr)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.ServeHTTP(w,r)

	    
}