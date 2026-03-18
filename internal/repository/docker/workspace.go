package docker

import (
	"log"
	"os/exec"
)

func Startworkspace(username string,instancename string ) (string,error) {

	log.Println(instancename)

	cmd := exec.Command(
		"docker", "exec", "-d",
		instancename,
		"code-server",
		"--bind-addr", "0.0.0.0:8081",
		"--auth", "none",
		"/app",
	)

     output,runerr := cmd.CombinedOutput()

	 log.Println(string(output))

	if runerr != nil {
         log.Println("Run Err from Workspace",runerr)
		 return "",runerr
	}

	url := "http://localhost:81"+"/"+instancename+".nodefy"

	return url,nil 
}