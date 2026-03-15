package linux

import (
	"log"
	"os/exec"
)

func CheckUserlinux(username string) error {
	 
	cmd := exec.Command("id",username)

	output,cmderr := cmd.CombinedOutput()

	if cmderr != nil {
          return cmderr
	}

	log.Println(string(output))

	return  nil 
}