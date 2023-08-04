# go-test-tutorial

[![Actions Status](https://github.com/go-zen-chu/golang-template/workflows/ci/badge.svg)](https://github.com/go-zen-chu/golang-template/actions)
[![GitHub issues](https://img.shields.io/github/issues/go-zen-chu/golang-template.svg)](https://github.com/go-zen-chu/golang-template/issues)

Go test tutorial repository.

```bash
# this command has bug
$ go run cmd/shp1
# try debugging

# for not VSCode user
$ cd usecase/shp1
$ go install github.com/cweill/gotests/...   
# generate table driven test
$ gotests -all -w credit_service.go 

# this command also has bug
$ go run cmd/shp2
# try debugging too
```
