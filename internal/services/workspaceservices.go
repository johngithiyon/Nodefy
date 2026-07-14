package services

import (

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
)

func Workspaceservices(userhash string) (string,error){

    containerip,checkerr := postgres.CheckWorkspace(userhash)

	if checkerr != nil {
          return  "",errors.ErrInternalserver
	}

	if len(containerip) > 0 {
              return containerip,nil 
		    
	} 
		return "",nil 
  
}