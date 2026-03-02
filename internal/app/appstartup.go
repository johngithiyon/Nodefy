package app

import (
	"log"

	"github.com/johngithiyon/Nodefy/internal/repository/docker"
	"github.com/johngithiyon/Nodefy/internal/routes"
)

func Appstartup() {
	 
	 routes.Routes()

	 // docker connections 

	 dockerconnerr :=  docker.DockerConn()

	 if dockerconnerr != nil {
		 log.Println("Docker Conn Error",dockerconnerr)
		 return 
	 }


}