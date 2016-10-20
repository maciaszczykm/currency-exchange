# Currency Exchange

Currency exchange service written in Go.

## Current status

[![Go Report Card](https://goreportcard.com/badge/github.com/maciaszczykm/currency-exchange)](https://goreportcard.com/report/github.com/maciaszczykm/currency-exchange)

## How to run?


## How to execute tests?

Open project directory in your terminal and use following command:

``` shell
go test ./...
```

## Usage

Asumming application is running on `localhost:8080` to query for a currency exchange rates you can use following
commands in your terminal:

``` shell
curl "http://localhost:8080/convert?amount=200&currency=SEK"

```

Returns JSON result for 200 SEK.


``` shell
curl -H "Accept: application/xml, */*" "http://localhost:8080/convert?amount=300&currency=PLN"

```

Returns XML result for 300 PLN.

Otherwise, you can just open `http://localhost:8080/convert?amount=200&currency=SEK`in your web browser.


## Possible enhacements

- change `float64` to more precise type, for example `github.com/shopspring/decimal` can be used easily with current 
code,
- build Docker image,
- deploy application to free server,
- add CI.