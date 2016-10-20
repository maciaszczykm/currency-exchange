package main

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/maciaszczykm/currency-exchange/handler/converter"
	"log"
	"net/http"
	"os"
)

// main runs currency-exchange application. Setting system variable $PORT forces application to run on different port,
// 8080 is default (needed for Heroku setup).
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		// Run on localhost:8080 by default.
		port = ":8080"
	} else {
		port = fmt.Sprintf(":%s", port)
	}

	log.Printf("Starting currency-exchange application on localhost%s\n", port)
	restful.Add(converter.ConverterHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
