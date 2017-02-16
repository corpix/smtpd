// The MIT License (MIT)

// Copyright © 2017 Dmitry Moskowski

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
package smtpd

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"net/mail"
)

var (
	// EnvelopeHashDelimiter is a delimiter that should be used
	// while calculating Hash.
	EnvelopeHashDelimiter = []byte{0}
)

// Envelope is a wrapper around general message fields and data.
type Envelope struct {
	From string
	To   []string
	Data []byte
}

// Message decodes Envelope.Data with mail.ReadMessage() and
// returns the result.
func (e *Envelope) Message() (*mail.Message, error) {
	return mail.ReadMessage(bytes.NewBuffer(e.Data))
}

// Hash calculates sha256 hash sum of the envelope contents.
func (e *Envelope) Hash() []byte {
	hash := sha256.New()
	hash.Write([]byte(e.From))
	hash.Write(EnvelopeHashDelimiter)

	for _, v := range e.To {
		hash.Write([]byte(v))
		hash.Write(EnvelopeHashDelimiter)
	}

	hash.Write(e.Data)

	return hash.Sum(nil)
}

// HashString is a Hash variant that encodes a result with hex.
func (e *Envelope) HashString() string {
	return hex.EncodeToString(e.Hash())
}
