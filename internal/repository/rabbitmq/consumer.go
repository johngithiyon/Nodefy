package rabbitmq

import (
	"encoding/json"


	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/docker"
)


func Consumer(data []byte,consumertype string) error {

	
	 if consumertype == "exploreos" {
		 
	        var exploreos models.Exploreos

			json.Unmarshal(data,&exploreos)

			explorerr := docker.Exploreos(&exploreos)

			if explorerr != nil {
				return explorerr
			}
	 }

	   if consumertype == "buildinstance" {
		 
		       var Build models.Build

			   json.Unmarshal(data,&Build)

				services := Build.Services

            for i:=0;i<len(services);i++ {

	         builderr :=   docker.BuildImage(Build.Instancename,Build.Username,services[i])

			if builderr != nil {
				return errors.ErrInternalserver
			}
       }

	 }

	 if consumertype == "deployinstance" {
		    
		    var deploy models.Deploy

			json.Unmarshal(data,&deploy)

			deployerr := docker.Deploydocker(deploy.Username,&deploy)

			if deployerr != nil {
				return deployerr 
			}

	 }


	 return nil
}