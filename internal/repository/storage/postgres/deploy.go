package postgres

import (
	"log"

	"github.com/johngithiyon/Nodefy/internal/models"
)

func Savedeployinstance(username string , Deploy models.Deploy) error {

	query := `
		INSERT INTO deploy_instances (appname, languages, services, username)
		VALUES ($1, $2, $3, $4)
	`

	_,saverr := Database.Db.Exec(query,Deploy.Appname,Deploy.Languages,Deploy.Services,username)

	if saverr != nil {
		log.Println("Save Err in the deploy_instances table",saverr)
		return  saverr
	}

	return nil 
}