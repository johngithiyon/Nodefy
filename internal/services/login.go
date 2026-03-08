package services

import (
	err "github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
)

func LoginServices(username string,Password string ) (string,error) {

	//check the user exist or not 

	 checkexists := postgres.CheckUserexists(username)
 
	if !checkexists  {
		return "",err.ErrAuthenticate
	}
	
	 //search the password in the postgres 

	 passwd,passwderr := postgres.SearchPassword(username)

	 if  passwd == "" && passwderr != nil {
		   return "",err.ErrInternalserver
	 }

	 //compare the password 

	 comparerr := Comparepassword(Password,passwd)

	 if comparerr != nil {
		  return "",err.ErrAuthenticate
	 }

     //generate the session id 

	 id := GenerateSessionStore(username)

	 //set the id as  key and username as  a value in the redis

	 seterr :=  redis.Storesessionid(id,username)

	 if seterr != nil {
		  return "",err.ErrInternalserver
	 }

	 return  id,nil 
 
}