
.EXPORT_ALL_VARIABLES:
VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
# PROJECTNAME := $(shell basename "$(PWD)")
PROJECTNAME := echov6

LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"
STDERR := /tmp/.$(PROJECTNAME)-stderr.txt
export CGO_ENABLED=0
# If the first argument is "run"...
ifeq (cert,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

.PHONY: web
web:
	go run main.go web

.PHONY: build
build: 
	go build $(LDFLAGS) -o dist/$(PROJECTNAME)
	chmod +x dist/$(PROJECTNAME)

.PHONY: docker
docker: 
	docker build -t nuxion/${PROJECTNAME} .


.PHONY: docker-run
docker-run:
	docker run --rm -p 5656:5656 nuxion/${PROJECTNAME} 

.PHONY: test
test:
	curl -6 "[::1]:5656/v1/get" 

.PHONY: release
release: docker
	docker tag nuxion/${PROJECTNAME} registry.int.deskcrash.com/nuxion/${PROJECTNAME}:$(VERSION)
	docker push registry.int.deskcrash.com/nuxion/$(PROJECTNAME):$(VERSION)

