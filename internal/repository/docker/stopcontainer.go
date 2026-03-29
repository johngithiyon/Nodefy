package docker

import (
	"log"
	"os/exec"

	"github.com/johngithiyon/Nodefy/internal/models"
)

func Stopcontainer(username string ,Stopcontainer models.Containermanage) error {

	cmd :=  exec.Command(
             "docker",
			"ps",
			"-aq",
			"--filter", "label=owner="+username,
			"--filter","label=instance="+Stopcontainer.Instancename,
	)

	output,outputerr :=  cmd.CombinedOutput()

	if outputerr != nil {
		log.Println("Docker Ps  Error in stop  container ",outputerr,string(output))
		return outputerr
	}

	
	stopcmd := exec.Command("docker","stop",string(output))

	stopoutput,stoperr := stopcmd.CombinedOutput()

	if stoperr != nil {
		log.Println("Stop Error",stoperr,string(stopoutput))
		return stoperr
	}

	return  nil 
}