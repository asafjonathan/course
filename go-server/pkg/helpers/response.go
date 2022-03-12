package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
)

func ServerError(w http.ResponseWriter, err error) {
	fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
func ResponseError(w http.ResponseWriter, message string, httpStatusCode int) {
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	json.NewEncoder(w).Encode(resp)
}
func ResponseErrors(w http.ResponseWriter, errors interface{}, httpStatusCode int) {
	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(errors)
}
func ResponseSuccess(w http.ResponseWriter, data interface{}, key string) {
	w.WriteHeader(200)
	resp := make(map[string]interface{})
	resp[key] = data
	json.NewEncoder(w).Encode(resp)

}
