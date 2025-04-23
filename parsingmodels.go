package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	stripalmostallstrings = regexp.MustCompile(`[^\d\-／\s]`)
	swapspanner           = regexp.MustCompile(`[\-／]`)
	stripforromans1       = regexp.MustCompile(`[^IVX\-／\s]`)
	stripforromans0       = regexp.MustCompile(`[^IVX\s]`)
)

// onearabicsimple: `101 bc` --> -101
func onearabicsimple(fp FingerPrint) FingerPrint {
	// fmt.Println("onearabicsimple", fp.OrigDateString)
	cleaned := stripalmostallstrings.ReplaceAllString(fp.OrigDateString, "")
	cleaned = strings.TrimSpace(cleaned)
	d, e := strconv.Atoi(cleaned)
	if e != nil {
		fmt.Printf("onearabicsimple parser failed '%s'\n", fp.OrigDateString)
		fp.ParserFailed = true
		return fp
	}
	fp.Calculated = d
	fp.ApplyBCE()
	fp.ApplySimpleFudges()
	// fmt.Println("onearabicsimple", fp.Calculated)
	return fp
}

// twoarabicsimple -  `1-50 ac` --> 25
func twoarabicsimple(fp FingerPrint) FingerPrint {
	cleaned := stripalmostallstrings.ReplaceAllString(fp.OrigDateString, "")
	cleaned = swapspanner.ReplaceAllString(cleaned, " ")
	cleaned = strings.ReplaceAll(cleaned, "  ", " ")
	halves := strings.Split(strings.TrimSpace(cleaned), " ")
	if len(halves) != 2 {
		fmt.Printf("twoarabicsimple parser failed ptA '%s'\n", fp.OrigDateString)
		fp.ParserFailed = true
		return fp
	}
	d1, e1 := strconv.Atoi(halves[0])
	d2, e2 := strconv.Atoi(halves[1])
	if e1 != nil || e2 != nil {
		fmt.Printf("twoarabicsimple parser failed ptB '%s'\n", fp.OrigDateString)
		fp.ParserFailed = true
		return fp
	}

	// `before 337／6 bc` does not involve the average of two numbers...
	if len(halves[0]) != 3 && len(halves[1]) != 1 {
		mid := (d1 + d2) / 2
		fp.Calculated = mid
	} else {
		fp.Calculated = d1
	}
	fp.ApplyBCE()
	fp.ApplySimpleFudges()
	return fp
}

// twoarabiccenturies -  `7th-9th ac` --> 750
func twoarabiccenturies(fp FingerPrint) FingerPrint {
	cleaned := stripalmostallstrings.ReplaceAllString(fp.OrigDateString, "")
	cleaned = swapspanner.ReplaceAllString(cleaned, " ")
	cleaned = strings.ReplaceAll(cleaned, "  ", " ")
	halves := strings.Split(strings.TrimSpace(cleaned), " ")
	// fmt.Println(halves)
	if len(halves) != 2 {
		fmt.Printf("twoarabiccenturies parser failed ptA '%s'\n", fp.OrigDateString)
		fp.ParserFailed = true
		return fp
	}
	d1, e1 := strconv.Atoi(halves[0])
	d2, e2 := strconv.Atoi(halves[1])
	if e1 != nil || e2 != nil {
		fmt.Printf("twoarabiccenturies parser failed ptB '%s'\n", fp.OrigDateString)
		fp.ParserFailed = true
		return fp
	}
	mid := (float32(d1) + float32(d2)) / 2
	mid = mid * 100
	fp.Calculated = int(mid)
	fp.ApplyBCE()
	if fp.Calculated < 0 {
		// 5th bce is not -500 but -450...
		fp.Calculated = fp.Calculated + 50
	} else {
		// 5th ce is not 500 but 450...
		fp.Calculated = fp.Calculated - 50
	}
	fp.ApplySimpleFudges()
	return fp
}

// onearabiccentury : `early 5th bc` --> -475
func onearabiccentury(fp FingerPrint) FingerPrint {
	cleaned := stripalmostallstrings.ReplaceAllString(fp.OrigDateString, "")
	cleaned = strings.TrimSpace(cleaned)
	d, e := strconv.Atoi(cleaned)
	if e != nil {
		fmt.Printf("onearabiccentury parser failed '%s'\n", fp.OrigDateString)
		fp.ParserFailed = true
		return fp
	}
	fp.Calculated = d * 100
	fp.ApplyBCE()
	if fp.Calculated < 0 {
		// 5th bce is not -500 but -450...
		fp.Calculated = fp.Calculated + 50
	} else {
		// 5th ce is not 500 but 450...
		fp.Calculated = fp.Calculated - 50
	}
	fp.ApplySimpleFudges()
	return fp
}

// twoaromancenturies -  `7th-9th ac` --> 800
func twoaromancenturies(fp FingerPrint) FingerPrint {
	cleaned := stripforromans1.ReplaceAllString(fp.OrigDateString, "")
	cleaned = swapspanner.ReplaceAllString(cleaned, " ")
	cleaned = strings.ReplaceAll(cleaned, "  ", " ")
	halves := strings.Split(strings.TrimSpace(cleaned), " ")
	// fmt.Println(halves)
	if len(halves) != 2 {
		fmt.Printf("twoaromancenturies parser failed ptA '%s' : %s\n", fp.OrigDateString, cleaned)
		fp.ParserFailed = true
		return fp
	}
	d1, e1 := romannumerals[halves[0]]
	d2, e2 := romannumerals[halves[1]]
	if e1 != true || e2 != true {
		fmt.Printf("twoaromancenturies parser failed ptB '%s' : %s\n", fp.OrigDateString, cleaned)
		fp.ParserFailed = true
		return fp
	}
	mid := (float32(d1) + float32(d2)) / 2
	mid = mid * 100
	fp.Calculated = int(mid)
	fp.ApplyBCE()
	if fp.Calculated < 0 {
		// 5th bce is not -500 but -450...
		fp.Calculated = fp.Calculated + 50
	} else {
		// 5th ce is not 500 but 450...
		fp.Calculated = fp.Calculated - 50
	}
	fp.ApplySimpleFudges()
	return fp
}

// oneromancentury - `VIII bc?` -->   -750
func oneromancentury(fp FingerPrint) FingerPrint {
	cleaned := stripforromans0.ReplaceAllString(fp.OrigDateString, "")
	cleaned = strings.TrimSpace(cleaned)
	d, e := romannumerals[cleaned]
	if e != true {
		fmt.Printf("oneromancentury parser failed '%s'\n", fp.OrigDateString)
		fp.ParserFailed = true
		return fp
	}
	fp.Calculated = d * 100
	fp.ApplyBCE()
	if fp.Calculated < 0 {
		// 5th bce is not -500 but -450...
		fp.Calculated = fp.Calculated + 50
	} else {
		// 5th ce is not 500 but 450...
		fp.Calculated = fp.Calculated - 50
	}
	fp.ApplySimpleFudges()
	return fp
}

// slashdated - `c.63／2-51／0 bc` --> -57
func slashdated(fp FingerPrint) FingerPrint {
	// note that there is an infinite loop possibility: pickandrunparser() is how you got here
	// if that "／" does not disappear, you could return

	// `(\d+)／\d(\D)`
	cleaned := hasslashdate1.ReplaceAllString(fp.OrigDateString, "$1$2")
	// fmt.Println("slashdated", cleaned)

	newfp := TakeFingerprint(cleaned)
	newfp = pickandrunparser(newfp)
	newfp.OrigDateString = fp.OrigDateString
	return newfp
}

// eitherordate - `618 or 633 ac` --> 618
func eitherordate(fp FingerPrint) FingerPrint {
	// note that there is an infinite loop possibility: pickandrunparser() is how you got here

	// `(.*)( or \d.*\s)
	cleaned := hasor.ReplaceAllString(fp.OrigDateString, "$1 ")
	newfp := TakeFingerprint(cleaned)
	newfp = pickandrunparser(newfp)
	newfp.OrigDateString = fp.OrigDateString
	return newfp
}

func bracketdate(fp FingerPrint) FingerPrint {
	// note that there is an infinite loop possibility: pickandrunparser() is how you got here
	cleaned := hasbracket.ReplaceAllString(fp.OrigDateString, "")
	// note that there is an infinite loop possibility: pickandrunparser() is how you got here
	newfp := TakeFingerprint(cleaned)
	newfp = pickandrunparser(newfp)
	newfp.OrigDateString = fp.OrigDateString
	return newfp
}
