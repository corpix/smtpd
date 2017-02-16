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
