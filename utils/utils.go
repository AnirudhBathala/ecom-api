package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJSON(r *http.Request,payload any) error {
	// check if payload is present in request
	body,err:=io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("unable to read request body")
	}

	if len(body) == 0 {
		return fmt.Errorf("missing request body")
	}
	
	return json.Unmarshal(body,payload)
}


func WriteJSON(w http.ResponseWriter,status int, v any) error {
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter,status int, err error){
	WriteJSON(w,status,map[string]string{"error":err.Error()})
}
