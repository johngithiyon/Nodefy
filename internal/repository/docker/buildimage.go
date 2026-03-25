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
				"--name",instancename+"-"+servicename,
				"--label", "owner="+username,
				"--label", "instance="+instancename+"-"+servicename,
				 servicename,
			)

			output,outputerr := cmd.CombinedOutput()
	
			if outputerr != nil {
                    log.Println("Failed in the services containers",outputerr)
					return  outputerr
			}
   
		   log.Println(string(output))

		  saverr :=  postgres.SaveInstance(instancename+"-"+servicename,username)

		  if saverr != nil {
                return errors.ErrInternalserver
		  }

			return nil 

}