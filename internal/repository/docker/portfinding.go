package docker

import (
	"log"
	"os/exec"
	"strings"
)

func Portfind(instancename string) (string,error) {

	cmd := exec.Command("docker","port",instancename)

    output,porterr := cmd.CombinedOutput()

	if porterr != nil {
		log.Println("Port Scan Err",porterr)
		return  "",porterr
	}

	 parts := strings.Split(string(output),":")

	 port := strings.TrimSpace(parts[len(parts)-1])

	  return port,nil 
}