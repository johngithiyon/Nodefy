package postgres

import (
	"database/sql"
	"log"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
)


func  CheckUserexists(username string) bool {

	      var user string 

        scanerr :=  Database.Db.QueryRow("select username from users where username = $1",username).Scan(&user)
       
	    if scanerr == sql.ErrNoRows {
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

func SaveInstance(instancename string,username string) error {

	    log.Println(instancename,username)

		updatequery := `
			UPDATE users
			SET containers = array_append(COALESCE(containers,'{}'),$1)
			WHERE username=$2
			`
	    _,updaterr :=  Database.Db.Exec(updatequery,instancename,username)
		
	    if updaterr != nil {
			log.Println("Update Err in container array",updaterr)
			return updaterr 
		}	

	return nil 
}

func CheckInstance(killinstancename string,username string) int {
 
	var res int 
	
	checkquery := `
	            SELECT array_position(containers,$1)
				FROM users
				WHERE username=$2;
	`

	Database.Db.QueryRow(checkquery,killinstancename,username).Scan(&res)

	if res == 1 {
		 return res
	}

	return 0
}

func RemoveInstance(killinstancename string,username string) error {
 
	deletequery := `
	         UPDATE users
			SET containers = array_remove(containers,$1)
			WHERE username=$2;

	`

	_,delerr := Database.Db.Exec(deletequery,killinstancename,username)

	if delerr != nil {
		log.Println("Delete Err from instance kill",delerr)
		return errors.ErrInternalserver
	}

	return nil 
}
