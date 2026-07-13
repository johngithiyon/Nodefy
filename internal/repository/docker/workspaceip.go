package docker

import (
	"log"
	"os/exec"
)

//Function used to find the ip address of the workspace container

func Findipworkspace(username string) (string,error) {
        
	cmdoutput,runerr := exec.Command(
		"docker",
		"inspect",
		"-f",
		"{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}",
		"workspace-"+username,
	).CombinedOutput()

	if runerr != nil {
		  log.Println("Ip find err in the workspace",string(cmdoutput))
		  return  "",runerr
	}

	ip := string(cmdoutput)

	return ip,nil 	    
}