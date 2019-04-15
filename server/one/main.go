// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

var (
	brexitDate = "29th March"
)

func brexitDateHandler(w http.ResponseWriter, r *http.Request) {
	println("server handling request: " + r.Method + " " + r.URL.String())
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte(brexitDate))
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		brexitDate = string(body)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func main() {
	http.HandleFunc("/brexitDate", brexitDateHandler)
	log.Fatal(http.ListenAndServe("localhost:1234", nil))
}
