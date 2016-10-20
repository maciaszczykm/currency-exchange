package converter

import (
	"github.com/emicklei/go-restful"
	"github.com/maciaszczykm/currency-exchange/client"
	"github.com/maciaszczykm/currency-exchange/handler"
	"log"
	"net/http"
	"strconv"
)

// ConverterHandler returns a new REST webservice with capability to convert currencies.
func ConverterHandler() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/convert").
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	service.Route(service.GET("").To(convert))

	return service
}

// convert is a handler function, returning response based on specific requests. Requests have to contain two following
// parameters:
//
// - amount - amount of money to exchange
// - currency - currency of money to exchange
func convert(request *restful.Request, response *restful.Response) {
	amount, err := strconv.ParseFloat(request.QueryParameter("amount"), 64)
	if err != nil {
		log.Println("Cannot read amount parameter from the request")
		handler.HandleError(response, err, http.StatusBadRequest)
		return
	}

	currency := request.QueryParameter("currency")
	if currency == "" {
		log.Println("Cannot read currency parameter from the request")
		handler.HandleError(response, err, http.StatusBadRequest)
		return
	}

	rates, err := client.GetCurrencyRates(currency)

	if err != nil {
		log.Printf("Cannot get currency rates for %.2f %s", amount, currency)
		handler.HandleError(response, err, http.StatusInternalServerError)
		return
	}

	// All the magic happens here
	for currency, rate := range rates {
		// TODO consider github.com/shopspring/decimal usage due to float64 inaccuraccy
		rates[currency] = handler.Round(rate*amount, 2)
	}

	result := &converterResponse{
		Amount:    amount,
		Currency:  currency,
		Converted: rates,
	}

	response.WriteEntity(result)
}
