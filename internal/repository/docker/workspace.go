package docker

import (
	"log"
	"os/exec"
)

func Startworkspace(username string,instancename string ) (string,error) {

   //TO Do : set the proxy and route the users to the containers

    cmd := exec.Command(
		"docker", "run", "-d",
		"--name", "workspace"+"-"+instancename,
		"--network", "nodefy-network",
		"-p", "0:8080",
		"codercom/code-server:latest",
		"--auth", "none",
	)

     output,runerr := cmd.CombinedOutput()

	 log.Println(string(output))

	if runerr != nil {
         log.Println("Run Err from Workspace",runerr)
		 return "",runerr
	}

	ports,porterr := Portfind("workspace-"+instancename)

	if porterr != nil {
		return "",porterr
	}

	url := "http://localhost:"+ports


	return url,nil 
}