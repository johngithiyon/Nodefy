package config

import (
	"log"
	"github.com/joho/godotenv"
)

func Loadenv () error {
   loaderr :=  godotenv.Load("internal/config/.env")

   if loaderr != nil {
	   log.Println("Load error",loaderr)
	   return loaderr
   }
   
   return nil 
}