# Checkfile

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
