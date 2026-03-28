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

	saveinstance := "insert into instances(user_id, instance_name,status) Values($1,$2,$3)"

	_,saverr := Database.Db.Exec(saveinstance,userid,Instancename,"Running")

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

	getinstances := "select id,instance_name,status from instances where user_id=$1"

	rows,getinstanceserr := Database.Db.Query(getinstances,userid)

	if getinstanceserr != nil {
		return nil,errors.ErrInternalserver
	}

	for rows.Next() {
         rows.Scan(&instance.Instanceid,&instance.Instancename,&instance.Status)

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

func DeleteInstances(killinstancename string,username string , userid int) error {

	username = utils.Lowercase(username)

	deleteinstance := "delete from instances where instance_name=$1 and user_id=$2"

	res,delerr := Database.Db.Exec(deleteinstance,username+"-"+killinstancename,userid)

	if delerr != nil {
		log.Println("Del Err from kill instance",delerr)
		return delerr
	}

	_,rowsafferr := res.RowsAffected()


	if rowsafferr != nil {
		log.Println(rowsafferr)
		return rowsafferr
	}

	return nil 
}