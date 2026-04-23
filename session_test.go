package entanglement

import (
	"testing"

	"github.com/borghives/websession"
)

func TestNewSession(t *testing.T) {
	ws := websession.NewWebSession("127.0.0.1", "ua")
	session := NewSession(*ws)

	if session.SystemFrame.Nonce == "" {
		t.Errorf("expected non-empty nonce")
	}
	if session.SystemFrame.Token == "" {
		t.Errorf("expected non-empty token")
	}
}

func TestEntangleSession(t *testing.T) {
	ws := websession.NewWebSession("127.0.0.1", "ua")
	frame := Create("n", "t")
	session := EntangleSession(frame, *ws)

	if session.SystemFrame.Nonce != "n" {
		t.Errorf("expected nonce n")
	}
}

func TestSession_CreateSubFrame(t *testing.T) {
	ws := websession.NewWebSession("127.0.0.1", "ua")
	session := NewSession(*ws)
	sub := session.CreateSubFrame("sub1")

	if sub.SystemFrame.Name != "sub1" {
		t.Errorf("expected name sub1")
	}
}

func TestSession_GenerateToken(t *testing.T) {
	ws := websession.NewWebSession("127.0.0.1", "ua")
	session := NewSession(*ws)
	session.SystemFrame.Name = "test"
	
	tok := session.GenerateToken()
	if tok == "" {
		t.Errorf("expected generated token")
	}
}

func TestSession_GenerateCorrelation(t *testing.T) {
	ws := websession.NewWebSession("127.0.0.1", "ua")
	session := NewSession(*ws)
	session.SystemFrame.Name = "test"
	
	corr := session.GenerateCorrelation("prop1")
	if corr == "" {
		t.Errorf("expected generated correlation")
	}
}
