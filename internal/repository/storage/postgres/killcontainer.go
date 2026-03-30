package postgres

import (
	"log"

	"github.com/johngithiyon/Nodefy/internal/models"
)

func Containerkill(username string, Containerkill models.Containermanage) error {

	if Containerkill.Type == "build" {

		userid, getiderr := Getuserid(username)
		if getiderr != nil {
			log.Println("Get err in kill container",getiderr)
			return getiderr
		}

		query := `
		DELETE FROM instances
		WHERE user_id = $1
		AND instance_name = $2;
		`

		_, deleteErr := Database.Db.Exec(query, userid, Containerkill.Instancename)
		if deleteErr != nil {
			log.Println("Delete Err in kill container",deleteErr)
			return deleteErr
		}
	}

	if Containerkill.Type == "deploy" {

	   appid,getiderr := GetAppId(Containerkill)

	   if getiderr != nil {
		  return getiderr
	   }

		query := `
		DELETE FROM deploy_instance_services
		WHERE app_id = $1
		AND services_name = $2;
		`

		_, deleteErr := Database.Db.Exec(query,appid, Containerkill.Instancename)
		if deleteErr != nil {
			log.Println("Error deleting in deploy_instances_services:", deleteErr)
			return deleteErr
		}
	}

	return nil
}