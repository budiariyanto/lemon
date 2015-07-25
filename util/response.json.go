package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Data interface{} 	`json:"data"`
	Links string 		`json:"links,omitempty"`
}

func JsonResponse(w http.ResponseWriter, data interface{}) {
	var response Response
	response.Data = data
	jsonString, _ := json.MarshalIndent(response, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonString))

}
