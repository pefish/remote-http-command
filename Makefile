
DEFAULT: build-cur

ifeq ($(GOPATH),)
  GOPATH = $(HOME)/go
endif

build-cur:
	GOPATH=$(GOPATH) go install github.com/pefish/go-build-tool/cmd/...@latest
	$(GOPATH)/bin/go-build-tool

install: build-cur
	sudo install -C ./build/bin/linux/remote-http-command /usr/local/bin/remote-http-command

install-service: install
	sudo mkdir -p /etc/systemd/system
	sudo install -C -m 0644 ./script/remote-http-command.service /etc/systemd/system/remote-http-command.service
	sudo systemctl daemon-reload
	@echo
	@echo "remote-http-command service installed."

