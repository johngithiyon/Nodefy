package linux

import (
	"log"
	"os/exec"

	"github.com/johngithiyon/Nodefy/internal/errors"
)

func Adduser(username string ) error {

    cmd :=  exec.Command("sudo","useradd",username)

	output,outputerr := cmd.CombinedOutput()

	if outputerr != nil {
		 log.Println("Output Err in the useradd command",outputerr)
		 return errors.ErrInternalserver
	}

	log.Println(string(output))
 
   passwdseterr := 	Setpasswd(username)

   if passwdseterr != nil {
	     return errors.ErrInternalserver
   }
	 
	return nil 
}