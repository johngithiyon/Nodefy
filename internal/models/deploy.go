package models


type Deploy struct {
	   Instancename string `json:"instancename"`
       OsName string `json:"osname"`
	   Services []string `json:"services"`
}