package postgres

import (
	"log"


	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func SaveInstance(Instancename string,username string) error {

	userid,geterr := Getuserid(username)

	if geterr != nil {
		return errors.ErrInternalserver
	}

	lowercase_username := utils.Lowercase(username)

	saveinstance := "insert into instances(user_id, instance_name) Values($1,$2)"

	_,saverr := Database.Db.Exec(saveinstance,userid,lowercase_username+"-"+Instancename)

	if saverr != nil {
		log.Println("Save Instances Error",saverr)
		return errors.ErrInternalserver
	}
          
	  return nil 
}

func GetInstances(username string) ([]models.Instances,error) {
	var instances[] models.Instances
    var instance models.Instances

	userid,getiderr := Getuserid(username)

	if userid == 0 && getiderr != nil {
		return nil,errors.ErrBadrequest
	}

	getinstances := "select id,instance_name from instances where user_id=$1"

	rows,getinstanceserr := Database.Db.Query(getinstances,userid)

	if getinstanceserr != nil {
		return nil,errors.ErrInternalserver
	}

	for rows.Next() {
         rows.Scan(&instance.Instanceid,&instance.Instancename)

		 instances = append(instances,instance)

	}


	return instances,nil 
}

func Getinstancebyid(ids string) (string,error){

	var instancename string 

	id,converr := utils.StrtoInt(ids)

	if converr != nil {
		log.Println("Convert string into int",converr)
		return "",converr
	}

	getinstaces := "select instance_name from instances where id=$1"

	instancegeterr := Database.Db.QueryRow(getinstaces,id).Scan(&instancename)

	if instancegeterr != nil {
		 log.Println("Instance name get err ",instancegeterr)
		 return "",instancegeterr
	}

	log.Println("This is instancename",instancename)


	return instancename,nil 
}