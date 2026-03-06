package redis

import (
	"context"
	"time"
)


func SetOtp(email string,otp string)  error{

	 seterr :=   Redisconn.Redisconn.Set(context.Background(),email,otp,5*time.Minute).Err()

	 if seterr != nil {
		 return seterr
	 }

	 return nil 
}