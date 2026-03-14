package docker

import (
	"log"
	"os/exec"
	"strings"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
)

func BuildImage(username string,Deploy *models.Deploy) error {

	        var service string 
			
            for _,services := range Deploy.Services {
				 
				   service += services + " "
			}

			lowercase_username := strings.ToLower(username)

			//Write  a commmand to build the image

			executer := exec.Command(
				"docker", "build",
				"--build-arg", "BASE_IMAGE="+Deploy.OsName,
				"--build-arg", "PACKAGES="+service,
				"-t", lowercase_username+Deploy.OsName,
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

			runexecuter := exec.Command("docker","run","-u","1000","-m","100m","--memory-swap=100m","--label","owner="+username,"--label","instance="+Deploy.Instancename,"-d",lowercase_username+Deploy.OsName,"sleep","infinity")

			runoutput,runerr := runexecuter.CombinedOutput()

			if runerr != nil {
				log.Println("Run failed:", runerr)
				log.Println(string(runoutput))
				return runerr
			}

		//save the instance name in db

		  saverr :=  postgres.SaveInstance(Deploy.Instancename,username)

		  if saverr != nil {
                return errors.ErrInternalserver
		  }

			return nil 

}