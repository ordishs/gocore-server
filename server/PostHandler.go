package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	decoder := json.NewDecoder(r.Body)
	var j interface{}

	_ = decoder.Decode(&j)

	jmap, ok := j.(map[string]interface{})
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("BAD"))
		return
	}

	host, ok := jmap["host"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("No host"))
		return
	}

	address, ok := jmap["address"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("NO address"))
		return
	}

	serviceName, ok := jmap["serviceName"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("NO serviceName"))
		return
	}

	key := fmt.Sprintf("%s:%s:%s", host, address, serviceName)

	jmap["_ts"] = time.Now().UTC().Format(time.RFC3339)

	services[key] = j

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("Created"))
}
