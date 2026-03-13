package linux

import (
	"log"
	"os/exec"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
)

func Setpasswd(username string) error {

	  passwd,searcherr := postgres.SearchPassword(username)
	  
	  if searcherr != nil {
		   log.Println("Search Error for password",searcherr)
		   return errors.ErrInternalserver
	  }

	  cmd := exec.Command("bash", "-c", "echo '"+username+":"+passwd+"' | sudo chpasswd -e")

	 output,outputerr :=  cmd.CombinedOutput()

	 if outputerr != nil {
		 log.Println("Err from the passwd execution",outputerr)
		 return errors.ErrInternalserver
	 }

	 log.Println("Output for the setpasswd",string(output))

	  return  nil 
}