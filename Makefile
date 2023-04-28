mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
PID = ./app.pid
GO_FILES = $(wildcard *.go)$(wildcard **/*.go)
APP = ./app
MAIN := $(dir $(mkfile_path))cmd

serve: start
	@fswatch -x -o --event Created --event Updated --event Renamed -r -e '.*' -i '\.go$$'  **/*.go | xargs -n1 -I{}  make restart || make kill

kill:
	@kill `cat $(PID)` || true

before:
	@echo "actually do nothing"

build: $(GO_FILES)
	@go build -o $(APP) $(MAIN)

$(APP): $(GO_FILES)
	@go build $? -o $@ $(MAIN)

start:
	# @sh -c "$(APP) & echo $$! > $(PID)"
	@./app & echo $$! > $(PID)

restart: kill before build start

.PHONY: start serve restart kill before # let's go to reserve rules names