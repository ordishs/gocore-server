package main

import (
	"gocore_server/webapp"
	"net/http"
	"os"
	"os/signal"

	"gocore_server/server"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ordishs/gocore"
)

// Name used by build script for the binaries. (Please keep on single line)
const progname = "gocore-server"

// Version & commit strings injected at build with -ldflags -X...
var version string
var commit string

var logger = gocore.Log(progname)

func init() {
	gocore.SetInfo(progname, version, commit)
}

func main() {
	stats := gocore.Config().Stats()
	logger.Infof("STATS\n%s\nVERSION\n-------\n%s (%s)\n\n", stats, version, commit)

	go func() {
		profilerAddr, ok := gocore.Config().Get("accountManager_profilerAddr")
		if ok {
			logger.Infof("Starting profile on http://%s/debug/pprof", profilerAddr)
			logger.Fatalf("%v", http.ListenAndServe(profilerAddr, nil))
		}
	}()

	// setup signal catching
	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan

		appCleanup()
		os.Exit(1)
	}()

	start()
}

func appCleanup() {
	logger.Infof("Shutting down...")
}

func start() {
	r := mux.NewRouter()
	r.HandleFunc("/api", server.GetHandler).Methods("GET")
	r.HandleFunc("/api", server.PostHandler).Methods("POST")
	r.HandleFunc("/api/{host}/{address}/{serviceName}", server.DeleteHandler).Methods("DELETE")

	r.PathPrefix("/").HandlerFunc(webapp.AppHandler).Methods("GET")

	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "DELETE", "OPTIONS"})

	address, _ := gocore.Config().Get("address", ":8889")

	certFile, _ := gocore.Config().Get("certFile")
	keyFile, _ := gocore.Config().Get("keyFile")

	if certFile != "" && keyFile != "" {
		logger.Infof("Starting TLS server on %s", address)
		logger.Fatal(http.ListenAndServeTLS(address, certFile, keyFile, handlers.CORS(originsOk, methodsOk)(r)))
	} else {
		logger.Infof("Starting server on %s", address)
		logger.Fatal(http.ListenAndServe(address, handlers.CORS(originsOk, methodsOk)(r)))
	}
}
