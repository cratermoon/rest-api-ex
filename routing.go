package main

import (
	"github.com/julienschmidt/httprouter"
	"fmt"
	"log"
	"math/rand"
  "net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "It works!\n")
}

type Router struct {
  router *httprouter.Router
}

var quips = []string{
	"I cannot say that truth is stranger than fiction,\nbecause I have never had acquaintance with either.\n",
	"All of the significant battles are waged within the self.\n",
	"An old pine tree preaches wisdom,\nAnd a wild bird is crying out Truth.\n",
	"This whole world reminds me of my dog.\nMy dog reminds of this whole world.\n",
	"Don't bite my finger, look where I'm pointing!",
}

func about(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	now := time.Now().Format(time.RFC1123Z)
	fmt.Fprintf(w, "The time is now %s. Time for a wise saying!\n\n", now)
	fmt.Fprintf(w,quips[rand.Intn(len(quips))])
}

func NewRouter() Router  {
	r := new(Router)
	r.router = httprouter.New()
	return *r
}

func (r Router) bind(path string, fn func(w http.ResponseWriter, r *http.Request, _ httprouter.Params)) {
	r.router.GET(path, fn)
}
func (r Router) Start() {
  r.router.GET("/", index)
  r.router.GET("/about", about)

  log.Fatal(http.ListenAndServe(":8080", r.router))
}
