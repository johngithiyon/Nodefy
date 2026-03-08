package services

import "github.com/johngithiyon/Nodefy/internal/repository/docker"

func Stopcontainerservices(username string) error{

	  dockerstoperr :=   docker.Stopcontainer(username)

	  if dockerstoperr != nil {
		    return dockerstoperr
	  }

	  return  nil 
}