package models

import "database/sql"

type Dbconn struct {
	 
	   Db *sql.DB
}