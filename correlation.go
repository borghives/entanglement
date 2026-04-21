package entanglement

import (
	"maps"
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

func (e TypeStateCorrelation) Update(source TypeStateCorrelation) {
	maps.Copy(e, source)
}

type EntangleProperties struct {
	Token        string               `xml:"-" json:"Token" bson:"-" `
	Correlations TypeStateCorrelation `xml:"-" json:"Correlations,omitempty" bson:"-" `
}

func (e *EntangleProperties) UpdateCorrelationProperties(typeCorrelation TypeStateCorrelation) {
	if e.Correlations == nil {
		e.Correlations = make(TypeStateCorrelation)
	}

	e.Correlations.Update(typeCorrelation)
}

type Correlatable interface {
	TransitionStates(frame Session) TypeStateCorrelation
	CheckTransition(frame Session) error
}
