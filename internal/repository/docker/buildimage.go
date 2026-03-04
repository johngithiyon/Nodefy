package docker

import (
	"log"
	"os/exec"

	"github.com/johngithiyon/Nodefy/internal/models"
)

func BuildImage(Deploy *models.Deploy) {

	        var service string 
	 
            for _,services := range Deploy.Services {
				 
				   service += services + " "
			}


			executer := exec.Command("sudo",
				"docker", "build",
				"--build-arg", "BASE_IMAGE="+Deploy.OsName,
				"--build-arg", "PACKAGES="+service,
				"-t", "nodefy-image",
				".",
			)
		
			output, err := executer.CombinedOutput()
			if err != nil {
				log.Println("Build failed:", err)
				log.Println(string(output))
				return
			}

		    log.Println(string(output))

			runexecuter := exec.Command("sudo","docker","run","-d","nodefy-image","sleep","infinity")

			runoutput,runerr := runexecuter.CombinedOutput()

			if runerr != nil {
				log.Println("Run failed:", runerr)
				log.Println(string(runoutput))
				return
			}

}