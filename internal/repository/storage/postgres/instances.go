package postgres

import (
	"log"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
)

func SaveInstance(Instancename string,username string) error {

	userid,geterr := Getuserid(username)

	if geterr != nil {
		return errors.ErrInternalserver
	}

	saveinstance := "insert into instances(user_id, instance_name) Values($1,$2)"

	_,saverr := Database.Db.Exec(saveinstance,userid,Instancename)

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