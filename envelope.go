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
