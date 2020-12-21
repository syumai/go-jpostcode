GOTEST=go test -v
TESTPKGS=$(shell go list ./... | grep -v -e example -e statik)

test:
# exclude TestAll_
	$(GOTEST) -race -run='^Test_' $(TESTPKGS)

test/all:
	$(GOTEST) $(TESTPKGS)

test/all/race:
	$(GOTEST) -race $(TESTPKGS)

