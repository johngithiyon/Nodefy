package utils

import "strings"

func Trimstring(str string,suffix string) string {

	      trim := strings.TrimSuffix(str,suffix)

		  return trim
	       
}