# FreePort

[![GoDoc](https://pkg.go.dev/badge/github.com/ansidev/freeport?status.svg)](https://pkg.go.dev/github.com/ansidev/freeport?tab=doc)
[![Release](https://img.shields.io/github/release/ansidev/freeport.svg)](https://github.com/ansidev/freeport/releases)
[![Build Status](https://github.com/ansidev/freeport/workflows/ci/badge.svg?branch=main)](https://github.com/ansidev/freeport/actions?query=branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/ansidev/freeport)](https://goreportcard.com/report/github.com/ansidev/freeport)
[![Sourcegraph](https://sourcegraph.com/github.com/ansidev/freeport/-/badge.svg)](https://sourcegraph.com/github.com/ansidev/freeport?badge)

Get a free open TCP port that is ready to use.

This repository is forked from [phayes/freeport](https://github.com/phayes/freeport). I forked this repository because the original repository is not maintained actively.

## Command Line Example

```bash
# Ask the kernel to give us an open port.
export port=$(freeport)

# Start standalone httpd server for testing
httpd -X -c "Listen $port" &

# Curl local server on the selected port
curl localhost:$port
```

## Golang example

```go
package main

import "github.com/ansidev/freeport"

func main() {
	port, err := freeport.GetFreePort()
	if err != nil {
		log.Fatal(err)
	}
	// port is ready to listen on
}

```

## Installation

#### Mac OSX
```bash
brew install ansidev/repo/freeport
```

#### CentOS and other RPM-based systems

```bash
wget https://github.com/ansidev/freeport/releases/download/1.0.3/freeport_1.0.3_linux_386.rpm
rpm -Uvh freeport_1.0.3_linux_386.rpm
```

#### Ubuntu and other DEB-based systems

```bash
wget wget https://github.com/ansidev/freeport/releases/download/1.0.3/freeport_1.0.3_linux_amd64.deb
dpkg -i freeport_1.0.3_linux_amd64.deb
```

#### Building From Source

```bash
sudo apt-get install golang  # Download go. Alternatively build from source: https://golang.org/doc/install/source
```

```bash
go install github.com/ansidev/freeport/cmd@latest
```

## Contact

Le Minh Tri [@ansidev](https://ansidev.xyz/about).

## License

This source code is available under the [Open Source License (BSD 3-Clause)](/LICENSE).
