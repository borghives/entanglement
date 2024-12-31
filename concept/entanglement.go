package concept

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/borghives/websession"
)

type Entanglement struct {
	SystemSession websession.Session
	Frame         string
	Nonce         string
	Token         string
	Properties    map[string]string
}

func (e Entanglement) CalculatePropertiesState() string {
	representativeState := [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for key, value := range e.Properties {
		propState := sha256.Sum256([]byte(key + " : " + value))

		for i := range representativeState {
			representativeState[i] = representativeState[i] ^ propState[i]
		}
	}
	return string(hex.EncodeToString(representativeState[:]))
}

func (e *Entanglement) SetFrame(frame string) {
	e.Frame = frame
}

func (e *Entanglement) CreatSubFrame(frame string) Entanglement {
	return Entanglement{
		SystemSession: e.SystemSession,
		Frame:         frame,
		Nonce:         e.Nonce,
		Token:         e.Token,
		Properties:    nil,
	}
}

func (e *Entanglement) ResetProperty() {
	e.Properties = nil
}

func (e *Entanglement) SetProperty(key string, state string) {
	key = replaceChar(key, ':', '_')
	state = replaceChar(state, ':', '_')

	if e.Properties == nil {
		e.Properties = make(map[string]string)
	}

	e.Properties[key] = state
}

func (e Entanglement) GenerateToken() string {
	salt := websession.GenerateSalt(e.Nonce, e.Frame)
	return e.SystemSession.GenerateTokenFromSalt(salt)
}

func (e Entanglement) GenerateCorrelation(property string) string {
	salt := websession.GenerateSalt(e.Frame, property)

	if len(e.Properties) > 0 {
		salt += e.CalculatePropertiesState()
	}
	log.Println("salt", salt)
	return e.SystemSession.GenerateTokenFromSalt(salt)
}

func (e Entanglement) CheckToken() error {
	if e.Nonce == "" || e.Frame == "" || e.Token == "" {
		return fmt.Errorf("Empty nonce, frame or token")
	}

	gentoken := e.GenerateToken()

	if gentoken != e.Token {
		return fmt.Errorf("token mismatch")
	}

	return nil
}

func xorBytes(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("byte arrays must have the same length")
	}
	result := make([]byte, len(a))
	for i := range a {
		result[i] = a[i] ^ b[i]
	}
	return result
}

func replaceChar(s string, oldChar, newChar rune) string {
	// Use strings.Map for efficient rune-by-rune replacement
	return strings.Map(func(r rune) rune {
		if r == oldChar {
			return newChar
		}
		return r
	}, s)
}
