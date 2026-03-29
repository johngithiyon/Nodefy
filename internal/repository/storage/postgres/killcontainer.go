package postgres

import (
	"log"

	"github.com/johngithiyon/Nodefy/internal/models"
)

func Containerkill(username string, Containerkill models.Containermanage) error {

	if Containerkill.Type == "build" {

		userid, getiderr := Getuserid(username)
		if getiderr != nil {
			return getiderr
		}

		query := `
		DELETE FROM instances
		WHERE user_id = $1
		AND instance_name = $2;
		`

		_, deleteErr := Database.Db.Exec(query, userid, Containerkill.Instancename)
		if deleteErr != nil {
			return deleteErr
		}
	}

	if Containerkill.Type == "deploy" {

	   appid,getiderr := GetAppId(Containerkill)

	   if getiderr != nil {
		  return getiderr
	   }

		query := `
		DELETE FROM deploy_instances_services
		WHERE app_id = $1
		AND service_name = $2;
		`

		_, deleteErr := Database.Db.Exec(query,appid, Containerkill.Instancename)
		if deleteErr != nil {
			log.Println("Error deleting in deploy_instances_services:", deleteErr)
			return deleteErr
		}
	}

	return nil
}