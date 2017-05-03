# simple-proxy

A simple proxy server that forwards requests to a targeted backend.

Written in [Golang.](http://golang.org)

## About

Sometimes you just need to springboard requests because of whitelist constraints.
This little server will forward those requests for you.

## Command Line Usage

```
Description: simple-proxy is a server that will pass on requests it receives
to a backendpoint.

Usage: simple-proxy  [options...]

Server options:
    -s, --backend-scheme SCHEME      http or https
    -b, --backend-target TARGET      Host to forward to.

    -d, --debug                      Enable debugging output (default: false)

Common options:
    -h, --help                       Show this message
    -V, --version                    Show version

Example:

    simple-proxy --backend-scheme "https" --backend-target "www.bar.com"

```

## Building

This code currently requires version 1.8.1 or higher of Go.

Information on Golang installation, including pre-built binaries, is available at <http://golang.org/doc/install>.

Run `go version` to see the version of Go which you have installed.

Run `go build` inside the directory to build.

Run `go test ./...` to run the unit regression tests.

A successful build run produces no messages and creates an executable called `simple-proxy` in this
directory.

Run `go help` for more guidance, and visit <http://golang.org/> for tutorials, presentations, references and more.

## Docker Images

A prebuilt docker image is available at (http://www.docker.com) [simple-proxy](https://registry.hub.docker.com/u/composer22/simple-proxy/)

If you have docker installed, run:
```
docker pull composer22/simple-proxy:latest

or

docker pull composer22/simple-proxy:<version>

if available.
```
See /docker directory README for more information on how to run the container.

## License

(The MIT License)

Copyright (c) 2015 Pyxxel Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to
deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
sell copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.
