package response

import (
	"encoding/json"
	"net/http"
)


func Response(w http.ResponseWriter,statuscode int,message string) {

	   w.Header().Set("Content-Type", "application/json")
       w.WriteHeader(statuscode)

	    json.NewEncoder(w).Encode(
		   map[string]string {
				"message":message,
		   },
		)
	
}