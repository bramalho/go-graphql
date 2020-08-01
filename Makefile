PID_FILE=/tmp/app.pid
GO_FILES=$(filter-out %_test.go, $(wildcard *.go))
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
BINARY_NAME=${APP_NAME:-app}
BINARY_UNIX=$(BINARY_NAME)_unix

start:
	$(GO_CMD) run $(GO_FILES) & echo $$! > $(PID_FILE)

stop:
	-kill `pstree -p \`cat $(PID_FILE)\` | tr "\n" " " |sed "s/[^0-9]/ /g" |sed "s/\s\s*/ /g"`

before:
	@echo "STOPED app" && printf '%*s\n' "40" '' | tr ' ' -

restart: stop before start
	@echo "STARTED app" && printf '%*s\n' "40" '' | tr ' ' -

serve: start
	fswatch -or --event=Updated /home/app | \
	xargs -n1 -I {} make restart

.PHONY: start before stop restart serve

# Build
all: clean build build-linux

clean:
	$(GO_CLEAN)
	rm -rf .build/

build:
	$(GO_BUILD) -o .build/$(BINARY_NAME) -v

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_BUILD) -o .build/$(BINARY_UNIX) -v
