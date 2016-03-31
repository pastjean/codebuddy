# codebuddy

[![Build Status](https://travis-ci.org/pastjean/codebuddy.svg?branch=master)](https://travis-ci.org/pastjean/codebuddy)
[![GoDoc](http://godoc.org/github.com/pastjean/codebuddy?status.svg)](http://godoc.org/github.com/pastjean/codebuddy)

> Have a buddy for the week


## Download

1. Get it from [Github releases](https://github.com/pastjean/codebuddy/releases/latest/)
2. Then create your [`teams.json`](teams.json) file
3. Execute
```sh
cat teams.json | buddy
```

## Development

Have a go environment set up

```sh
git clone git@github.com:pastjean/codebuddy.git
cd codebuddy
go get ./...
go build ./cmd/buddy

cat teams.json | ./buddy
```
