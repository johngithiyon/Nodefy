package models

type Deployform struct {
	       Username string 
	       Appname string `json:"appname"`
	       Domainame string `json:"domainame"`
		   Portnumber string `json:"portnumber"`
		   Containerip string 
}