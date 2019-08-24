package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", IndexHandler)
	router.POST("/snip", SnipHandler)
	router.GET("/:shortURL", RedirectHandler)
	router.NotFound = http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":3000", router)
}
