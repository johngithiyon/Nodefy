package models

type Deploy struct {
	Username string    `json:"username"`
	Appname   string   `json:"appname"`
	Gitrepo   string   `json:"gitrepo"`
	Languages []string `json:"languages"`
	Services  []string `json:"services"`
	Domainame string   `json:"domainame"`
    Portnumber string  `json:"portnumber"`
	Containerip string 
}