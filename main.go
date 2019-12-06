package main

import (
	"fmt"
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)
	log.Infof(context, "Request from IP Address \"%v\"\n", r.RemoteAddr)
	fmt.Fprintln(w, r.RemoteAddr)

}

func main() {
	http.HandleFunc("/", rootHandler)
		appengine.Main()
}
