# Harbor: CLI client

[![Build Status](http://github.dronehippie.de/api/badges/harborapp/harbor-cli/status.svg)](http://github.dronehippie.de/harborapp/harbor-cli)
[![Coverage Status](http://coverage.dronehippie.de/badges/harborapp/harbor-cli/coverage.svg)](http://coverage.dronehippie.de/harborapp/harbor-cli)
[![Go Doc](https://godoc.org/github.com/harborapp/harbor-cli?status.svg)](http://godoc.org/github.com/harborapp/harbor-cli)
[![Go Report](http://goreportcard.com/badge/github.com/harborapp/harbor-cli)](http://goreportcard.com/report/github.com/harborapp/harbor-cli)
[![Join the chat at https://gitter.im/harborapp/harbor](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/harborapp/harbor-api)
![Release Status](https://img.shields.io/badge/status-beta-yellow.svg?style=flat)

**This project is under heavy development, it's not in a working state yet!**

This project acts as a CLI client implementation to interact with the Harbor
API implementation. You can find the sources of the Harbor API at
https://github.com/harborapp/harbor-api.

The structure of the code base is heavily inspired by Drone, so those credits
are getting to [bradrydzewski](https://github.com/bradrydzewski), thank you for
this awesome project!


## Install

You can download prebuilt binaries from the GitHub releases or from our
[download site](http://dl.webhippie.de/harbor-cli). You are a Mac user? Just take
a look at our [homebrew formula](https://github.com/harborapp/homebrew-harbor).
If you are missing an architecture just write us on our nice
[Gitter](https://gitter.im/harborapp/harbor-api) chat. Take a look at the help
output, you can enable auto updates to the binary to avoid bugs related to old
versions. If you find a security issue please contact thomas@webhippie.de first.


## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions](http://golang.org/doc/install.html).
As this project relies on vendoring of the dependencies and we are not
exporting `GO15VENDOREXPERIMENT=1` within our makefile you have to use a Go
version `>= 1.6`

```bash
go get -d github.com/harborapp/harbor-cli
cd $GOPATH/src/github.com/harborapp/harbor-cli
make deps build

bin/harbor-cli -h
```


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
