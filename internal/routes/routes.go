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
	  http.HandleFunc("/buildinstancespage",handlers.RenderBuildpage)
	  http.HandleFunc("/choosepage",handlers.RenderChoosepage)
	  http.HandleFunc("/deploypage",handlers.RenderDeploypage)
	  http.HandleFunc("/explorepage",handlers.RenderExplorepage)
	  http.HandleFunc("/teachpage",handlers.RenderTeachpage)

	  // Api endpoints
	 
	  http.HandleFunc("/buildinstance",handlers.Buildhandler)
	  http.HandleFunc("/deployinstance",handlers.DeployHandler)
	  http.HandleFunc("/exploreos",handlers.ExploreosHandler)
	  http.HandleFunc("/signup",handlers.Signuphandler)
	  http.HandleFunc("/otp",handlers.Otphandler)
	  http.HandleFunc("/login",handlers.Loginhandler)
	  http.HandleFunc("/startinstance",handlers.Startcontainerhandler)
	  http.HandleFunc("/stopinstance",handlers.StopcontainerHandler)
	  http.HandleFunc("/killinstance",handlers.Killcontainerhandler)
	  http.HandleFunc("/getbuildinstance",handlers.GetBuildInstancesHandler)
      http.HandleFunc("/getdeployinstance",handlers.GetDeployinstance)
	  http.HandleFunc("/workspace",handlers.WorkspaceHanlder)
	  http.HandleFunc("/addinstancesbuild",handlers.Addinstanceshandler)
	  http.HandleFunc("/addinstancesdeploy",handlers.AddinstancesDeployHandler)
}