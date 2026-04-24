package docker

import (
	"log"
	"os/exec"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
)



func BuildImage(instancename string,username string,servicename string) error {  

            //pull and run the docker contiainer 

            cmd := exec.Command(
				"docker", "run", "-d",
				"--name", instancename+"-"+servicename,
				"--network", username+"-"+"network",

				//this is a container config you can change based on your server capacity

				"--memory=512m",
				"--cpus=0.5",
				"--pids-limit=100",


				"--label", "owner="+username,
				"--label", "instance="+instancename+"-"+servicename,
				"-e", "POSTGRES_HOST_AUTH_METHOD=trust", 
				servicename,
			)

			output,outputerr := cmd.CombinedOutput()
	
			if outputerr != nil {
                    log.Println("Failed in the services containers",outputerr,string(output))
					return  outputerr
			}
   
		  saverr :=  postgres.SaveInstance(instancename+"-"+servicename,username)

		  if saverr != nil {
                return errors.ErrInternalserver
		  }

			return nil 

}