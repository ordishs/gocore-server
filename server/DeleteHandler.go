package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	host := vars["host"]
	address := vars["address"]
	serviceName := vars["serviceName"]
	key := fmt.Sprintf("%s:%s:%s", host, address, serviceName)

	mu.Lock()
	defer mu.Unlock()

	delete(services, key)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Deleted"))
}
