package docker

import (
	"log"
	"os/exec"

	"github.com/johngithiyon/Nodefy/internal/models"
)

func BuildImage(Deploy *models.Deploy) error {

	        var service string 
	 
            for _,services := range Deploy.Services {
				 
				   service += services + " "
			}

			//Write  a commmand to build the image

			executer := exec.Command(
				"docker", "build",
				"--build-arg", "BASE_IMAGE="+Deploy.OsName,
				"--build-arg", "PACKAGES="+service,
				"-t", "nodefy-image",
				".",
			)

			//Used to show the output and error from the cmd 
		
			output, outerr := executer.CombinedOutput()
			if outerr != nil {
				log.Println("Build failed:", outerr)
				log.Println(string(output))
				return outerr
			}

		    log.Println(string(output))

			// write the command to run the container

			runexecuter := exec.Command("docker","run","-m","100m","--memory-swap=100m","-d","nodefy-image","sleep","infinity")

			runoutput,runerr := runexecuter.CombinedOutput()

			if runerr != nil {
				log.Println("Run failed:", runerr)
				log.Println(string(runoutput))
				return runerr
			}

			return nil 

}