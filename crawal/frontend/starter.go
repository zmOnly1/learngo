package main

import (
	"learngo2/crawal/frontend/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("crawal/frontend/view")))
	http.Handle("/search", controller.CreateSearchResultHandler("crawal/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
