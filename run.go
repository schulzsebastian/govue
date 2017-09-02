package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/schulzsebastian/gowebpack/api"
)

func runWebpack() {
	cmd := exec.Command("sh", "-c", "cd static && npm run dev")
	_ = cmd.Start()
	log.Print("Running Webpack on port :5001")
}

func startHTTPServer(prod bool) {
	router := mux.NewRouter()
	router.HandleFunc("/data", api.Data)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	if prod == true {
		log.Print("NPM building...")
		cmd := exec.Command("sh", "-c", "cd static && npm run build")
		_ = cmd.Start()
		cmd.Wait()
	}
	log.Print("Running HTTP server on port :5000")
	http.ListenAndServe(":5000", router)
}

func listenStop(close chan bool) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for _ = range c {
		log.Print("Shutdowning HTTP server on port :5000")
		log.Print("Shutdowning Webpack on port :5001")
		close <- true
	}
}

func main() {
	// Listen for CTRL+C
	close := make(chan bool)
	go listenStop(close)
	// Parsing production flag
	prod := flag.Bool("P", false, "Production")
	flag.Parse()
	// Run prod or dev mode
	if *prod == true {
		go startHTTPServer(true)
	} else {
		go startHTTPServer(false)
		go runWebpack()
	}
	// Wait for value from listenStop
	<-close
}
