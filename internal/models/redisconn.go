package models

import (
	
	red "github.com/redis/go-redis/v9"
)


type Redisconn struct {
	Redisconn *red.Client 
}