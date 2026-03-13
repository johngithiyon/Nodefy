package docker

import (
	"log"
	"os/exec"
)

func Stopcontainer(username string ) error {

	cmd :=  exec.Command("docker","stop",username)

	output,outputerr := cmd.CombinedOutput()

	if outputerr != nil {
		 log.Println("Docker Stop Error",outputerr)
		 return outputerr
	}

	log.Println("Docker Stop Output",string(output))

	return  nil 
}