package api

import (
	"encoding/json"
	"net/http"
)

type exampleResponse struct {
	Data string `json:"data"`
}

// Data -- Example routing
func Data(w http.ResponseWriter, r *http.Request) {
	resp := exampleResponse{"ok"}
	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
}
