package services

import (

	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func Deployproxy(domainname string) (string,string,error) {

	domainname = utils.Trimstring(domainname,":8080")
 
	containerip,portnumber,geterr := postgres.Getexposeappdetails(domainname)

	if geterr != nil {
		 return "","",geterr 
	}

	 return  containerip,portnumber,nil 
}