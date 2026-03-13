package services

import (
	"github.com/google/uuid"
) 

func GenerateSessionStore(username string) string {

	uuid := uuid.New()

	sid := uuid.String()

	id := sid[0:16]

	return id


}