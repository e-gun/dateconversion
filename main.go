package main

import (
	"strings"
)

func main() {

	td := strings.Split(TESTDATA, "\n")
	for _, line := range td {
		fp := TakeFingerprint(line)
		fp.Rationalize()
		fp = pickbasicparser(fp)
		fp.PrintWithCalc()
	}
}
