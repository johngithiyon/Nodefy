package services

import (
	"log"
	"os/exec"

	"github.com/johngithiyon/Nodefy/internal/models"
)

func DeployInstances(Deploy *models.Deploy ) error  {


	cmd  := exec.Command("sudo" ,"docker", "run",Deploy.OsName)

	bytes,runerr := cmd.CombinedOutput()

	log.Println(string(bytes))

	if runerr != nil {
		return runerr 
	}
    
	return nil 

}