package main

import (
		"bytes"
		"encoding/json"
		"fmt"
		"github.com/julienschmidt/httprouter"
		"log"
		"net/http"
		"strconv"
)

const DATA_URL = "https://s3-us-west-2.amazonaws.com/sample-coding-dataset/organization_sample_data.csv"

var data *Organizations

type App struct {
	router Router
}

func orgCallback(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	queryValues := r.URL.Query()
	responseJson := filterAndSort(queryValues)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, responseJson)
}

func valueFromQueryParam(values []string) string {
	if (values != nil) {
		return values[0]
	}
	return ""
}

func filterAndSort(queryValues map[string][]string) string {
	var result Organization
	id := queryValues["id"]
	log.Println("id param is ", id)
	if (id != nil) {
		val, err := strconv.Atoi(id[0])
		if (err == nil) {
			result = data.Get(val)
		}
	}
	// we're only going to take the first qs value.

	name := valueFromQueryParam(queryValues["name"])
	city := valueFromQueryParam(queryValues["city"])
	state := valueFromQueryParam(queryValues["state"])
	postal := valueFromQueryParam(queryValues["postal"])
	category := valueFromQueryParam(queryValues["category"])
	data.Search(name, city, state, postal, category)
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
  enc.Encode(result)
	log.Println("result in filterAndSort is ", result)
	jsonString := b.String();
	log.Println("jsonString in filterAndSort is ", jsonString)
	return jsonString
}

func (app App) Run() {
	entries := LoadData(DATA_URL, true)
	data = populate(entries)
	app.router = NewRouter()
	app.router.bind("/organizations", orgCallback)
	app.router.Start()
}
