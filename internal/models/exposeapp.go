package models

type Exposeappform struct {
	       Username string 
	       Appname string `json:"appname"`
	       Domainame string `json:"domainame"`
		   Portnumber string `json:"portnumber"`
		   Containerip string 
}