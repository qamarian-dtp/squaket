package squaket

import (
	"gopkg.in/qamarian-lib/str.v2"
	"testing"
)

// TestSquaket () tests the Squaket data type.
func TestSquaket (t *testing.T) {
	// Testing function New (). ... {
	sA, errA := New ([]interface {} {false, "1", 2})
	if sA != nil {
		str.PrintEtr ("Test failed! Ref: 0", "err", "TestSquaket ()")
		t.FailNow ()
	}
	if errA == nil {
		str.PrintEtr ("Test failed! Ref: 1", "err", "TestSquaket ()")
		t.FailNow ()
	}
	// ... }

	// Testing method Group (). Confirming it does not invalid property name. ... {
	e0 := testType1 {"Ibrahim Qamardeen", 22, "Oyo"}
	e1 := testType2 {"Qamarian", "IBZ", "Oyo"}
	e2 := testType3 {"Birth", "22/04", "Lag"}

	sB, errB := New ([]interface {} {e0, e1, e2})
	if errB != nil {
		str.PrintEtr ("Test failed! Ref: 2", "err", "TestSquaket ()")
		t.FailNow ()
	}

	sC, errC := sB.Group ("_location")
	sD, errD := sB.Group ("location")
	sE, errE := sB.Group ("")
	sF, errF := sB.Group ("%%")
	sG, errG := sB.Group ("$d")

	if sC != nil || sD != nil || sE != nil || sF != nil || sG != nil || errC == nil ||
		errD == nil || errE == nil || errF == nil || errG == nil {

		str.PrintEtr ("Test failed! Ref: 3", "err", "TestSquaket ()")
		t.FailNow ()
	}
	// ... }

	// Testing method Group (). Confirming it does now allow grouping with uncommon
	// property. ... {
	e3 := testType1 {"Ibrahim Qamardeen", 22, "Oyo"}
	e4 := testType2 {"Qamarian", "IBZ", "Oyo"}
	e5 := testType3 {"Birth", "22/04", "Lag"}

	sH, errH := New ([]interface {} {e3, e4, e5})
	if errH != nil {
		str.PrintEtr ("Test failed! Ref: 4", "err", "TestSquaket ()")
		t.FailNow ()
	}

	resultI, errI := sH.Group ("Name")
	if resultI != nil {
		str.PrintEtr ("Test failed! Ref: 5", "err", "TestSquaket ()")
		t.FailNow ()
	}
	if errI == nil {
		str.PrintEtr ("Test failed! Ref: 6", "err", "TestSquaket ()")
		t.FailNow ()
	}
	// ... }

	// Testing method Group (). Confirming it does an accurate grouping. ... {
	e6 := testType1 {"Ibrahim Qamardeen", 22, "Oyo"}
	e7 := testType2 {"Qamarian", "IBZ", "Oyo"}
	e8 := testType3 {"Birth", "22/04", "Lag"}

	sJ, errJ := New ([]interface {} {e6, e7, e8})
	if errJ != nil {
		str.PrintEtr ("Test failed! Ref: 7", "err", "TestSquaket ()")
		t.FailNow ()
	}

	resultK, errK := sJ.Group ("Location")
	if errK != nil {
		str.PrintEtr ("Test failed! Ref: 8", "err", "TestSquaket ()")
		t.FailNow ()
	}

	if !((resultK ["Oyo"][0] == e6 && resultK ["Oyo"][1] == e7 &&
		resultK ["Lag"][0] == e8) ||

		(resultK ["Oyo"][0] == e7 && resultK ["Oyo"][1] == e6 &&
		resultK ["Lag"][0] == e8)) {

		str.PrintEtr ("Test failed! Ref: 9", "err", "TestSquaket ()")
		t.FailNow ()
	}
	// ... }

	str.PrintEtr ("Test passed.", "std", "TestSquaket ()")
}

type testType1 struct {
	Name     string
	Age      int
	Location string
}

type testType2 struct {
	Name     string
	Ceo      string
	Location string
}

type testType3 struct {
	Theme    string
	Account  string
	Location string
}
