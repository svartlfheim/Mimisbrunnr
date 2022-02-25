package server

import (
	"net/http"

	"github.com/rs/zerolog"
)

func handleError(w http.ResponseWriter, l zerolog.Logger, err error) bool {
	if err == nil {
		return false
	}

	statusCode := 500
	message := "An unknown error occurred!"

	switch err.(type) {
	case ErrBadRequestInputData, ErrEmptyRequestBodyNotAllowed:
		statusCode = 400
		message = err.Error()
		l.Warn().Err(err).Int("status-code", statusCode).Msg("bad request data received")
	default:
		l.Error().Err(err).Int("status-code", statusCode).Msg("unknown error occurred")
	}

	w.WriteHeader(statusCode)

	if _, err := w.Write([]byte(message)); err != nil {
		l.Fatal().Err(err).Msg("failed to write response")
	}

	return true
}