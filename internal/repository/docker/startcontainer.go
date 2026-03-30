package docker

import (
	"log"
	"os/exec"
	"strings"

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

		id := strings.TrimSpace(string(output))

		startcmd := exec.Command("docker","start",id)

	    startoutput,starterr := startcmd.CombinedOutput()

		if starterr != nil {
			log.Println("Start Error",starterr,string(startoutput))
			return starterr
		}

	return nil
}
