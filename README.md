GoPuny
======

[![Build Status](https://travis-ci.org/olshevskiy87/gopuny.svg?branch=master)](https://travis-ci.org/olshevskiy87/gopuny) [![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://lbesson.mit-license.org/) [![Go Report Card](https://goreportcard.com/badge/github.com/olshevskiy87/gopuny)](https://goreportcard.com/report/github.com/olshevskiy87/gopuny)

A package and CLI tool to use [SAPO](http://sapo.pt/)'s URL shortening service.

Install
=======

```bash
$ go get github.com/olshevskiy87/gopuny
```

Usage
=====

1. CLI tool `gopuny`

```bash
$ gopuny -h
Usage: gopuny ACTION URL

Positional arguments:
  ACTION                 "short" or "expand"
  URL                    url to short or expand

Options:
  --help, -h             display this help and exit
```

```bash
$ gopuny short https://www.google.com
Unicode   http://㐃鍍.sl.pt
ASCII     http://37.sp.sl.pt
Preview   http://37.sp.sl.pt/-
Original  https://www.google.com
```

```bash
$ gopuny expand http://37.sp.sl.pt
https://www.google.com/
```

2. Using package `punyurl`

```go
package main

import (
	"fmt"
	"os"

	"github.com/olshevskiy87/gopuny/punyurl"
)

func main() {
	puny, err := punyurl.New("https://www.google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	result, err := puny.Short()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%#v\n", result)
	// output:
	// &result.Result{ASCII:"http://37.sp.sl.pt", Preview:"http://37.sp.sl.pt/-", Puny:"http://㐃鍍.sl.pt", URL:"https://www.google.com"}
}
```

Motivations
===========

Inspired by perl-module [WWW::Shorten::PunyURL](http://search.cpan.org/dist/WWW-Shorten-PunyURL/)

License
=======

MIT. See LICENSE for details.
