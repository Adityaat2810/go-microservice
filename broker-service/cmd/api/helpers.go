package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)


type jsonResponse struct {
  Error  bool  `json:"error"`
  Message  string  `json:"message"`
  Data  any  `json:"data,omitempty"`
}

// readJson - it is for reading the json from the request body
func (app *Config) readJson(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576 // one mb

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)

	if err != nil {
		return err
	}

	if err := dec.Decode(&struct{}{}); err != io.EOF {
	  return errors.New("body must only contain a single JSON value")
	}

	return nil

}

// writeJson - it is for writing the json to the response body
func (app *Config) writeJson(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
  out, err := json.Marshal(data)
  if err != nil {
    return err
  }

  if len(headers) > 0 {
	for k, v := range headers[0] {
		w.Header()[k] = v
	}
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)

  _, err = w.Write(out)
  if err != nil {
	return err
  }

  return nil
}

// errorJson - it is for writing the error json to the response body
func (app *Config) errorJson(w http.ResponseWriter, err error, status ...int) error {
  statusCode := http.StatusBadRequest

  if len(status) > 0 {
	statusCode = status[0]
  }

  var payload jsonResponse
  payload.Error = true
  payload.Message = err.Error()

  return app.writeJson(w, statusCode, payload)
}