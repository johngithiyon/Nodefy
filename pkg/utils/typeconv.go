package utils

import (
	"log"
	"strconv"
)

func StrtoInt(ids string) (int,error){
       
	  id,converr := strconv.Atoi(ids)

	  if converr != nil {
		log.Println("String to Int conv err",converr)
		return 0,converr
	}

	return id,nil 
}