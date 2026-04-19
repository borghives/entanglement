package entanglement

import "github.com/borghives/websession"

type Session struct {
	WebFrame
	Session websession.Session
}

func NewSession(session websession.Session) Session {
	return Session{
		WebFrame: WebFrame{
			Nonce: websession.GetRandomHexString(),
			Token: session.GenerateSessionToken(),
		},
		Session: session,
	}
}

func EntangleSession(web WebFrame, session websession.Session) Session {
	return Session{
		WebFrame: web,
		Session:  session,
	}
}

func (e Session) CreateSubFrame(frame string) Session {
	return Session{
		WebFrame: e.WebFrame.CreateSubFrame(frame),
		Session:  e.Session,
	}
}

func (e Session) GenerateToken() string {
	return e.WebFrame.GenerateToken(e.Session)
}

func (e Session) GenerateCorrelation(property string) string {
	salt := websession.GenerateSalt(e.WebFrame.Frame, property)

	salt += e.WebFrame.CalculatePropertiesState()
	return e.Session.GenerateTokenFromSalt(salt)
}
