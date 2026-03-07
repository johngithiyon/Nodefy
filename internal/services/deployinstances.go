package services

import (
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/docker"
)

func DeployInstances(Deploy *models.Deploy ) error  {

   builerr :=  docker.BuildImage(Deploy)

   if builerr != nil {
	     
	     return builerr
   }

	return nil 

}