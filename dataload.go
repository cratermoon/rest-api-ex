package main

import (
  "bytes"
  "encoding/csv"
  "io/ioutil"
  "log"
  "net/http"
  "strings"
)

func download(url string) []byte {

  resp, err := http.Get(url)
  defer resp.Body.Close()
  if (err != nil) {
    log.Fatal(err)
  }
  responseBytes, err := ioutil.ReadAll(resp.Body)

  if (err != nil) {
    log.Fatal(err)
  }
	return responseBytes
}

func LoadData(url string, hasHeader bool) [][]string {
	respBytes := download(url)
	trimmed := bytes.Trim(respBytes, "\xef\xbb\xbf")
  fixedInput := strings.Replace(string(trimmed), "\r", "\n", -1)
  r := csv.NewReader(strings.NewReader(fixedInput))
	if hasHeader {
		// the first line is just header data, discard it
		header, err := r.Read()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Header read and discarded: %v\n", header)
	}

	records, err := r.ReadAll()
	if err != nil {
			log.Fatal(err)
	}
  return records
}
