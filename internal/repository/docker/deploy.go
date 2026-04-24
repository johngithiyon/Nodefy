package docker

import (
	"log"
	"os"
	"os/exec"

	"github.com/johngithiyon/Nodefy/internal/models"
)


func Deploydocker(username string,Deploy models.Deploy) error {

	NGROK_AUTHTOKEN := os.Getenv("AUTH_TOKEN")
	 
	   packages := ""
	    
		if len(Deploy.Languages) > 0 {
				
			for i:=0;i<len(Deploy.Languages);i++ {
				packages += Deploy.Languages[i] + " "
			}
		} 

		if len(Deploy.Services) > 0 {
		for i:=0;i<len(Deploy.Services);i++ {

			if Deploy.Services[i] == "" {
				break
			}
			
			servicecmd := exec.Command(
				"docker", "run", "-d",
				"--name",Deploy.Appname+"-"+Deploy.Services[i],

				//this is a container config you can change based on your server capacity

				"--memory=512m",
				"--cpus=0.5",
				"--pids-limit=100",
				"--label", "owner="+username,
				"--label", "instance="+Deploy.Appname+"-"+Deploy.Services[i],
				"-e", "POSTGRES_HOST_AUTH_METHOD=trust",
				 Deploy.Services[i],
			)

			output,outputerr := servicecmd.CombinedOutput()
	
			if outputerr != nil {
                    log.Println("Failed in the services containers",outputerr,string(output))
					return  outputerr
			}
   
		   log.Println(string(output))
		}
		}

		//build the docker image 

		cmd := exec.Command(
		"docker", "build",
		"-t",Deploy.Appname+"-"+username, // image name

		"--build-arg", "PACKAGES="+packages,
		"--build-arg", "USERNAME="+username,

		".", 
		)

		output,outputerr := cmd.CombinedOutput()

		if outputerr != nil {

		log.Println("Deploy Build Err",outputerr,string(output))
		return  outputerr

		}

		log.Println(string(output))

		runcmd := exec.Command("docker","run","-d","--name",Deploy.Appname+"-"+username,"-e","NGROK_AUTHTOKEN="+NGROK_AUTHTOKEN,"-e","REPO_URL="+Deploy.Gitrepo,
		 
		  "--label", "owner="+username,
		 "--label", "instance="+Deploy.Appname+"-"+username,
		   Deploy.Appname+"-"+username,
	       
	      )

		runout,runerr := runcmd.CombinedOutput()

		if runerr != nil {
             log.Println("Run Err from deploy container",runerr,string(runout))
			 return  runerr 
		}

		log.Println(string(runout))

		return  nil 
	       
}