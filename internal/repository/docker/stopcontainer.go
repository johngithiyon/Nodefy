package docker

import (
	"log"
	"os/exec"
	"strings"

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

	id := strings.TrimSpace(string(output))


	log.Println("Docker ps output",string(output),id)

	
	stopcmd := exec.Command("docker","stop",id)

	stopoutput,stoperr := stopcmd.CombinedOutput()

	if stoperr != nil {
		log.Println("Stop Error",stoperr,string(stopoutput))
		return stoperr
	}

	return  nil 
}