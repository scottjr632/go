package helpers

import (
    "regexp"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// WriteJSON returns a application/json response of the model
func WriteJSON(w http.ResponseWriter, model interface{}) error {
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

// ReadJSON de-serializes a JSON request into a model
func ReadJSON(r *http.Request, model interface{}) error {
	bod, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(bod, model)
	return err
}

// WriteError takes an error and returns a JSON response with the
// error string as the message
func WriteError(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	errJSON := &struct {
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	errJSON.Error.Message = err.Error()
	WriteJSON(w, errJSON)
}

func BuildGetParams(regEx string) func(string) (paramsMap map[string]string) {
	var compRegEx = regexp.MustCompile(regEx)
	return func(url string) (paramsMap map[string]string) {
		match := compRegEx.FindStringSubmatch(url)

		paramsMap = make(map[string]string)
		for i, name := range compRegEx.SubexpNames() {
			if i > 0 && i <= len(match) {
				paramsMap[name] = match[i]
			}
		}
		return
	}
}
