package routes

import (
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/api/handlers"
)

func Routes() {

	 //serve the static files

	 http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	 //serve the profile images

	 http.Handle("/profileimage/", http.StripPrefix("/profileimage/", http.FileServer(http.Dir("./profileimage"))))

	  // frontend endpoints 

	  http.HandleFunc("/signupage",handlers.Rendersignupage)
	  http.HandleFunc("/loginpage",handlers.Renderloginpage)
	  http.HandleFunc("/otpverificationpage",handlers.RenderOtpverifypage)
	  http.HandleFunc("/",handlers.Roothandler)
	  http.HandleFunc("/aboutpage",handlers.Renderaboutpage)
	  http.HandleFunc("/buildpage",handlers.RenderBuildpage)
	  http.HandleFunc("/choosepage",handlers.RenderChoosepage)
	  http.HandleFunc("/deploypage",handlers.RenderDeploypage)
	  http.HandleFunc("/exploreospage",handlers.RenderExplorepage)
	  http.HandleFunc("/teachpage",handlers.RenderTeachpage)
	  http.HandleFunc("/buildinfopage",handlers.RenderBuildinstancepage)
	  http.HandleFunc("/deployinfopage",handlers.RenderDeployinstancepage)
	  http.HandleFunc("/exposeappage",handlers.Exposeapphandler)

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
	  http.HandleFunc("/addinstancesbuild",handlers.Addinstanceshandler)
	  http.HandleFunc("/addinstancesdeploy",handlers.AddinstancesDeployHandler)
	  http.HandleFunc("/createworkspace",handlers.Createworkspace)
}