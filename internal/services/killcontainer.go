package services

import "github.com/johngithiyon/Nodefy/internal/repository/docker"

func Killcontainerservices(username string) error{

	  killerr :=  docker.Killcontainer(username)

	  if killerr != nil {
		  return killerr
	  }
     
	   return nil 
}