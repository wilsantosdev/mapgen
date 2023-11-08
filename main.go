package main

import (
	"mapgen/worldmap"
	"net/http"
	"text/template"
)

func main() {

	tlp := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		worldmap := worldmap.NewMap(50, 10)
		gmap := worldmap.GetMap()
		tlp.Execute(w, gmap)
	})

	http.ListenAndServe(":8080", nil)

}
