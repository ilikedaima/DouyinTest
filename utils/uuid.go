package utils

import (
	uuid "github.com/satori/go.uuid"
)

func GetUUID() string{

	u1 := uuid.NewV4().String()
	return u1
}

