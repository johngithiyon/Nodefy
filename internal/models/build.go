package models


type Build struct {
	   Username string `json:"username"`
	   Instancename string `json:"instancename"`
	   Services []string `json:"services"`
}