package server

import (
	"encoding/json"
	"net/http"
	"sort"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	values := make([]interface{}, 0)
	for _, v := range services {
		values = append(values, v)
	}

	encoder := json.NewEncoder(w)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	sort.SliceStable(values, func(i, j int) bool {
		imap, ok := values[i].(map[string]interface{})
		if !ok {
			return i < j
		}
		jmap, ok := values[j].(map[string]interface{})
		if !ok {
			return i < j
		}

		iService, ok := imap["serviceName"].(string)
		if !ok {
			return i < j
		}

		jService, ok := jmap["serviceName"].(string)
		if !ok {
			return i < j
		}

		if iService != jService {
			return iService < jService
		}

		iHost, ok := imap["host"].(string)
		if !ok {
			return i < j
		}

		jHost, ok := jmap["host"].(string)
		if !ok {
			return i < j
		}

		if iHost != jHost {
			return iHost < jHost
		}

		iPort, ok := imap["port"].(float64)
		if !ok {
			return i < j
		}

		jPort, ok := jmap["port"].(float64)
		if !ok {
			return i < j
		}

		if iPort != jPort {
			return iPort < jPort
		}

		return i < j
	})

	_ = encoder.Encode(&values)
}
