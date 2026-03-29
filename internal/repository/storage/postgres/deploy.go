package postgres

import (
	"log"
   pq "github.com/lib/pq"
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

func Getdeployinstance(username string) ([]models.Deployinfo,error) {

     var Deployinfo models.Deployinfo
	 var deployinfo []models.Deployinfo 

     getquery := "select appname,services from deploy_instances where username=$1" 

	 getrows,geterr :=  Database.Db.Query(getquery,username)

	 if geterr != nil {
		log.Println("Get Err in deploy instances",geterr)
		return nil,geterr
	 }

	 for getrows.Next() {
		 
		scanerr := getrows.Scan(&Deployinfo.Appname,pq.Array(&Deployinfo.Services))
		if scanerr != nil {
            log.Println("Scan Err in deploy instances",scanerr)
			return nil,scanerr
		}

		deployinfo = append(deployinfo, Deployinfo)
	 }

	 return deployinfo,nil 
}

func SaveDeployAddinstances(username string,Deployaddinstances models.Deployaddinstances) error {
	 
	    query := `UPDATE deploy_instances
		SET services = array_append(services,$1)
		WHERE appname = $2 and username = $3; `

		_,saverr := Database.Db.Exec(query,Deployaddinstances.Imagename,Deployaddinstances.Appname,username)

		if saverr != nil {
			log.Println("Save Err in the deploy_instances table for add instances",saverr)
			return  saverr
		}
	
        return  nil 

}
