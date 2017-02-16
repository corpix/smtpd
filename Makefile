.DEFAULT_GOAL  = all
NAME           = smtpd
PACKAGE        = github.com/corpix/$(NAME)
NUMCPUS        = $(shell cat /proc/cpuinfo | grep '^processor\s*:' | wc -l)
VERSION        = $(shell git rev-list --count HEAD).$(shell git rev-parse --short HEAD)
LDFLAGS        = -X $(PACKAGE)/cmd.version=$(VERSION) \
                 -B 0x$(shell dd if=/dev/urandom bs=40 count=1 | sha1sum | awk '{print $$1}')

all: tools

$(NAME):
	govendor remove +u
	govendor add +e
	govendor sync
	mkdir -p build
	go build -a -ldflags "$(LDFLAGS)" -v \
		-o build/$(NAME)             \
		$(PACKAGE)/$(NAME)


test: tools
	go test -v $(PACKAGE)/...

lint: tools
	go vet $(PACKAGE)/...
	gometalinter                     \
		--deadline=5m            \
		--concurrency=$(NUMCPUS) \
		--exclude="(^|/)vendor/" \
		$(PACKAGE)/...

check: lint test

tools:
	if [ ! -e "$(GOPATH)"/src/"github.com/kardianos/govendor" ]; then go get github.com/kardianos/govendor; fi
	if [ ! -e "$(GOPATH)"/src/"github.com/rogpeppe/godef" ]; then go get github.com/rogpeppe/godef; fi
	if [ ! -e "$(GOPATH)"/src/"github.com/nsf/gocode" ]; then go get github.com/nsf/gocode; fi
	if [ ! -e "$(GOPATH)"/src/"github.com/stretchr/testify/assert" ]; then go get github.com/stretchr/testify/assert; fi
	if [ ! -e "$(GOPATH)"/src/"github.com/alecthomas/gometalinter" ]; then go get github.com/alecthomas/gometalinter && gometalinter --install; fi

.PHONY: all $(NAME) test lint check tools
