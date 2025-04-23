package main

import "testing"

func TestTwoarabicsimple(t *testing.T) {
	d := `1011／1012 ac`
	fp := TakeFingerprint(d)
	fp = twoarabicsimple(fp)
	fp.PrintWithCalc()
}

func TestOnearabiccentury(t *testing.T) {
	d := `5th ad`
	fp := TakeFingerprint(d)
	fp = onearabiccentury(fp)
	fp.PrintWithCalc()
}

func TestOneromancentury(t *testing.T) {
	d := `mid-IV ac?`
	fp := TakeFingerprint(d)
	fp = oneromancentury(fp)
	fp.PrintWithCalc()
}

func TestSlashdated(t *testing.T) {
	d := `c.63／2-51／0 bc`
	fp := TakeFingerprint(d)
	fp = slashdated(fp)
	fp.PrintWithCalc()
}

func TestEitherordate(t *testing.T) {
	d := `618 or 633 ac`
	fp := TakeFingerprint(d)
	fp = eitherordate(fp)
	fp.PrintWithCalc()
}

func TestPickandrunparser(t *testing.T) {
	d := `c.10／9-3／2`
	fp := TakeFingerprint(d)
	fp = pickandrunparser(fp)
	fp.PrintWithCalc()
}
