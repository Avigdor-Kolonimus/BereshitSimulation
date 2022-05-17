# Go related variables
PROJECTNAME := $(shell basename "$(PWD)")
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)
CSVFILE := *Bereshit*Landing.csv

build:
	@echo "  >  Building binary..."
	go build -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

runBoaz:
	go run main.go -algorithm 0

run:
	go run main.go -algorithm 1

runTwoPID:
	go run main.go -algorithm 2

clean:
	@echo "  >  Cleaning build cache"
	go clean
	rm $(GOBIN)/$(PROJECTNAME)
	rm $(GOBASE)/$(CSVFILE)