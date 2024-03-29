package helper

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, payload interface{}, status int) {

	mapdata := map[string]interface{}{}
	var (
		ok = http.StatusOK
		created = http.StatusCreated
		internalservererror = http.StatusInternalServerError
		methodnotallowed = http.StatusMethodNotAllowed
	)

	switch status {
	case internalservererror:
		mapdata["status"] = internalservererror
		mapdata["message"] = "failed"
		mapdata["data"] = payload

	case methodnotallowed:
		mapdata["status"] = methodnotallowed
		mapdata["message"] = "failed"
		mapdata["data"] = payload

	case created:
		mapdata["status"] = created
		mapdata["message"] = "success"
		mapdata["data"] = payload

	default:
		mapdata["status"] = ok
		mapdata["message"] = "success"
		mapdata["data"] = payload
	}


	response, err := json.Marshal(mapdata)


	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(response)

}
