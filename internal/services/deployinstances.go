package services

import (
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/docker"
)

func DeployInstances(Deploy *models.Deploy ) error  {

    docker.BuildImage(Deploy)
	return nil 

}