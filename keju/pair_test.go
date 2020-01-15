package keju

import "testing"

func TestQAPair_HashCode(t *testing.T) {
	p1 := QAPair{
		Question: "Question",
		Answer: "Answer",
	}
	p2 := QAPair{
		Question: "Question",
		Answer: "Answer",
	}
	p3 := QAPair{
		Question: "Question2",
		Answer: "Answer2",
	}
	p4 := QAPair{
		Question: "",
		Answer: "",
	}
	if p1.HashCode() != p1.HashCode() {
		t.Error("Case 1",  p1.HashCode(), p1.HashCode())
		t.Fail()
	}
	if p1.HashCode() != p2.HashCode() {
		t.Error("Case 2")
		t.Fail()
	}
	if p1.HashCode() == p3.HashCode() {
		t.Error("Case 3")
		t.Fail()
	}
	if p1.HashCode() == p4.HashCode() {
		t.Error("Case 4")
		t.Fail()
	}
	if p4.HashCode() != p4.HashCode() {
		t.Error("Case 5")
		t.Fail()
	}
}
