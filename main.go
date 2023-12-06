package ezjson

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type JSONPayLoad struct {
	Data    any    `json:"data,omitempty"`
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func ReadRequest(w http.ResponseWriter, r *http.Request, destination any, maxBytesLength int64) error {
	var maxBytes int64
	if maxBytesLength == 0 {
		maxBytes = 1048576
	} else {
		maxBytes = maxBytesLength
	}

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(destination)
	if err != nil {
		return err
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only have a single json value")
	}

	return nil
}

func WriteResponse(w http.ResponseWriter, status int, responseData JSONPayLoad, headers ...http.Header) error {
	jsonOutput, err := json.Marshal(responseData)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, val := range headers[0] {
			w.Header()[key] = val
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(jsonOutput)
	if err != nil {
		return err
	}

	return nil
}

func WriteErrorResponse(w http.ResponseWriter, err error, statusCodes ...int) error {
	statusCode := http.StatusBadRequest

	if len(statusCodes) > 0 {
		statusCode = statusCodes[0]
	}

	var payload JSONPayLoad
	payload.Error = true
	payload.Message = err.Error()

	return WriteResponse(w, statusCode, payload)
}
