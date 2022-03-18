.PHONY: all
all: build

MADDY ?= $(shell command -v maddy.cover 2> /dev/null)
TLS_DIR ?= $(shell pwd)/test/tls/
TLS_CERT_FILE ?= $(TLS_DIR)ensmail.test.pem
TLS_KEY_FILE ?= $(TLS_DIR)ensmail.test-key.pem

.PHONY: init
init:
ifndef MADDY
	$(error "maddy.cover is not available please install https://github.com/foxcpp/maddy/blob/v0.5.4/tests/build_cover.sh")
endif
ifeq (,$(wildcard $(TLS_CERT_FILE)))
	@mkdir -p $(TLS_DIR)
	mkcert -cert-file="$(TLS_CERT_FILE)" -key-file="$(TLS_KEY_FILE)" ensmail.test mx.ensmail.test 127.0.0.1 localhost 
endif

.PHONY: test
test:
	@go test ./pkg/...

test-full: test init
	@go test ./test -integration.executable="$(MADDY)" -cert="$(TLS_CERT_FILE)" -key="$(TLS_KEY_FILE)"

.PHONY: build
build:
	@go build -ldflags \
		"-X main.version=${shell git describe --always --dirty --tags}" \
		-o build/ensmail \
		./cmd/ensmail.go

install: build
	export GOBIN=/usr/bin; go install github.com/foxcpp/maddy/cmd/maddy@latest
	cp ./build/ensmail /usr/bin
	-useradd ensmail -U -M -s /sbin/nologin
	-mkdir /run/ensmail /etc/ensmail
	chown -R ensmail:ensmail /run/ensmail
	@echo HTTP_WEB3_PROVIDER=$(HTTP_WEB3_PROVIDER) > /etc/ensmail/web3.env
	cp ./init/* /etc/systemd/system/
	systemctl daemon-reload
	systemctl enable ensmail
