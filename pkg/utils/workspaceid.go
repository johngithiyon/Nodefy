package utils


import (
    "crypto/rand"
    "encoding/hex"
)


func Generateworkspaceid() (string,error) {
     
	b := make([]byte, 16)

    _, randerr := rand.Read(b)

    if randerr != nil {
        return "", randerr
    }

    return hex.EncodeToString(b), nil 
	     
}