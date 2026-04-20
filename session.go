package entanglement

import "github.com/borghives/websession"

type Session struct {
	SystemFrame
	Session websession.Session
}

func NewSession(session websession.Session) Session {
	return Session{
		SystemFrame: SystemFrame{
			Nonce: websession.GetRandomHexString(),
			Token: session.GenerateSessionToken(),
		},
		Session: session,
	}
}

func EntangleSession(web SystemFrame, session websession.Session) Session {
	return Session{
		SystemFrame: web,
		Session:     session,
	}
}

func (e Session) CreateSubFrame(frame string) Session {
	return Session{
		SystemFrame: e.SystemFrame.CreateSubFrame(frame),
		Session:     e.Session,
	}
}

func (e Session) GenerateToken() string {
	return e.SystemFrame.GenerateToken(e.Session)
}

func (e Session) GenerateCorrelation(property string) string {
	salt := websession.GenerateSalt(e.SystemFrame.Frame, property)

	salt += e.SystemFrame.CalculateEntangledState()
	return e.Session.GenerateTokenFromSalt(salt)
}
