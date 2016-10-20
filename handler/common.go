package handler

import (
	"github.com/emicklei/go-restful"
	"log"
	"math"
)

// Round returns rounded float64 value to the specified number of places after comma.
func Round(value float64, places int) float64 {
	pow := math.Pow(10, float64(places))
	digit := pow * value
	_, div := math.Modf(digit)
	if div >= .5 {
		return math.Ceil(digit) / pow
	}
	return math.Floor(digit) / pow
}

// HandleError logs error message and writes error response with specific status code.
func HandleError(response *restful.Response, err error, status int) {
	log.Print(err)
	response.AddHeader("Content-Type", "text/plain")
	response.WriteErrorString(status, err.Error())
}
