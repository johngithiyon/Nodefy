package utils

import "regexp"



func Regexcompile() *regexp.Regexp {
	 
	Emailcheck := regexp.MustCompile(`^[a-z0-9._%+-]+@gmail\.com$`)

	return Emailcheck

}

