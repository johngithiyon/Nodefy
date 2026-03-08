package postgres

import (
	"database/sql"
	"log"

	"github.com/johngithiyon/Nodefy/internal/models"
)


func  CheckUserexists(username string) bool {

	      var user string 

        scanerr :=  Database.Db.QueryRow("select username from users where username = $1",username).Scan(&user)
       
	    if scanerr != nil && scanerr == sql.ErrNoRows {
			   return true
		}  else {
			return false 
		}
	   
}

func StoreUser(Verify *models.Verify) error {
  
  Res,inserterr := Database.Db.Exec("insert into users (username, email, password) VALUES ($1, $2, $3)",Verify.Username,Verify.Email,Verify.Password)

  rowno, _ :=  Res.RowsAffected()

  if rowno != 0 && inserterr != nil {
	  log.Println("Insert Error for Users",inserterr)
	  return inserterr
  }
	  
  return nil 
}

func SearchPassword(username string) (string,error) {
  
	var passwd string 

  searcherr := Database.Db.QueryRow("select password from users where username=$1",username).Scan(&passwd)

  if searcherr != nil {
	  log.Println("Search Err",searcherr)
	  return "",searcherr
  }
	 
	return  passwd,nil 
	 
}
