package crudinator

import (
	"errors"
	"log"
	"net/http"
)

var (
	ErrNotFound = errors.New("Not Found")
)

func HandleError(rw http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	if errors.Is(ErrNotFound, err) {
		rw.WriteHeader(http.StatusNotFound)
	}

	log.Println(err.Error())

	return true
}
