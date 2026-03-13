package models


type Deploy struct {
       OsName string `json:"osname"`
	   Services []string `json:"services"`
}