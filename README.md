smtpd
-----

[![Build Status](https://travis-ci.org/corpix/smtpd.svg?branch=master&12312556)](https://travis-ci.org/corpix/smtpd)

This package is trying to implement RFC-compatible SMTP server fo go.

## Example

### Requirements

* `swaks` to test our SMTP server

### Starting

> See `examples/`.

Save this code to `example.go`(same code could be found in `examples/logging-smtpd/main.go`):

``` go
package main

import (
	"io/ioutil"
	"net"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/davecgh/go-spew/spew"

	"github.com/corpix/smtpd"
)

type smtpServer struct{}

func (s *smtpServer) ServeSMTP(c net.Conn, e *smtpd.Envelope) {
	logrus.Infof(
		"Received message from %s those envelope: %s\n",
		c.RemoteAddr(),
		spew.Sdump(e),
	)
	msg, err := e.Message()
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(msg.Body)
	if err != nil {
		panic(err)
	}
	logrus.Infof("Message body: %s\n", body)
}

func main() {
	var err error

	c, err := net.Listen("tcp", "127.0.0.1:2525")
	if err != nil {
		panic(err)
	}

	for {
		err = smtpd.Serve(c, &smtpServer{})
		if err != nil {
			logrus.Error(err)
			time.Sleep(1 * time.Second)
		}
	}
}

```

And run the server:

``` shell
go run example.go
```

It will listen on `127.0.0.1:2525`.

### Send mail

``` shell
swaks -f me@example.org -t test@example.com --server 127.0.0.1:2525
```

You should see output similar to this:

``` text
# remote 127.0.0.1:55912 at 2017-02-16 15:42:56 +0000
w 220 localhost go-smtpd
r EHLO localhost
w 250-localhost Hello 127.0.0.1:55912
w 250-8BITMIME
w 250-PIPELINING
w 250 HELP
r MAIL FROM:<me@example.org>
w 250 Okay, I'll believe you for now
r RCPT TO:<test@example.com>
w 250 Okay, I'll believe you for now
r DATA
w 354 Send away
r . <end of data>
w 250 I've put it in a can
r QUIT
w 221 Goodbye
# finished at 2017-02-16 15:42:56 +0000
INFO[0004] Received message from 127.0.0.1:55912 those envelope: (*smtpd.Envelope)(0xc4200d2000)({
 From: (string) (len=14) "me@example.org",
 To: ([]string) (len=1 cap=1) {
  (string) (len=16) "test@example.com"
 },
 Data: ([]uint8) (len=208 cap=208) {
  00000000  44 61 74 65 3a 20 54 68  75 2c 20 31 36 20 46 65  |Date: Thu, 16 Fe|
  00000010  62 20 32 30 31 37 20 31  35 3a 34 32 3a 35 36 20  |b 2017 15:42:56 |
  00000020  2b 30 30 30 30 0a 54 6f  3a 20 74 65 73 74 40 65  |+0000.To: test@e|
  00000030  78 61 6d 70 6c 65 2e 63  6f 6d 0a 46 72 6f 6d 3a  |xample.com.From:|
  00000040  20 6d 65 40 65 78 61 6d  70 6c 65 2e 6f 72 67 0a  | me@example.org.|
  00000050  53 75 62 6a 65 63 74 3a  20 74 65 73 74 20 54 68  |Subject: test Th|
  00000060  75 2c 20 31 36 20 46 65  62 20 32 30 31 37 20 31  |u, 16 Feb 2017 1|
  00000070  35 3a 34 32 3a 35 36 20  2b 30 30 30 30 0a 58 2d  |5:42:56 +0000.X-|
  00000080  4d 61 69 6c 65 72 3a 20  73 77 61 6b 73 20 76 32  |Mailer: swaks v2|
  00000090  30 31 33 30 32 30 39 2e  30 20 6a 65 74 6d 6f 72  |0130209.0 jetmor|
  000000a0  65 2e 6f 72 67 2f 6a 6f  68 6e 2f 63 6f 64 65 2f  |e.org/john/code/|
  000000b0  73 77 61 6b 73 2f 0a 0a  54 68 69 73 20 69 73 20  |swaks/..This is |
  000000c0  61 20 74 65 73 74 20 6d  61 69 6c 69 6e 67 0a 0a  |a test mailing..|
 }
})


INFO[0004] Message body: This is a test mailing
```

## Credits

* [@siebenmann](https://github.com/siebenmann/smtpd) for awesome work on SMTPD package this packaged is based on.

## License

MIT
