package docker

import (
	"log"
	"os/exec"

	"github.com/johngithiyon/Nodefy/internal/models"
)

func Startcontainer(username string,startcontainer models.Containermanage) error {

		cmd := exec.Command(
			"docker",
			"ps",
			"-aq",
			"--filter", "label=owner="+username,
			"--filter","label=instance="+startcontainer.Instancename,
		)

	   output,outputerr :=  cmd.CombinedOutput()

		if outputerr != nil {
			log.Println("Docker Ps  Error in start container ",outputerr,string(output))
			return outputerr
		}

		startcmd := exec.Command("docker","start",string(output))

	    startoutput,starterr := startcmd.CombinedOutput()

		if starterr != nil {
			log.Println("Start Error",starterr,string(startoutput))
			return starterr
		}

	return nil
}
