package postgres

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/johngithiyon/Nodefy/internal/models"
	_ "github.com/lib/pq"
)

var Database models.Dbconn

func  GetPostgresConnection () error {

	connstr := os.Getenv("CONN_STR")
 
	//Create a connection driver 

    postgresdb,driveropenerr := sql.Open("postgres",connstr)

	if driveropenerr != nil {
          
		  log.Println("Open driver connection error",driveropenerr)
		  return  driveropenerr  
	}  

	// connection pooling 
	postgresdb.SetMaxOpenConns(50)
    postgresdb.SetMaxIdleConns(25)
	postgresdb.SetConnMaxIdleTime(100 *time.Second)

	Database.Db = postgresdb
	
	return nil 


}