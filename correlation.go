package entanglement

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type StateCorrelation map[string]string               //From one state (ID) relating to the next state (ID)
type TypeStateCorrelation map[string]StateCorrelation //Entity Type and its states correlation

func (e TypeStateCorrelation) AddCorrelation(frameName string, originState string, nextState string) {
	properties := e[frameName]
	if properties == nil {
		properties = make(StateCorrelation)
	}

	properties[originState] = nextState

	e[frameName] = properties
}

type EntangleProperties struct {
	Token        string               `xml:"-" json:"Token" bson:"-" `
	Correlations TypeStateCorrelation `xml:"-" json:"Correlations,omitempty" bson:"-" `
}

func (e *EntangleProperties) SetCorrelationProperties(name string, properties StateCorrelation) {
	if e.Correlations == nil {
		e.Correlations = make(TypeStateCorrelation)
	}

	e.Correlations[name] = properties
}

func (e *EntangleProperties) UpdateCorrelationProperties(typeCorrelation TypeStateCorrelation) {
	for key, value := range typeCorrelation {
		e.SetCorrelationProperties(key, value)
	}
}

type Correlatable interface {
	GetID() bson.ObjectID
	TransitionStates(frame Session) TypeStateCorrelation
	CheckExpectation(frame Session) error
}
