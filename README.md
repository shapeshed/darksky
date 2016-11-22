# Darksky

[![Build Status](https://travis-ci.org/shapeshed/darksky.svg?branch=master)](https://travis-ci.org/shapeshed/darksky)

A Go client for the [Dark Sky][1] API.

## Installation

    go get -v github.com/shapeshed/darksky

## Usage

The client requires that a pointer to a Struct of parameter data is passed to the Get method. 

    type RequestParams struct {
      Key       string
      Latitude  float64
      Longitude float64
      Exclude   string
      Extend    string
      Lang      string
      Units     string
    }

You must provide at least Key, Latitude and Longitude.

    package main

    import (
      "fmt"
      "github.com/shapeshed/darksky"
      "log"
    )

    func main() {
      params := darksky.RequestParams{
        Key:       "17b1e8cae7b654290659b438557def7e",
        Latitude:  52.847875,
        Longitude: -0.664397,
        Units:     "si",
      }

      forecast, err := darksky.Get(&params)
      if err != nil {
        log.Fatal(err)
      }
      fmt.Println(forecast.Currently.Temperature)

    }

## Principles

* Be lightweight. 
* Don't modify any data.
* Provide useful errors.
* Do one thing and do it well.

[1]: https://darksky.net/
