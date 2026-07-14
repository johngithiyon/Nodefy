package postgres

import (
	"database/sql"
	"log"
)

func SaveWorkspace(username string ,url string,containername string,containerip string) error {
            
	       insertworkspace := "INSERT INTO workspace (username, workspace_url,containername,containerip) VALUES($1,$2,$3,$4)"

		  _,inserterr :=  Database.Db.Exec(insertworkspace,username,url,containername,containerip)

		  if inserterr != nil {
			    log.Println("Insert Err in workspace",inserterr)
				return inserterr
		  }

		  return  nil 
}

func CheckWorkspace(userhash string) (string,error) {

	    var containername string 

	    checkworkspace := "SELECT  containerip FROM workspace WHERE  workspace_url = $1"

	   scanerr := Database.Db.QueryRow(checkworkspace,userhash).Scan(&containername)

	   if scanerr != nil {
		   if scanerr == sql.ErrNoRows {
			   return "",nil 
		   } else {
			   log.Println("Scan Err from check workspace",scanerr)
			   return "",scanerr
		   }
	   } 

	   return  containername,nil 

       
}