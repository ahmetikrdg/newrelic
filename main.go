package main

import (
	"context"
	"encoding/json"
	"github.com/newrelic/go-agent/v3/integrations/logcontext/nrlogrusplugin"
	"github.com/newrelic/go-agent/v3/newrelic"
	log "github.com/sirupsen/logrus"
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

	logger := log.New()
	logger.SetFormatter(nrlogrusplugin.ContextFormatter{})

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("examples"),
		newrelic.ConfigLicense("eu01xxac9cc171d0c4693aef5b12709b0802NRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	txn := app.StartTransaction("example")
	defer txn.End()
	ctx := newrelic.NewContext(context.Background(), txn)
	logger.WithContext(ctx).Info("Hello New Relic!")

	if err != nil {
		log.Error(err.Error())
	}
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/movie", getMovies))

	http.ListenAndServe(":8800", nil)
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("movies")
}
