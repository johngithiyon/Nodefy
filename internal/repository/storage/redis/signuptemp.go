package redis

import (
	"context"
	"log"
)

func Storetemp(id string,signupdetails map[string]interface{}) error  {
 
	 storerr :=  Redisconn.Redisconn.HSet(context.Background(),id,signupdetails).Err()

	 if storerr != nil {
		   log.Println("Store error temp data for signup",storerr)
		   return storerr
	 }

	 return nil 

	 
}