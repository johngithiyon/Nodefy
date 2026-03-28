package postgres

import (
	"log"

	"github.com/johngithiyon/Nodefy/internal/models"
)

func  Stopcontainer(username string,stopconatiner models.Stopcontainer) error {
 
		query := `
		UPDATE instances
		SET status = 'Stopped'
		WHERE instance_name = $1 and username=$2
	`

	_,updaterr := Database.Db.Exec(query,stopconatiner.Instancename,username)
	if updaterr != nil {
		log.Println("Update Err in Stop Container",updaterr)
		return updaterr
	}

   return nil 
}