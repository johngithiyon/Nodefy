package postgres

import (
	"log"

	"github.com/johngithiyon/Nodefy/internal/models"
	
)

func SaveDeployinstances(username string, deploy models.Deploy) error {


	var appID int
	insertQuery := `
	INSERT INTO deploy_instances (appname, username)
	VALUES ($1, $2)
	RETURNING id
	`

	err := Database.Db.QueryRow(insertQuery, deploy.Appname, username).Scan(&appID)
	if err != nil {
		return err
	}


	serviceQuery := `
	INSERT INTO deploy_instance_services (app_id, services_name, status)
	VALUES ($1, $2, $3)
	`

	for _, service := range deploy.Services {
		_, err := Database.Db.Exec(serviceQuery, appID, service, "Running")
		if err != nil {
			log.Println("Error inserting service:", service, err)
			return err
		}
	}

	return nil
}

func GetDeployinstances(username string) ([]map[string]interface{}, error) {

	var result []map[string]interface{}

	appnamequery := "select id, appname from deploy_instances where username=$1"

	approw, appnamerr := Database.Db.Query(appnamequery, username)
	if appnamerr != nil {
		log.Println("Appname And Id err from deploy_instances", appnamerr)
		return nil, appnamerr
	}
	defer approw.Close()

	for approw.Next() {
		var appid int
		var appname string

		approw.Scan(&appid, &appname)

		servicesquery := "select services_name, status from deploy_instance_services where app_id=$1"

		servicesrow, serviceserr := Database.Db.Query(servicesquery, appid)
		if serviceserr != nil {
			log.Println("Get Err in Services", serviceserr)
			return nil, serviceserr
		}

		var services []map[string]string

		for servicesrow.Next() {
			var servicename string
			var servicestatus string

			servicesrow.Scan(&servicename, &servicestatus)

			services = append(services, map[string]string{
				"name":   servicename,
				"status": servicestatus,
			})
		}
		servicesrow.Close()

		result = append(result, map[string]interface{}{
			"appname":  appname,
			"services": services,
		})
	}

	return result, nil
}


func SaveDeployAddinstances(username string, deploy models.Deployaddinstances) error {

	var appID int
	query := `
	SELECT id FROM deploy_instances
	WHERE appname = $1 AND username = $2
	`

	err := Database.Db.QueryRow(query, deploy.Appname, username).Scan(&appID)
	if err != nil {
		return err
	}

	insertQuery := `
	INSERT INTO deploy_instance_services (app_id, services_name, status)
	VALUES ($1, $2, $3)
	`

	_, err = Database.Db.Exec(insertQuery, appID, deploy.Appname+deploy.Imagename, "Running")
	if err != nil {
		return err
	}

	return nil
}

func GetAppId(deploy models.Containermanage) (int,error) {

	   var appid int 
   
	   query := "select app_id from deploy_instance_services where services_name=$1"

	  appiderr :=  Database.Db.QueryRow(query,deploy.Instancename).Scan(&appid)
       
      if appiderr != nil {
		  return 0,appiderr
	  }

	  return appid,nil 
}