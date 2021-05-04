all: checkfile

checkfile:
	CGO_ENABLED=0 go build -o bin/checkfile cmd/checkfile/main.go

# Custom targets for checkfile
.PHONY: test test-verify test-lint test-unit

test: test-verify test-unit test-lint

# Verify if go.sum is valid
test-verify: save-sha-mod verify compare-sha-mod

# Unit test
test-unit:
	go test -v ./...

# Test code lint
test-lint:
	golangci-lint run ./... -v

save-sha-mod:
	$(eval MODSHA=$(shell sha512sum go.mod))
	$(eval SUMSHA=$(shell sha512sum go.sum))

verify:
	go mod verify

compare-sha-mod:
	$(eval MODSHA_AFTER=$(shell sha512sum go.mod))
	$(eval SUMSHA_AFTER=$(shell sha512sum go.sum))
	@if [ "${MODSHA_AFTER}" = "${MODSHA}" ]; then echo "go.mod is not changed"; else echo "go.mod file is changed"; exit 1; fi
	@if [ "${SUMSHA_AFTER}" = "${SUMSHA}" ]; then echo "go.sum is not changed"; else echo "go.sum file is changed"; exit 1; fi
