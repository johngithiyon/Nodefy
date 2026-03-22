package docker

import (
	"log"
	"os/exec"

	"github.com/johngithiyon/Nodefy/internal/models"
)

func Exploreos(username string,exploreos models.Exploreos) error {

	cmd := exec.Command(
		"docker", "run", "-d",
		"--name", username+"-"+exploreos.Osname,
		exploreos.Osname,
		"sh", "-c", "sleep 1800 && kill 1", // kill the container after 30 minutes 
	)

	output,outputerr := cmd.CombinedOutput()

	if outputerr != nil {
		log.Println("Err in Explore Os",outputerr)
		return outputerr
	}

	log.Println(string(output))

	return  nil 
}