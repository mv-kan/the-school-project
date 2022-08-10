package controller

import (
	"encoding/json"
	"net/http"

	"github.com/mv-kan/the-school-project/logger"
)

const (
	internalServerErrorMessage = "Internal server error. We are sorry for inconvenience!"
	badFieldError              = "field name cannot be empty"
	notImplemtedMessage        = "not implemented"
	notFoundMessage            = "not found"
)

type errorMessage struct {
	Message string `json:"error"`
}

func respondWithErrorLog(log logger.Logger, w http.ResponseWriter, code int, err string) error {
	switch {
	case code >= 500 && code < 600:
		log.Error(err)
		return respondWithError(w, code, internalServerErrorMessage)
	case code == 404:
		log.Error(err)
		return respondWithError(w, code, notFoundMessage)
	case code >= 400 && code < 500:
		log.Error(err)
		return respondWithError(w, code, err)
	}
	log.Error(err)
	return respondWithError(w, code, err)
}

func respondWithError(w http.ResponseWriter, code int, message string) error {
	return respondWithJSON(w, code, errorMessage{Message: message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	return nil
}
