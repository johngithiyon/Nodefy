package postgres

import (
	"log"

	"github.com/johngithiyon/Nodefy/internal/models"
)

func Containerstatus(username string, Containerstatus models.Containermanage) error {

       if Containerstatus.Type == "build" {
   
		userid,getiderr := Getuserid(username)

		if getiderr != nil {
			return getiderr
		}
		
		query := `UPDATE instances
			SET status = $1
			WHERE user_id = $2
			AND instance_name = $3;`

        _,updaterr := Database.Db.Exec(query,Containerstatus.Status,userid,Containerstatus.Instancename)

		if updaterr != nil {
			return updaterr
		}

	   }

	   if Containerstatus.Type == "deploy" {

		query := `
		UPDATE deploy_instances_services
		SET status = $1
		WHERE app_id = $2 AND service_name = $3
		`
		
		_, upadeterr := Database.Db.Exec(query,Containerstatus.Status, 9,Containerstatus.Instancename)
		if upadeterr != nil {
			log.Println("Error updating status in deploy_instances_services:",upadeterr)
			return  upadeterr
		}
		         
	   }

	   return  nil 
	  
}