package docker

import (
	"log"
	"os/exec"

	"github.com/johngithiyon/Nodefy/internal/errors"
)

func Killcontainer(killinstancename,username string) error {

	     cmd := exec.Command(
			"docker",
			"ps",
			"-aq",
			"--filter", "label=owner="+username,
			"--filter", "label=instance="+killinstancename,
		)

		output,outputerr :=  cmd.CombinedOutput()

		if outputerr != nil {
			log.Println("Docker Kill Error",outputerr)
			return errors.ErrInternalserver
		}
   
		log.Println("Docker Kill Id Get output",string(output))

		rmcmd := exec.Command("docker","rm","-f",string(output))

		rmoutput,rmerr := rmcmd.CombinedOutput()

		if rmerr != nil {
			log.Println("Remove Error",rmerr)
			return errors.ErrInternalserver
		}

		log.Println("Docker Kill Output",string(rmoutput))

		 return nil 
}