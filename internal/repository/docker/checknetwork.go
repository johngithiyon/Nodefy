package docker

import (
	"log"
	"os/exec"
)


func Checknetwork(username string) error {

    cmd := exec.Command("docker" ,"network","inspect",username+"-"+"network")	 

	output,outputerr := cmd.CombinedOutput()

	if outputerr != nil {
		log.Println("Network Does Not Exists")
		 log.Println("Err in the network check",outputerr)
		 return outputerr
	}

	log.Println("Network Exists")

	log.Println(string(output))

	 return nil 
}