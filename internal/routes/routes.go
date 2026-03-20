package routes

import (
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/api/handlers"
)

func Routes() {

	  // frontend endpoints 

	  http.HandleFunc("/signupage",handlers.Rendersignupage)
	  http.HandleFunc("/loginpage",handlers.Renderloginpage)
	  http.HandleFunc("/otpverificationpage",handlers.RenderOtpverifypage)
	  http.HandleFunc("/",handlers.Renderhomepage)
	  http.HandleFunc("/aboutpage",handlers.Renderaboutpage)

	  // Api endpoints
	 
	  http.HandleFunc("/buildinstance",handlers.Buildhandler)
	  http.HandleFunc("/signup",handlers.Signuphandler)
	  http.HandleFunc("/otp",handlers.Otphandler)
	  http.HandleFunc("/login",handlers.Loginhandler)
	  http.HandleFunc("/stopinstance",handlers.StopcontainerHandler)
	  http.HandleFunc("/killinstance",handlers.Killcontainerhandler)
	  http.HandleFunc("/getinstance",handlers.GetInstancesHandler)
	  http.HandleFunc("/workspace",handlers.WorkspaceHanlder)
}