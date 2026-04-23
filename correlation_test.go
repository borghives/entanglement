package entanglement

import (
	"reflect"
	"testing"
)

func TestTypeStateCorrelation_AddCorrelation(t *testing.T) {
	tc := make(TypeStateCorrelation)
	tc.AddCorrelation("frame1", "origin", "next")

	if len(tc) != 1 {
		t.Fatalf("expected length 1, got %d", len(tc))
	}
	if tc["frame1"]["origin"] != "next" {
		t.Errorf("expected 'next', got '%s'", tc["frame1"]["origin"])
	}

	tc.AddCorrelation("frame1", "origin2", "next2")
	if len(tc["frame1"]) != 2 {
		t.Errorf("expected 2 correlations in frame1, got %d", len(tc["frame1"]))
	}
}

func TestTypeStateCorrelation_Update(t *testing.T) {
	tc1 := make(TypeStateCorrelation)
	tc1.AddCorrelation("f1", "s1", "s2")

	tc2 := make(TypeStateCorrelation)
	tc2.AddCorrelation("f2", "s3", "s4")
	tc2.AddCorrelation("f1", "s5", "s6")

	tc1.Update(tc2)

	if len(tc1) != 2 {
		t.Fatalf("expected length 2, got %d", len(tc1))
	}
	if !reflect.DeepEqual(tc1["f2"], StateCorrelation{"s3": "s4"}) {
		t.Errorf("unexpected f2: %v", tc1["f2"])
	}
	if !reflect.DeepEqual(tc1["f1"], StateCorrelation{"s5": "s6"}) {
		t.Errorf("unexpected f1: %v", tc1["f1"])
	}
}

func TestEntangleProperties_UpdateCorrelationProperties(t *testing.T) {
	ep := &EntangleProperties{}
	tc := make(TypeStateCorrelation)
	tc.AddCorrelation("f1", "s1", "s2")

	ep.UpdateCorrelationProperties(tc)

	if len(ep.Correlations) != 1 {
		t.Fatalf("expected correlations length 1, got %d", len(ep.Correlations))
	}
	if ep.Correlations["f1"]["s1"] != "s2" {
		t.Errorf("unexpected value in correlation")
	}
}
