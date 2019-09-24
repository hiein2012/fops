# fops
## Overview

This is a small command-line application that could implement two functions as below:
- Print line count of input file
- Print the checksum of file, support multiple algorithms: md5, sha1 and sha256

## Installation

Make sure you have a working Go environment. Go version 1.10+ is supported. [See the install instructions for Go.](http://golang.org/doc/install.html)

To install fops, simply run:

```
$ go get github.com/hiein2012/fops
```
You can find the installed folder is in $GOPATH/src/github.com/hiein2012/fops.

Make sure your PATH includes the `$GOPATH/bin` directory so your commands can be easily used:

```
export PATH=$PATH:$GOPATH/bin
```
This application was made base on `urfave/cli`. To install cli , you can run 

```
$ go get github.com/urfave/cli
```

## Getting Started

Under the path `$GOPATH/src/github.com/hiein2012/fops` , you can compiler the `fops.go` to output a binary file.

```
$ go build fops.go
```


##  Prerequisites

Make sure you have installed all of the following prerequisites on your development machine:

1. [GO](https://golang.org/dl/) 
2. [Git](https://git-scm.com/downloads)
3. [urfave/cli](https://github.com/urfave/cli)

