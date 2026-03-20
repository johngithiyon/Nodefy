package docker

import (
	"log"
	"os/exec"
)

func Createnetwork(username string) error {

	     cmd := exec.Command("docker","network","create",username+"-"+"network")

		 output,outputerr := cmd.CombinedOutput()

		 if outputerr != nil {

			log.Println("Failed To Create Network For",username,outputerr)
			return  outputerr
		 }

		 log.Println(string(output))

		 return nil 
}

func Deletenetwork(username string) error {
	 
	   cmd := exec.Command("docker","network","rm",username+"-network")

	   output,outputerr :=  cmd.CombinedOutput()

	   if outputerr != nil {
		    log.Println("Delete Network",outputerr)
			return outputerr
	   }

	   log.Println(string(output))

	   return nil 
}
