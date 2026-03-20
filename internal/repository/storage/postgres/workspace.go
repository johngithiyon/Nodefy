package postgres

import (
	"database/sql"
	"log"
)

func SaveWorkspace(username string ,url string ) error {
            
	       insertworkspace := "INSERT INTO workspace (username, workspace_url) VALUES($1,$2)"

		  _,inserterr :=  Database.Db.Exec(insertworkspace,username,url)

		  if inserterr != nil {
			    log.Println("Insert Err in workspace",inserterr)
				return inserterr
		  }

		  return  nil 
}

func CheckWorkspace(username string) (string,error) {

	    url := ""

	    checkworkspace := "SELECT  workspace_url FROM workspace WHERE username = $1"

	   scanerr := Database.Db.QueryRow(checkworkspace,username).Scan(&url)

	   if scanerr != nil {
		   if scanerr == sql.ErrNoRows {
			   return "",nil 
		   } else {
			   log.Println("Scan Err from check workspace",scanerr)
			   return "",scanerr
		   }
	   } 

	   return  url,nil 

       
}