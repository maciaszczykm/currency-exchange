package main

import (
	"github.com/emicklei/go-restful"
	"github.com/maciaszczykm/currency-exchange/handler/converter"
	"log"
	"net/http"
)

// main runs currency-exchange application on localhost:8080.
func main() {
	restful.Add(converter.ConverterHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
