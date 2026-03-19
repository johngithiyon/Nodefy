package docker

import (
	"log"
	"os/exec"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func BuildImage(username string,Deploy *models.Deploy) error {

	        var service string 
			
            for _,services := range Deploy.Services {
				 
				   service += services + " "
			}

			lowercase_username := utils.Lowercase(username)

			//Write  a commmand to build the image

			executer := exec.Command(
				"docker", "build",
				"--build-arg", "BASE_IMAGE="+Deploy.OsName,
				"--build-arg", "PACKAGES="+service,
				"-t", lowercase_username+"-"+Deploy.Instancename,
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

			name := lowercase_username + "-" + Deploy.Instancename

			runexecuter := exec.Command(
				"docker", "run",

				"--name", name,

				"-u", "1000",
				"-m", "100m",
				"--memory-swap=100m",

				"--network", "nodefy-network",

				"--label", "owner="+username,
				"--label", "instance="+Deploy.Instancename,

				"-d",

				
				name,

				"sleep",
				"infinity",
			)

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