package docker

import (
	"log"
	"os/exec"
)

func Startworkspace(username string) (string,error) {

   containername := "workspace"+"-"+username

    cmd := exec.Command(
		"docker", "run", "-d",
		"--name", containername,
		"--network", username+"-network",
		"codercom/code-server:latest",
		"--auth", "none",
	)

     output,runerr := cmd.CombinedOutput()

	 log.Println(string(output))

	if runerr != nil {
         log.Println("Run Err from Workspace",runerr)
		 return "",runerr
	}

	return containername,nil 
}