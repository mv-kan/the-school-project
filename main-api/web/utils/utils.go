package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/logger"
)

const (
	InternalServerErrorMessage = "Internal server error. We are sorry for inconvenience!"
	BadFieldError              = "field name cannot be empty"
	NotImplemtedMessage        = "not implemented"
	NotFoundMessage            = "404 not found"
)

type ErrorMessage struct {
	Message string `json:"error"`
}

func RespondWithErrorLog(log logger.Logger, w http.ResponseWriter, code int, err string) error {
	switch {
	case code >= 500 && code < 600:
		log.Error(err)
		return RespondWithError(w, code, InternalServerErrorMessage)
	case code == 404:
		log.Error(err)
		return RespondWithError(w, code, NotFoundMessage)
	case code >= 400 && code < 500:
		log.Error(err)
		return RespondWithError(w, code, err)
	}
	log.Error(err)
	return RespondWithError(w, code, err)
}

func RespondWithError(w http.ResponseWriter, code int, message string) error {
	return RespondWithJSON(w, code, ErrorMessage{Message: message})
}

func RespondWithCode(w http.ResponseWriter, code int) error {
	w.WriteHeader(code)
	return nil
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	return nil
}

func ParseJSONFromBody[T entity.Base](body io.ReadCloser) (T, error) {
	var instance T
	if body == nil {
		return instance, errors.New("body is empty")
	}
	err := json.NewDecoder(body).Decode(&instance)
	if err != nil {
		return instance, err
	}
	validate := validator.New()
	err = validate.Struct(instance)
	if err != nil {
		return instance, err
	}
	return instance, nil
}

func ParseJSONNoValidator[T any](body io.ReadCloser) (T, error) {
	var instance T
	if body == nil {
		return instance, errors.New("body is empty")
	}
	err := json.NewDecoder(body).Decode(&instance)
	if err != nil {
		return instance, err
	}
	return instance, nil
}
