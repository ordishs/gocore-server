package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	mu       sync.RWMutex
	services map[string]interface{}
)

func init() {
	services = make(map[string]interface{})
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	decoder := json.NewDecoder(r.Body)
	var j interface{}

	decoder.Decode(&j)

	jmap, ok := j.(map[string]interface{})
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("BAD"))
		return
	}

	host, ok := jmap["host"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No host"))
		return
	}

	address, ok := jmap["address"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("NO address"))
		return
	}

	serviceName, ok := jmap["serviceName"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("NO serviceName"))
		return
	}

	key := fmt.Sprintf("%s:%s:%s", host, address, serviceName)

	jmap["_ts"] = time.Now().UTC().Format(time.RFC3339)

	services[key] = j

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Created"))
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	host := vars["host"]
	address := vars["address"]
	serviceName := vars["serviceName"]
	key := fmt.Sprintf("%s:%s:%s", host, address, serviceName)

	mu.Lock()
	defer mu.Unlock()

	delete(services, key)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted"))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
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

	encoder.Encode(&values)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/gocore", getHandler).Methods("GET")
	r.HandleFunc("/gocore", postHandler).Methods("POST")
	r.HandleFunc("/gocore/{host}/{address}/{serviceName}", deleteHandler).Methods("DELETE")

	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "DELETE", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8889", handlers.CORS(originsOk, methodsOk)(r)))
}
