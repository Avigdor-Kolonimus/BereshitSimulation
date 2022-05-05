# Go related variables
PROJECTNAME := $(shell basename "$(PWD)")
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)
CSVFILE := BereshitLanding.csv

build:
	@echo "  >  Building binary..."
	go build -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

run:
	go run main.go

clean:
	@echo "  >  Cleaning build cache"
	go clean
	rm $(GOBIN)/$(PROJECTNAME)
	rm $(GOBASE)/$(CSVFILE)