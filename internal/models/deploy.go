package models

type Deploy struct {
	Appname   string   `json:"appname"`
	Gitrepo   string   `json:"gitrepo"`
	Languages []string `json:"languages"`
	Services  []string `json:"services"`
}