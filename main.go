package main

import (
	"bufio"
	"fmt"
	"github.com/e-gun/dateconversion/structs"
	"log"
	"os"
	"strings"
)

func main() {

	//td := strings.Split(TESTDATA, "\n")
	//for _, line := range td {
	//	fp := TakeFingerprint(line)
	//	fp = pickandrunparser(fp)
	//	fp.PrintWithCalc()
	//}
	testcoverage()
}

func testcoverage() {
	const (
		DISTINCT = `work_distinct_dates.txt`
	)
	f, err := os.Open(DISTINCT)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	var prints []structs.FingerPrint
	for scanner.Scan() {
		// do something with a line
		txt := scanner.Text()
		fp := TakeFingerprint(strings.TrimSpace(txt))
		prints = append(prints, fp)
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i, fp := range prints {
		prints[i] = pickandrunparser(fp)
	}

	badcount := 0
	lacking := 0
	var failures []structs.FingerPrint
	for _, fp := range prints {
		if fp.LacksParser {
			lacking++
		}
		if fp.ParserFailed {
			badcount++
			failures = append(failures, fp)
		}
	}

	fmt.Println("lack parser:", lacking)
	fmt.Printf("parser failures: %d of %d\n", badcount, len(prints))
	fmt.Println("sample failurse")
	for i, fp := range failures {
		if i%25 == 0 {
			fp.PrintWithCalc()
		}
	}
}
