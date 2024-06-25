package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/Alfred-Onuada/go-dropbox/internals/types"
)

func JSONResponse(w http.ResponseWriter, status bool, message string, statuscode int, data interface{}) {
	response := types.BaseResponse{
		Status: status,
		Data:    data,
		Message: message,

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	json.NewEncoder(w).Encode(response)
	return
}

