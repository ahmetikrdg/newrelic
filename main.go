package main

import (
	"encoding/json"
	"github.com/newrelic/go-agent/v3/integrations/nrlogrus"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	/*app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("talk"),
		newrelic.ConfigLicense("eu01xx44a2ddad38de524815bf026778b2c7NRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
		func(config *newrelic.Config) {
			config.Enabled = false
			logrus.SetLevel(logrus.DebugLevel)
			config.Logger = nrlogrus.StandardLogger()
		},
	)*/

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("example"),
		newrelic.ConfigLicense("eu01xx25be526dcf0936fbaca2f2a6c379dfNRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
		func(config *newrelic.Config) {
			config.Enabled = true
			logrus.SetLevel(logrus.DebugLevel)
			config.Logger = nrlogrus.StandardLogger()
		},
	)

	if err != nil {
		logrus.Error(err.Error())
	}
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/movie", getMovies))

	http.ListenAndServe(":8800", nil)
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	logrus.Info("okey")

	json.NewEncoder(w).Encode("movies")
}
