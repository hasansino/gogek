package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	listenOn := flag.String("listen", ":8080", "http server to listen")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		payload, err := io.ReadAll(r.Body)
		if err != nil {
			return
		}
		defer func() { _ = r.Body.Close() }()
		logger.Printf("Received request [%s], with data [%s] \n", r.URL.String(), string(payload))
		w.WriteHeader(http.StatusOK)
	})

	log.Printf("Listnening on %s\n", *listenOn)
	err := http.ListenAndServe(*listenOn, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
