package utils

import (
	"encoding/json"
	"io"
	"log"

)

func Jsoncov(req io.ReadCloser,convstruct any) error {
	  
	decoderr := json.NewDecoder(req).Decode(&convstruct)

	if decoderr != nil {
		log.Println(decoderr)
		return decoderr
	 }

	 return nil 
}