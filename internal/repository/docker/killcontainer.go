package docker

import (
	"log"
	"os/exec"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
)

func Killcontainer(killinstance models.Containermanage,username string) error {

	     cmd := exec.Command(
			"docker",
			"ps",
			"-aq",
			"--filter", "label=owner="+username,
			 "--filter","label=instance="+killinstance.Instancename,
		)

		output,outputerr :=  cmd.CombinedOutput()

		if outputerr != nil {
			log.Println("Docker Kill Error",outputerr)
			return errors.ErrInternalserver
		}

		log.Println(string(output))
   
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