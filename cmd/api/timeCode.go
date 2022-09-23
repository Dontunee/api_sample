package main

import (
	"github.com/dontunee/api_sample/cmd/internal/data"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const defaultError = "Error occurred processing the request"

const invalidTimeZoneError = "invalid timezone"

func (app *application) timeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var times data.Times
	tz, exist := app.readTzQueryParam(r)

	if exist {
		times.Tz = tz
	}

	data, err := times.GetTime()

	if err != nil {
		app.logger.Println(err)
		http.Error(w, invalidTimeZoneError, http.StatusNotFound)
		return
	}

	err = app.writeJson(data, http.StatusOK, w)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, defaultError, http.StatusInternalServerError)
	}
}
