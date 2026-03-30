package docker

import (
	"log"
	"os/exec"

	"github.com/johngithiyon/Nodefy/internal/models"
)

func Addinstances(username string,Addinstances models.Addinstancesbuild) error {

	// start the container by the user given image name 
	
	cmd := exec.Command(
		"docker", "run",
		"--name", Addinstances.Instancename,
		"--label", "owner="+username,
		"--label", "instance="+Addinstances.Instancename,
		"-d",
		Addinstances.Imagename,
		"sleep", "infinity",
	)
	 
      output,outputerr :=  cmd.CombinedOutput()

	  if outputerr != nil {
		  log.Println("Outputerr in the Add Instances",outputerr,string(output))
		  return  outputerr
	  }

	  log.Println(string(output))

	   return  nil 
}