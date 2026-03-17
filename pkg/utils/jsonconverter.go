package utils

import (
	"encoding/json"
	"log"
)

func Jsonconvertor(data interface{}) (string,error) {

    jsondata, converr := json.Marshal(data)
    if converr != nil {
		log.Println("Json Conv Err",converr)
        return "",converr
    }
	
	return  string(jsondata),nil 
}