// frontend project main.go
package main

import (
	"frontend/core"
	//	"frontend/structs"
	// "fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ajax/", handlerAjax)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("fonts"))))
	http.ListenAndServe(":8000", nil)
}

func handlerAjax(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if strings.Contains(r.URL.Path, ".html") {
			// Static handler
			f, _ := ioutil.ReadFile("." + r.URL.Path)
			w.Write(f)
		} else {
			// Dynamic handler (Controller)
			path := strings.Split(r.URL.Path, "/")
			if path[1] == "ajax" {
				if path[2] == "arena" {
					if len(path) == 3 {
						core.GetGames(w)
					}
				}
				if path[2] == "rating" {
					if len(path) == 3 {
						core.GetRating(w)
					}
				}
			}
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	f, _ := ioutil.ReadFile("./html/index.html")
	w.Write(f)
}
