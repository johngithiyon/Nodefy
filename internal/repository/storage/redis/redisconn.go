package redis

import (
	"os"

	"github.com/johngithiyon/Nodefy/internal/models"
	red "github.com/redis/go-redis/v9"
)


var Redisconn models.Redisconn


func RedisGetConnection() {

      conn := red.NewClient(&red.Options{
		    
		Addr: os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB: 0,
	  })

	  Redisconn.Redisconn = conn 
	    
}