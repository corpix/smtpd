.DEFAULT_GOAL  = all

build_id      := 0x$(shell dd if=/dev/urandom bs=40 count=1 2> /dev/null | sha1sum | awk '{print $$1}')

NAME           = smtpd
PACKAGE        = github.com/corpix/$(NAME)
NUMCPUS        = $(shell cat /proc/cpuinfo | grep '^processor\s*:' | wc -l)
VERSION        = $(shell git rev-list --count HEAD).$(shell git rev-parse --short HEAD)
LDFLAGS        = -X $(PACKAGE)/cli.version=$(VERSION) \
                 -B $(build_id)


.PHONY: all
all: tools

.PHONY: $(NAME)
$(NAME):
	govendor remove +u
	govendor add +e
	govendor sync
	mkdir -p build
	@echo "Build id: $(build_id)"
	go build -a -ldflags "$(LDFLAGS)" -v \
		-o build/$(NAME)             \
		$(PACKAGE)/$(NAME)

.PHONY: build
build: $(NAME)


.PHONY: test
test: tools
	go test -v ./...

.PHONY: bench
bench: tools
	go test -bench=. -v ./...

.PHONY: lint
lint: tools
	go vet ./...
	gometalinter                     \
		--deadline=5m            \
		--concurrency=$(NUMCPUS) \
		--exclude="(^|/)vendor/" \
		./...

.PHONY: check
check: lint test

.PHON: tools
tools:
	if [ ! -e "$(GOPATH)"/src/"github.com/kardianos/govendor" ]; then go get github.com/kardianos/govendor; fi
	if [ ! -e "$(GOPATH)"/src/"github.com/rogpeppe/godef" ]; then go get github.com/rogpeppe/godef; fi
	if [ ! -e "$(GOPATH)"/src/"github.com/nsf/gocode" ]; then go get github.com/nsf/gocode; fi
	if [ ! -e "$(GOPATH)"/src/"github.com/stretchr/testify/assert" ]; then go get github.com/stretchr/testify/assert; fi
	if [ ! -e "$(GOPATH)"/src/"github.com/alecthomas/gometalinter" ]; then go get github.com/alecthomas/gometalinter && gometalinter --install; fi
