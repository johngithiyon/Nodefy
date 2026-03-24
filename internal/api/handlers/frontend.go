package handlers

import (
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/services"
)

func Rendersignupage(w http.ResponseWriter, r * http.Request) {
          services.Render(w,"signup.html")
}

func Renderloginpage(w http.ResponseWriter, r *http.Request) {
	     services.Render(w,"login.html")
}

func RenderOtpverifypage(w http.ResponseWriter, r *http.Request) {
	 services.Render(w,"otpverification.html")
}

func Renderhomepage(w http.ResponseWriter,r *http.Request) {
	services.Render(w,"index.html")
}

func Renderaboutpage(w http.ResponseWriter, r *http.Request) {
	services.Render(w,"about.html")
}

func RenderBuildpage(w http.ResponseWriter, r *http.Request) {
	services.Render(w,"build.html")
}

func RenderChoosepage(w http.ResponseWriter, r *http.Request) {
	services.Render(w,"choose.html")
}

