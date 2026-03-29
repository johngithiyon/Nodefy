package docker

import (
	"log"
	"os/exec"

	"github.com/johngithiyon/Nodefy/internal/models"
)

func DeployAddinstances(username string,DeployAddinstances models.Deployaddinstances) error {

	// start the container by the user given image name 
	
	  cmd :=  exec.Command("docker","run","--name",DeployAddinstances.Appname,"-d",DeployAddinstances.Imagename)

      output,outputerr :=  cmd.CombinedOutput()

	  if outputerr != nil {
		  log.Println("Outputerr in the Add Instances",outputerr,string(output))
		  return  outputerr
	  }

	   return  nil 
}