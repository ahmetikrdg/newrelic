package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
	/*
		Logger := log.New()
		Logger.SetFormatter(nrlogrusplugin.ContextFormatter{})
		app, err := newrelic.NewApplication(
			newrelic.ConfigAppName("end"),
			newrelic.ConfigLicense("eu01xx25be526dcf0936fbaca2f2a6c379dfNRAL"),
			newrelic.ConfigDistributedTracerEnabled(true),
		)

		txn := app.StartTransaction("example")
		defer txn.End()
		ctx := newrelic.NewContext(context.Background(), txn)
		Logger.WithContext(ctx).Info("Hello New Relic!")

		if err != nil {
			log.Error(err.Error())
		}*/
	/*
		app, err := newrelic.NewApplication(
			newrelic.ConfigAppName("bismillah"),
			newrelic.ConfigLicense("eu01xx25be526dcf0936fbaca2f2a6c379dfNRAL"),
			newrelic.ConfigDistributedTracerEnabled(true),
			func(c *newrelic.Config) {
				c.DatastoreTracer.SlowQuery.Threshold = slowQueryThreshold
				c.ErrorCollector.IgnoreStatusCodes = []int{
					http.StatusUnauthorized,
					http.StatusNotFound,
				}
				logrus.SetLevel(logrus.InfoLevel)
				c.Logger = nrlogrus.StandardLogger()
				c.Enabled = true
			},
		)*/

	/*
		w, err := os.OpenFile("my_log_file", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		app, _ := newrelic.NewApplication(
			newrelic.ConfigAppName("ddd"),
			newrelic.ConfigLicense("eu01xx25be526dcf0936fbaca2f2a6c379dfNRAL"),
			newrelic.ConfigInfoLogger(w),
		)

		logrus.Info("deneme logu")
		logrus.Error("deneme erroru")*/

	/*
			l := logrus.New()
			l.SetLevel(logrus.DebugLevel)
			app, err := newrelic.NewApplication(
				newrelic.ConfigAppName("ddd"),
				newrelic.ConfigLicense("eu01xx25be526dcf0936fbaca2f2a6c379dfNRAL"),
				newrelic.ConfigFromEnvironment(),
				nrlogrus.ConfigLogger(l),
			)


		app, err := newrelic.NewApplication(
			newrelic.ConfigAppName("ss"),
			newrelic.ConfigLicense("eu01xx25be526dcf0936fbaca2f2a6c379dfNRAL"),
			func(config *newrelic.Config) {
				logrus.SetLevel(logrus.InfoLevel)
				config.Logger = nrlogrus.StandardLogger()
			},
		)*/

	Logger := log.New()
	Logger.SetFormatter(nrlogrusplugin.ContextFormatter{})
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("new"),
		newrelic.ConfigLicense("eu01xxac9cc171d0c4693aef5b12709b0802NRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	txn := app.StartTransaction("example")
	defer txn.End()
	ctx := newrelic.NewContext(context.Background(), txn)
	Logger.WithContext(ctx).Info("Hello New Relic!")

	if err != nil {
		fmt.Println(err)
	}
	Logger.WithContext(ctx).Error("Hello New Relic!")

	http.HandleFunc(newrelic.WrapHandleFunc(app, "/movie", getMovies))
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/err", getError))
	http.ListenAndServe(":8800", nil)
}

func getMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	log.Info("coming data")
	json.NewEncoder(w).Encode("movies")
}

func getError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	errors.New("hata")
	log.Error("come error")
	json.NewEncoder(w).Encode("movies")
}
