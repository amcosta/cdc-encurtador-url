package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amcosta/cdc-encurtador/url"
)

var (
	porta   int
	urlBase string
)

func init() {
	porta = 8888
	urlBase = fmt.Sprintf("http://localhost:%d", porta)
}

func main() {
	http.HandleFunc("/api/encurtar", url.Encurtador)
	http.HandleFunc("/r/", Redirecionador)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", porta), nil))
}
