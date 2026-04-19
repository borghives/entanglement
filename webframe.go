package entanglement

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/borghives/websession"
)

type WebFrame struct {
	Frame      string
	Nonce      string
	Token      string
	Properties map[string]string
}

func CreateWeb(nonce string, token string) WebFrame {
	return WebFrame{
		Nonce: nonce,
		Token: token,
	}
}

func (e WebFrame) CreateSubFrame(frame string) WebFrame {
	return WebFrame{
		Frame: frame,
		Nonce: e.Nonce,
		Token: e.Token,
	}
}

func (e WebFrame) GenerateToken(session websession.Session) string {
	salt := websession.GenerateSalt(e.Nonce, e.Frame)
	return session.GenerateTokenFromSalt(salt)
}

func (e WebFrame) VerifyTokenAlignment(session websession.Session) error {
	if e.Nonce == "" || e.Frame == "" || e.Token == "" {
		return fmt.Errorf("Empty nonce, frame or token")
	}

	gentoken := e.GenerateToken(session)

	if gentoken != e.Token {
		return fmt.Errorf("token mismatch")
	}

	return nil
}

func (e WebFrame) CalculatePropertiesState() string {
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

func (e *WebFrame) SetProperty(key string, state string) *WebFrame {
	key = replaceChar(key, ':', '_')
	state = replaceChar(state, ':', '_')

	if e.Properties == nil {
		e.Properties = make(map[string]string)
	}

	e.Properties[key] = state
	return e
}

func (e *WebFrame) SetFrame(frame string) *WebFrame {
	e.Frame = frame
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
