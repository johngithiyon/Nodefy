package services

import (
	"log"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
)

func Otpverificationservices(otp string,sessionid string,Ver *models.Verify) error {

	storedata,geterr :=  redis.Getstoretemp(sessionid)

	if storedata == nil  && geterr != nil {
		 log.Println("Get Temp Data err from otpverify",geterr)
		 return geterr
	}


	if otp == storedata["otp"] {

		Ver.Username = storedata["username"]
		Ver.Email = storedata["email"]
		Ver.Password = storedata["password"]

		storerr := postgres.StoreUser(Ver)

		if storerr != nil {
			log.Println("Store temp data err in otp verify",storerr)
			return errors.ErrInternalserver
		}
		
	} else {
		 log.Println("Authenticate err in otp verify")
		 return errors.ErrAuthenticate
	}

	return nil 

}