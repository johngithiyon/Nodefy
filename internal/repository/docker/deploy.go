package docker

import (
	"log"
	"os"
	"os/exec"

	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)


func Deploydocker(username string,Deploy models.Deploy) error {

	lowercase_username := utils.Lowercase(username)

	 AUTH_TOKEN := os.Getenv("AUTH_TOKEN")
	 
	   packages := ""
	    
		if len(Deploy.Languages) > 0 {
				
			for i:=0;i<len(Deploy.Languages);i++ {
				packages += Deploy.Languages[i] + " "
			}
		} 

		if len(Deploy.Services) > 0 {
		for i:=0;i<len(Deploy.Services);i++ {
			
			servicecmd := exec.Command(
				"docker", "run", "-d",
				"--name",lowercase_username+"-"+Deploy.Services[i],
				"--label", "owner="+lowercase_username,
				"--label", "instance="+lowercase_username+"-"+Deploy.Services[i],
				 Deploy.Services[i],
			)

			output,outputerr := servicecmd.CombinedOutput()
	
			if outputerr != nil {
                    log.Println("Failed in the services containers",outputerr)
					return  outputerr
			}
   
		   log.Println(string(output))
		}
		}

		//build the docker image 

		cmd := exec.Command(
		"docker", "build",
		"-t",lowercase_username+"-"+Deploy.Appname, // image name

		"--build-arg", "PACKAGES="+packages,
		"--build-arg", "USERNAME="+lowercase_username,

		".", 
		)

		output,outputerr := cmd.CombinedOutput()

		if outputerr != nil {

		log.Println("Deploy Build Err",outputerr)
		return  outputerr

		}

		log.Println(string(output))

		runcmd := exec.Command("docker","run","-d","--name",lowercase_username+"-"+Deploy.Appname,"-e","AUTH_TOKEN="+AUTH_TOKEN,"-e","REPO_URL="+Deploy.Gitrepo,
		 
		  "--label", "owner="+lowercase_username,
		 "--label", "instance="+lowercase_username+"-"+Deploy.Appname,
		 lowercase_username+"-"+Deploy.Appname,
	       
	      )

		runout,runerr := runcmd.CombinedOutput()

		if runerr != nil {
             log.Println("Run Err from deploy container",runerr)
			 return  runerr 
		}

		log.Println(string(runout))

		return  nil 
	       
}