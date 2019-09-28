package squaket

import (
	"fmt"
	"gopkg.in/qamarian-lib/str.v2"
	"testing"
)

func TestSquaket (t *testing.T) {
	// Testing function New (). ... {
	invalidData := []interface {} {false, "1", 2}
	sA, errA := New (invalidData)
	if sA != nil || errA == nil {
		str.PrintEtr ("Test failed! Ref: 0", "err", "TestSquaket ()")
		t.FailNow ()
	}
	// ... }

	// Testing method Group (). Confirming it does now allow grouping with uncommon property. ... {
	e1 := testType1 {"Ibrahim Qamardeen", 22, "Oyo"}
	e2 := testType2 {"Qamarian", "IBZ", "Oyo"}
	e3 := testType3 {"Birth", "22/04", "Lag"}

	sB, errB := New ([]interface {} {e1, e2, e3})
	if errB != nil {
		str.PrintEtr ("Test failed! Ref: 1", "err", "TestSquaket ()")
		t.FailNow ()
	}

	resultC, errC := sB.Group ("Name")
	if resultC != nil {
		str.PrintEtr ("Test failed! Ref: 2", "err", "TestSquaket ()")
		t.FailNow ()
	}
	if errC == nil {
		str.PrintEtr ("Test failed! Ref: 3", "err", "TestSquaket ()")
		t.FailNow ()
	}
	// ... }

	// Testing method Group (). Confirming it does an accurate grouping. ... {
	e4 := testType1 {"Ibrahim Qamardeen", 22, "Oyo"}
	e5 := testType2 {"Qamarian", "IBZ", "Oyo"}
	e6 := testType3 {"Birth", "22/04", "Lag"}

	sD, errD := New ([]interface {} {e1, e2, e3})
	if errD != nil {
		str.PrintEtr ("Test failed! Ref: 4", "err", "TestSquaket ()")
		t.FailNow ()
	}

	resultE, errE := sD.Group ("Location")
	if errE != nil {
		str.PrintEtr ("Test failed! Ref: 5", "err", "TestSquaket ()")
		fmt.Println (errE)
		t.FailNow ()
	}

	if !((resultE ["Oyo"][0] == e4 && resultE ["Oyo"][1] == e5 && resultE ["Lag"][0] == e6) ||
		(resultE ["Oyo"][0] == e5 && resultE ["Oyo"][1] == e4 && resultE ["Lag"][0] == e6)) {
		str.PrintEtr ("Test failed! Ref: 3", "err", "TestSquaket ()")
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
