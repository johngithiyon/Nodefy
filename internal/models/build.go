package models


type Build struct {
	   Instancename string `json:"instancename"`
	   Services []string `json:"services"`
}