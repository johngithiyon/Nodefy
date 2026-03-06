package postgres

import (
	"database/sql"


	"github.com/johngithiyon/Nodefy/internal/models"
)


func  CheckUserexists(signup *models.Signup) (bool,error) {

       _,usernamerr :=  Database.Db.Exec("select username from users where username = $1",signup.Username)
       
	   if usernamerr != nil && usernamerr == sql.ErrNoRows {   
            return true,nil 
		        
	   } else {  
		   return false,usernamerr
	   }

	   
}
