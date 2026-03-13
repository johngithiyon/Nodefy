package redis

import (
	"context"
	"log"
	"time"
)

func Storetemp(id string,signupdetails map[string]interface{}) error  {
 
	 storerr :=  Redisconn.Redisconn.HSet(context.Background(),id,signupdetails).Err()

	 if storerr != nil {
		   log.Println("Store error temp data for signup",storerr)
		   return storerr
	 }

	 expirerr := Redisconn.Redisconn.Expire(context.Background(), id,5*time.Minute).Err()
	 if expirerr != nil {
		 log.Println("Error setting expiration for temp data:", expirerr)
		 return expirerr
	 }

	 return nil 

	 
}

func Getstoretemp(id string ) (map[string]string,error) {
   	 
      storedata,geterr := Redisconn.Redisconn.HGetAll(context.Background(),id).Result()

	  if geterr != nil {
		   log.Println("Get error temp data for signup",geterr)
		   return nil,geterr 
	  }

	  return storedata,nil 
}