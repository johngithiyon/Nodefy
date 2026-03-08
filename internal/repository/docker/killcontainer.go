package docker

import (
	"log"
	"os/exec"
)

func Killcontainer(username string) error {

	     cmd := exec.Command("docker","rm","-f",username)

		output,outputerr :=  cmd.CombinedOutput()

		if outputerr != nil {
			log.Println("Docker Kill Error",outputerr)
			return outputerr
		}

		log.Println("Docker Kill Output",string(output))

		 return nil 
}