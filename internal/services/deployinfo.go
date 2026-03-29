package services

import (
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
		 return "",errors.ErrInternalserver
	}

	resp,converr := utils.Jsonconvertor(deployinstances)

	if converr != nil {
		return "",errors.ErrInternalserver
	}

	return resp,nil 
}