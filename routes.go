package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"JSONLogsIndex",
		"POST",
		"/json/logs",
		JsonLogsIndex,
	},
	Route{
		"JSONMetricsIndex",
		"POST",
		"/json/metrics",
		JsonMetricsIndex,
	},
	Route{
		"JSONTracesIndex",
		"POST",
		"/json/traces",
		JsonTracesIndex,
	},
	Route{
		"PrometheusIndex",
		"POST",
		"/prometheus",
		PrometheusIndex,
	},
}
