package entanglement

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/borghives/websession"
)

type SystemFrame struct {
	Name       string
	Nonce      string
	Token      string
	Properties map[string]string
}

func Create(nonce string, token string) SystemFrame {
	return SystemFrame{
		Nonce: nonce,
		Token: token,
	}
}

func (e SystemFrame) CreateSubFrame(frame string) SystemFrame {
	return SystemFrame{
		Name:  frame,
		Nonce: e.Nonce,
		Token: e.Token,
	}
}

func (e SystemFrame) GenerateToken(session websession.Session) string {
	salt := websession.GenerateSalt(e.Nonce, e.Name)
	return session.GenerateTokenFromSalt(salt)
}

func (e SystemFrame) VerifyTokenAlignment(session websession.Session) error {
	if e.Nonce == "" || e.Name == "" || e.Token == "" {
		return fmt.Errorf("Empty nonce, frame or token")
	}

	gentoken := e.GenerateToken(session)

	if gentoken != e.Token {
		return fmt.Errorf("token mismatch")
	}

	return nil
}

func (e SystemFrame) CalculateEntangledState() string {
	if len(e.Properties) == 0 {
		return ""
	}

	representativeState := [32]byte{}
	for key, value := range e.Properties {
		propState := sha256.Sum256([]byte(key + " : " + value))

		for i := range representativeState {
			representativeState[i] = representativeState[i] ^ propState[i]
		}
	}
	return string(hex.EncodeToString(representativeState[:]))
}

func (e *SystemFrame) EntangleProperty(key string, state string) *SystemFrame {
	key = replaceChar(key, ':', '_')
	state = replaceChar(state, ':', '_')

	if e.Properties == nil {
		e.Properties = make(map[string]string)
	}

	e.Properties[key] = state
	return e
}

func (e *SystemFrame) SetFrame(frame string) *SystemFrame {
	e.Name = frame
	return e
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
