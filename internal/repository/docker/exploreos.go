package docker

import (
	"log"
	"os/exec"

	"github.com/johngithiyon/Nodefy/internal/models"
)

func Exploreos(username string,exploreos models.Exploreos) error {

   //Fix: Other linux name pull problm

	cmd := exec.Command(
		"docker", "run", "-d",
		"--name",exploreos.Instancename,
		"--label", "owner="+username,
		"--label", "instance="+exploreos.Instancename,
		exploreos.Osname,
		"sh", "-c", "sleep 1800 && kill 1", // kill the container after 30 minutes 
	)

	output,outputerr := cmd.CombinedOutput()

	if outputerr != nil {
		log.Println("Err in Explore Os",outputerr.Error())
		return outputerr
	}

	log.Println(string(output))

	return  nil 
}