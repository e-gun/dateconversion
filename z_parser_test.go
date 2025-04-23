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
