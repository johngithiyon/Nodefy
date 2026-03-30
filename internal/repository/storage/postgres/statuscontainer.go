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

		appid,getiderr := GetAppId(Containerstatus)

		if getiderr != nil {
			log.Println("Get Id err in deploy",getiderr)
			return getiderr
		}

		query := `
		UPDATE deploy_instance_services
		SET status = $1
		WHERE app_id = $2 AND services_name = $3
		`
		
		_, upadeterr := Database.Db.Exec(query,Containerstatus.Status,appid,Containerstatus.Instancename)
		if upadeterr != nil {
			log.Println("Error updating status in deploy_instances_services:",upadeterr)
			return  upadeterr
		}
		         
	   }

	   return  nil 
	  
}