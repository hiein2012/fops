# fops
## Overview

This is a small command-line application that could implement two functions as below:
- Print line count of input file
- Print the checksum of file, support multiple algorithms: md5, sha1 and sha256

The whole command lines mentioned  below are base on linux-base OS.

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
This application was made base on `urfave/cli`. To install cli , you can run :

```
$ go get github.com/urfave/cli
```

## Getting Started

Under the path `$GOPATH/src/github.com/hiein2012/fops` , you can compiler the `fops.go` to output a binary file.

```
$ go build fops.go
```

To Run `fops` function globally , you can run : 

``` 
$ cp fops /usr/local/bin 
```

- ### Function one : Print line count of input file

You can use `linecount` to get the rows number of input file , and use `-f` or `--file` as flag to take input file with the following code in it:

```
$ fops linecount -f inputfile.text
```
In function one , the file type **does not** support the binary file.

- ### Function two : Print the checksum of file (support algorithms : md5 / sha1 / sha256 )

You can assign specific flag to get checksum of file including three algorithms :
```
$ fops checksum -f  inputfile.text --md5 
$ fops checksum -f  inputfile.text --sha1
$ fops checksum -f  inputfile.text --sha256
```
In function two , the file type support the binary file.

##  Prerequisites

Make sure you have installed all of the following prerequisites on your development machine:

1. [GO](https://golang.org/dl/) 
2. [Git](https://git-scm.com/downloads)
3. [urfave/cli](https://github.com/urfave/cli)

