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

func TestMultislashdate(t *testing.T) {
	d := `109／108／106／105 BC`
	fp := TakeFingerprint(d)
	fp = multislashdate(fp)
	fp.PrintWithCalc()
}

func TestMulticommadate(t *testing.T) {
	d := `114, 116, ﹠ 156 ac`
	fp := TakeFingerprint(d)
	fp = multicommadate(fp)
	fp.PrintWithCalc()
}

func TestAndsigndate(t *testing.T) {
	d := `1299 ﹠ 1344 ac`
	fp := TakeFingerprint(d)
	fp = andsigndate(fp)
	fp.PrintWithCalc()
}

func TestPickandrunparser(t *testing.T) {
	d := `80-58／55-51 bc?`
	// d := `109／108／106／105 BC`
	fp := TakeFingerprint(d)
	fp = pickandrunparser(fp)
	fp.PrintWithCalc()
}
