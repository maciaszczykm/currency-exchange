# Currency Exchange

Currency exchange service written in Go.

## Current status

[![Build Status](https://travis-ci.org/maciaszczykm/currency-exchange.svg?branch=master)](https://travis-ci.org/maciaszczykm/currency-exchange)
[![Go Report Card](https://goreportcard.com/badge/github.com/maciaszczykm/currency-exchange)](https://goreportcard.com/report/github.com/maciaszczykm/currency-exchange)

## How to run?

Make sure, that you have valid `$GOPATH` set.

Clone repository into `$GOPATH/src/github.com/maciaszczykm/`:

``` shell
mkdir -p $GOPATH/src/github.com/maciaszczykm/
cd $GOPATH/src/github.com/maciaszczykm/
git clone https://github.com/maciaszczykm/currency-exchange.git
```

Run application:

``` shell
go run main.go
```

## How to use?

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


## How to execute tests?

Open project directory in your terminal and use following command:

``` shell
go test ./...
```

## Possible enhacements

- add more tests,
- fix Heroku setup, currently it gives an error R10 after deployment (due to failed web process binding to `$PORT`),
- change `float64` to more precise type, for example `github.com/shopspring/decimal` can be used easily with current 
code.
