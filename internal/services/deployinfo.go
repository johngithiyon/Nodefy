package services

import (
	"log"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func Getdeployinfo(sessionid string) (string,error)  {
 
	username,getusernamerr := redis.Getusername(sessionid)
 
	if getusernamerr != nil {
		return "",errors.ErrInternalserver
	}

	deployinstances,geterr := postgres.GetDeployinstances(username)

	if geterr != nil {
		 log.Println("postgres geterr",geterr)
		 return "",errors.ErrInternalserver
	}

	resp,converr := utils.Jsonconvertor(deployinstances)

	if converr != nil {
		log.Println("jsonconverr",converr)
		return "",errors.ErrInternalserver
	}

	return resp,nil 
}