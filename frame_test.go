package entanglement

import (
	"strings"
	"testing"

	"github.com/borghives/websession"
)

func TestSystemFrame_Create(t *testing.T) {
	sf := Create("my_nonce", "my_token")
	if sf.Nonce != "my_nonce" {
		t.Errorf("expected my_nonce, got %s", sf.Nonce)
	}
	if sf.Token != "my_token" {
		t.Errorf("expected my_token, got %s", sf.Token)
	}
}

func TestSystemFrame_CreateSubFrame(t *testing.T) {
	sf := Create("my_nonce", "my_token")
	sub := sf.CreateSubFrame("sub_frame")
	if sub.Name != "sub_frame" {
		t.Errorf("expected sub_frame, got %s", sub.Name)
	}
	if sub.Nonce != "my_nonce" || sub.Token != "my_token" {
		t.Errorf("unexpected values in sub frame")
	}
}

func TestSystemFrame_EntangleProperty(t *testing.T) {
	sf := Create("nonce", "token")
	sf.EntangleProperty("key:1", "val:1")

	if len(sf.Properties) != 1 {
		t.Fatalf("expected 1 property")
	}
	if sf.Properties["key_1"] != "val_1" {
		t.Errorf("expected val_1, got %s", sf.Properties["key_1"])
	}
}

func TestSystemFrame_SetFrame(t *testing.T) {
	sf := Create("nonce", "token")
	sf.SetFrame("new_frame")
	if sf.Name != "new_frame" {
		t.Errorf("expected new_frame, got %s", sf.Name)
	}
}

func TestSystemFrame_StateString(t *testing.T) {
	sf := Create("nonce", "token")
	sf.SetFrame("frame1")
	sf.EntangleProperty("k1", "v1")

	str := sf.StateString()
	if !strings.Contains(str, "Name: frame1") {
		t.Errorf("StateString missing name: %s", str)
	}
	if !strings.Contains(str, "Properties: (k1:v1,)") {
		t.Errorf("StateString missing properties: %s", str)
	}
}

func TestSystemFrame_CalculateEntangledState(t *testing.T) {
	sf := Create("nonce", "token")
	res := sf.CalculateEntangledState()
	if res != "" {
		t.Errorf("expected empty string for no properties")
	}

	sf.EntangleProperty("k1", "v1")
	res1 := sf.CalculateEntangledState()
	if res1 == "" {
		t.Errorf("expected non-empty state")
	}
}

func TestSystemFrame_TokenVerification(t *testing.T) {
	session := websession.NewWebSession("127.0.0.1", "ua")
	
	sf := Create("", "")
	err := sf.VerifyTokenAlignment(*session)
	if err == nil || err.Error() != "Empty nonce, frame or token" {
		t.Errorf("expected empty token error, got %v", err)
	}

	sf.Name = "test_frame"
	sf.Nonce = "test_nonce"
	token := sf.GenerateToken(*session)
	sf.Token = token

	err = sf.VerifyTokenAlignment(*session)
	if err != nil {
		t.Errorf("expected valid token alignment, got %v", err)
	}

	sf.Token = "tampered"
	err = sf.VerifyTokenAlignment(*session)
	if err == nil || err.Error() != "token mismatch" {
		t.Errorf("expected token mismatch error, got %v", err)
	}
}

func TestReplaceChar(t *testing.T) {
	res := replaceChar("a:b:c", ':', '_')
	if res != "a_b_c" {
		t.Errorf("expected a_b_c, got %s", res)
	}
}
