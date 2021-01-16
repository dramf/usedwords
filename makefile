MAINPACKAGE=main
EXENAME=uw
BUILDPATH=$(CURDIR)
export GOPATH=$(CURDIR)

default : all

all : makedir get build

makedir :
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin ; fi
	@if [ ! -d $(BUILDPATH)/pkg ] ; then mkdir -p $(BUILDPATH)/pkg ; fi

build :
	@echo "building...."
	@go build -o $(BUILDPATH)/bin/$(EXENAME) $(MAINPACKAGE)

get :
	@echo "download packages...."
	@go get github.com/gorilla/mux

clean :
	@echo "cleaning...."
	@rm -rf $(BUILDPATH)/bin/$(EXENAME)
	@rm -rf $(BUILDPATH)/pkg
	@rm -rf $(BUILDPATH)/bin

test :
	go test -race -coverprofile=coverage.txt -covermode=atomic -short ./src/app ./src/movie

cover :
	go test -v -coverprofile=app.out ./src/app ./src/movie
	go tool cover -html=app.out