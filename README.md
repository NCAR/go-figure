# go-figure
[![GoDoc](https://godoc.org/github.com/NCAR/go-figure?status.svg)](https://godoc.org/github.com/NCAR/go-figure)  [![codebeat badge](https://codebeat.co/badges/deb537c1-7654-41a5-8bf8-1a8166904af0)](https://codebeat.co/projects/github-com-ncar-go-figure) [![Build Status](https://travis-ci.org/NCAR/go-figure.svg?branch=master)](https://travis-ci.org/NCAR/go-figure) [![Coverage Status](https://coveralls.io/repos/github/NCAR/go-figure/badge.svg?branch=master)](https://coveralls.io/github/NCAR/go-figure?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/NCAR/go-figure)](https://goreportcard.com/report/github.com/NCAR/go-figure)

*go-figure* is a multi-file, multi-configuration package configuration  to handle old-school configuration techniques such searching for as multiple files in multiple folders (`~/something.d, /etc/something.d, ...`).

It does expand standard shell parameters in the path list via `os.ExpandEnv()`.

## Usage

Generally the following should be done:

```go
    ...
    v := viper.New()
    v.SetConfigType("yml|json|yaml")
    gofigure.Parse(v, []string{"${HOME}/.app.d", "${CFGD}", "/etc/app.d"})
    ...
    //do something with v

```
