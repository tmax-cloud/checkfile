# Checkfile

[![Go Report Card](https://goreportcard.com/badge/github.com/cqbqdd11519/checkfile)](https://goreportcard.com/report/github.com/cqbqdd11519/checkfile)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cqbqdd11519/checkfile)
![GitHub](https://img.shields.io/github/license/cqbqdd11519/checkfile)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/cqbqdd11519/checkfile)

Checkfile is a light-weight sha1sum checker

## Initiate files checksum DB
```console
export CHECK_FILES='/home/ubuntu/test-file,/tmp/test-tmp'
checkfile init
```

## Verify files checksum
```console
export CHECK_FILES='/home/ubuntu/test-file,/tmp/test-tmp'
checkfile verify
```
