package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func writeJSON(w http.ResponseWriter, model interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	jsonModel, err := json.Marshal(model)
	if err != nil {
		return err
	}

	if _, err := w.Write(jsonModel); err != nil {
		return err
	}

	return nil
}

func readJSON(r *http.Request, model interface{}) error {
	bod, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(bod, model)
	return err
}

func writeError(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	errJSON := &struct {
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	errJSON.Error.Message = err.Error()
	writeJSON(w, errJSON)
}
