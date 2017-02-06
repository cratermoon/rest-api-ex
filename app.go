package main

import (
		"bytes"
		"encoding/json"
		"fmt"
		"github.com/julienschmidt/httprouter"
		"net/http"
)

const DATA_URL = "https://s3-us-west-2.amazonaws.com/sample-coding-dataset/organization_sample_data.csv"

type App struct {
  data *OrganizationList
	router Router
}

func orgCallback(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	queryValues := r.URL.Query()
	responseJson := filterAndSort(queryValues)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, responseJson)
}

func filterAndSort(queryValues map[string][]string) string {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	queryValues["testing"] = []string{"1,2,3"}
  enc.Encode(queryValues)
	return b.String()
}

func (app App) Run() {
	entries := LoadData(DATA_URL, true)
	app.data = populate(entries)
	app.router = NewRouter()
	app.router.bind("/organizations", orgCallback)
	app.router.Start()
}
