# Checkfile

[![Go Report Card](https://goreportcard.com/badge/github.com/tmax-cloud/checkfile)](https://goreportcard.com/report/github.com/tmax-cloud/checkfile)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tmax-cloud/checkfile)
![GitHub](https://img.shields.io/github/license/tmax-cloud/checkfile)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/tmax-cloud/checkfile)

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
