package utils

import "strings"

func Istringcheck(str string , check string) bool  {
 
          if strings.HasSuffix(str,check) {
			     return  true 
		  }  

		  return false
}