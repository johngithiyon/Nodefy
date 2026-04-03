package app

import (
	"log"

	"github.com/johngithiyon/Nodefy/internal/config"
	"github.com/johngithiyon/Nodefy/internal/repository/docker"
	"github.com/johngithiyon/Nodefy/internal/repository/rabbitmq"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
	"github.com/johngithiyon/Nodefy/internal/routes"
)

func Appstartup() {

	//load the env file

	 loaderr := config.Loadenv()

	 if loaderr != nil {
		 return 
	 }

	 //routes for the api 
	 
	 routes.Routes()

	 //Create a connection object call

	 connerr :=  postgres.GetPostgresConnection()

	 if connerr != nil {
           log.Println("Db conn err",connerr)
		   return
	 }

	 // Create a  redis connection call 
     
	redis.RedisGetConnection()

	 // docker connections 

	 dockerconnerr :=  docker.DockerConn()

	 if dockerconnerr != nil {
		 log.Println("Docker Conn Error",dockerconnerr)
		 return 
	 }

	 //Rabbitma connections 

	 rabbitmqconnerr := rabbitmq.CreateRabbitconn()

	 if rabbitmqconnerr != nil {
            return
	 }


}