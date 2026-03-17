package services

import (
	"log"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func Getinstanceservices(sessionid string ) (string,error){

	username,getusernamerr := redis.Getusername(sessionid)
 
	if getusernamerr != nil {
		return "",errors.ErrInternalserver
	}

	instances,getinstancerr := postgres.GetInstances(username)

	if getinstancerr != nil {
		log.Println("Get Instances Err",getinstancerr)
		return "",errors.ErrInternalserver
	}

	respinstances,converr := utils.Jsonconvertor(instances)

	if converr != nil {
		return "",errors.ErrInternalserver																											
	}
	
	return  respinstances,nil 
}