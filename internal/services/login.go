package services

import (
	"log"

	err "github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
)

func LoginServices(username string,Password string ) (string,error) {

	//check the user exist or not 

	 checkexists := postgres.CheckUserexists(username)
 
	if checkexists  {
		return "",err.ErrAuthenticate
	}
	
	 //search the password in the postgres 

	 passwd,passwderr := postgres.SearchPassword(username)

	 if  passwd == "" && passwderr != nil {
		   log.Println("Passwd err",passwderr)
		   return "",err.ErrInternalserver
	 }

	 //compare the password 

	 comparerr := Comparepassword(Password,passwd)

	 if comparerr != nil {
		  log.Println("compare",comparerr)
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