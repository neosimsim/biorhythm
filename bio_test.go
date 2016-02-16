package main

import "math"
import "testing"

const eps = 1e-13

func TestBiorythmPhysical(t *testing.T) {
	p, _, _ := biorythm(9080)
	expectedP := -0.979084087682321
	if math.Abs(p-expectedP) > eps {
		t.Error("Unexpected physical:", p, "expected:", expectedP)
	}
}
func TestBiorythmEmotional(t *testing.T) {
	_, e, _ := biorythm(9080)
	expectedE := 0.974927912181876
	if math.Abs(e-expectedE) > eps {
		t.Error("Unexpected emotional:", e, "expected:", expectedE)
	}
}
func TestBiorythmIntelligence(t *testing.T) {
	_, _, i := biorythm(9080)
	expectedI := 0.81457595205029
	if math.Abs(i-expectedI) > eps {
		t.Error("Unexpected intelligence:", i, "expected:", expectedI)
	}
}
