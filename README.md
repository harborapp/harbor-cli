# Umschlag: CLI client

[![Build Status](http://drone.umschlag.tech/api/badges/umschlag/umschlag-cli/status.svg)](http://drone.umschlag.tech/umschlag/umschlag-cli)
[![Build Status](https://ci.appveyor.com/api/projects/status/bqy7swd1sn32k6vq?svg=true)](https://ci.appveyor.com/project/umschlagz/umschlag-cli)
[![Stories in Ready](https://badge.waffle.io/umschlag/umschlag-api.svg?label=ready&title=Ready)](http://waffle.io/umschlag/umschlag-api)
[![Join the Matrix chat at https://matrix.to/#/#umschlag:matrix.org](https://img.shields.io/badge/matrix-%23umschlag-7bc9a4.svg)](https://matrix.to/#/#umschlag:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/a8a9dd7a2e0a4437a56db779e38b47ed)](https://www.codacy.com/app/umschlag/umschlag-cli?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=umschlag/umschlag-cli&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/umschlag/umschlag-cli?status.svg)](http://godoc.org/github.com/umschlag/umschlag-cli)
[![Go Report](http://goreportcard.com/badge/github.com/umschlag/umschlag-cli)](http://goreportcard.com/report/github.com/umschlag/umschlag-cli)
[![](https://images.microbadger.com/badges/image/umschlag/umschlag-cli.svg)](http://microbadger.com/images/umschlag/umschlag-cli "Get your own image badge on microbadger.com")

**This project is under heavy development, it's not in a working state yet!**

Within this repository we are building the command-line client to interact with the [Umschlag API](https://github.com/umschlag/umschlag-api) server, for further information take a look at our [documentation](https://umschlag.tech).

*Where does this name come from or what does it mean? It's quite simple, it's one german word for transshipment, I thought it's a good match as it is related to containers and a harbor.*


## Install

You can download prebuilt binaries from the GitHub releases or from our [download site](http://dl.umschlag.tech/cli). You are a Mac user? Just take a look at our [homebrew formula](https://github.com/umschlag/homebrew-umschlag).


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.8.

```bash
go get -d github.com/umschlag/umschlag-cli
cd $GOPATH/src/github.com/umschlag/umschlag-cli

# install retool
make retool

# sync dependencies
make sync

# generate code
make generate

# build binary
make build

./bin/umschlag-cli -h
```


## Security

If you find a security issue please contact umschlag@webhippie.de first.


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```
