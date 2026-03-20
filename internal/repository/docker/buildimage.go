package docker

import (
	"log"
	"os/exec"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)



func BuildImage(username string,servicename string) error {
	     

			lowercase_username := utils.Lowercase(username)

            //pull and run the docker contiainer 

			cmd := exec.Command(
				"docker", "run", "-d",
				"--name",lowercase_username+"-"+servicename,
				"--label", "owner="+username,
				"--label", "instance="+lowercase_username+"-"+servicename,
				 servicename,
			)

			output,outputerr := cmd.CombinedOutput()
	
			if outputerr != nil {
                    log.Println("Failed in the services containers",outputerr)
					return  outputerr
			}
   
		   log.Println(string(output))

		  saverr :=  postgres.SaveInstance(lowercase_username+"-"+servicename,username)

		  if saverr != nil {
                return errors.ErrInternalserver
		  }

			return nil 

}