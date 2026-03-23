package docker

import (
	"log"
	"os/exec"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
)



func BuildImage(username string,servicename string) error {
	    
	         var Build models.Build

            //pull and run the docker contiainer 

			cmd := exec.Command(
				"docker", "run", "-d",
				"--name",Build.Instancename+"-"+servicename,
				"--label", "owner="+username,
				"--label", "instance="+Build.Instancename+"-"+servicename,
				 servicename,
			)

			output,outputerr := cmd.CombinedOutput()
	
			if outputerr != nil {
                    log.Println("Failed in the services containers",outputerr)
					return  outputerr
			}
   
		   log.Println(string(output))

		  saverr :=  postgres.SaveInstance(Build.Instancename+"-"+servicename,username)

		  if saverr != nil {
                return errors.ErrInternalserver
		  }

			return nil 

}