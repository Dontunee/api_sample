package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (app *application) writeJson(data interface{}, status int, w http.ResponseWriter) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func (app *application) readTzQueryParam(r *http.Request) ([]string, bool) {
	queryValues := r.URL.Query()

	tz := queryValues.Get("tz")
	if tz != "" {
		return strings.Split(tz, ","), true
	}
	return []string{}, false
}
