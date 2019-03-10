package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/go-toschool/dlog"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func main() {
	port := flag.Int("port", 8092, "HTTP server port")
	kind := flag.String("type", "json", "log style (json, or text)")

	flag.Parse()

	http.HandleFunc("/log", logHandler)

	if *kind == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{})
	}

	logger.SetOutput(os.Stdout)

	logger.WithFields(logrus.Fields{
		"port": *port,
	}).Info("running server")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		logger.Fatal(err)
	}
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error while read body", http.StatusBadRequest)
			return
		}

		var msg dlog.Message
		if err := json.Unmarshal(body, &msg); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer r.Body.Close()
		if msg.Level == "info" {
			logger.WithFields(logrus.Fields{
				"service": msg.Service,
				"time":    time.Now(),
			}).Info(msg.Info)
		} else if msg.Level == "error" {
			logger.WithFields(logrus.Fields{
				"service": msg.Service,
				"time":    time.Now(),
			}).Error(msg.Error)
		} else if msg.Level == "warn" {
			logger.WithFields(logrus.Fields{
				"service": msg.Service,
				"time":    time.Now(),
			}).Warn(msg.Warn)
		}
		w.WriteHeader(http.StatusOK)
	default:
		fmt.Fprintf(w, "Sorry, only POST method is supported")
	}
}
