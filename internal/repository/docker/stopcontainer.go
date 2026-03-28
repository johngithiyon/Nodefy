package docker

import (
	"log"
	"os/exec"

	"github.com/johngithiyon/Nodefy/internal/models"
)

func Stopcontainer(Stopcontainer models.Stopcontainer) error {

	cmd :=  exec.Command("docker","stop",Stopcontainer.Instancename)

	output,outputerr := cmd.CombinedOutput()

	if outputerr != nil {
		 log.Println("Docker Stop Error",outputerr)
		 return outputerr
	}

	log.Println("Docker Stop Output",string(output))

	return  nil 
}