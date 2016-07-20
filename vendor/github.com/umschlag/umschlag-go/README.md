# Umschlag: SDK for Golang

[![Build Status](http://github.dronehippie.de/api/badges/umschlag/umschlag-go/status.svg)](http://github.dronehippie.de/umschlag/umschlag-go)
[![Coverage Status](http://coverage.dronehippie.de/badges/umschlag/umschlag-go/coverage.svg)](http://coverage.dronehippie.de/umschlag/umschlag-go)
[![Go Doc](https://godoc.org/github.com/umschlag/umschlag-go?status.svg)](http://godoc.org/github.com/umschlag/umschlag-go)
[![Go Report](http://goreportcard.com/badge/github.com/umschlag/umschlag-go)](http://goreportcard.com/report/github.com/umschlag/umschlag-go)
[![Join the chat at https://gitter.im/umschlag/umschlag](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/umschlag/umschlag)
![Release Status](https://img.shields.io/badge/status-beta-yellow.svg?style=flat)

**This project is under heavy development, it's not in a working state yet!**

Where does this name come from or what does it mean? It's quite simple, it's one
german word for transshipment, I thought it's a good match as it is related to
containers and a harbor.

This project acts as a client SDK implementation written in Go to interact with
the Umschlag API implementation. You can find the sources of the Umschlag API at
https://github.com/umschlag/umschlag-api.

The structure of the code base is heavily inspired by Drone, so those credits
are getting to [bradrydzewski](https://github.com/bradrydzewski), thank you for
this awesome project!


## Install

```
go get -d github.com/umschlag/umschlag-go/umschlag
```


## Usage

Import the package:

```go
import (
  "github.com/umschlag/umschlag-go/umschlag"
)
```

Create the client:

```go
const (
  host  = "http://umschlag.example.com"
  token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZXh0IjoiYWRtaW4iLCJ0eXBlIjoidXNlciJ9.rm4cq4Jupb8BvvDdbwyVwC3rr_WDpdEbCTO0-DCYTWQ"
)

client := umschlag.NewClientToken(host, token)
```

Get the current user:

```go
profile, err := client.ProfileGet()
fmt.Println(profile)
```

For a further reference please take a look at Godoc, you can find a link to it
above within the list of badges.


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2016 Thomas Boerger <http://www.webhippie.de>
```