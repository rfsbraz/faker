faker [![GoDoc](https://godoc.org/github.com/rfsbraz/faker?status.png)](https://godoc.org/github.com/rfsbraz/faker) [![Build Status](https://travis-ci.org/rfsbraz/faker.svg?branch=master)](https://travis-ci.org/rfsbraz/faker)

*Faker* is a golang package that generates fake data.

It was inspired by '[Python Faker](https://github.com/joke2k/faker)' from [joke2k](https://github.com/joke2k) and '[Ruby Faker](https://github.com/stympy/faker)' from [stympy](https://github.com/stympy)

----

# Getting Started

## Install

```
go get -u github.com/rfsbraz/faker
```

## Generate data

```go
package main

import (
	"github.com/rfsbraz/faker"
	"log"
)

func main() {
	faker := NewFaker("pt_PT")
	log.Println(faker.Person.Name())
}

```

# Roadmap

----

## v1.0

Version 1.0 will feature the following modules:

* Address
* Barcode
* Color
* Company
* CreditCard
* Currency
* DateTime
* File
* Internet
* Job
* Lorem
* Miscellenous
* Person
* PhoneNumber
* Profile
* UserAgent