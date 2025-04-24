package main

import (
	"fmt"
	"github.com/e-gun/dateconversion/structs"
)

func pickandrunparser(fp structs.FingerPrint) structs.FingerPrint {
	// order of the tests matters
	if fp.HasTwoArabic && fp.HasOneSpan && (fp.Has2dArabic || fp.Has3dArabic || fp.Has4Arabic) {
		// twoarabicsimple -  `1-50 ac` --> 25
		return twoarabicsimple(fp)
	}
	if fp.HasOneArabic && (fp.Has2dArabic || fp.Has3dArabic || fp.Has4Arabic) {
		// onearabicsimple: `101 bc` --> -101
		return onearabicsimple(fp)
	}
	if (fp.HasTH || fp.HasGerman) && fp.HasTwoArabic && (fp.Has1dArabic || fp.Has2dArabic) {
		// twoarabiccenturies -  `7th-9th ac` --> 750
		return twoarabiccenturies(fp)
	}
	if (fp.HasTH || fp.HasGerman) && fp.HasOneArabic && (fp.Has1dArabic || fp.Has2dArabic) {
		// onearabiccentury : `early 5th bc` --> -475
		return onearabiccentury(fp)
	}
	if fp.HasTwoRoman && fp.HasOneSpan {
		// twoaromancenturies -  `7th-9th ac` --> 800
		return twoaromancenturies(fp)
	}
	if fp.Has1dArabic && !fp.HasOneSpan {
		// more dangerous, but the way to do 'c 5p
		return onearabiccentury(fp)
	}
	if fp.HasOneRoman {
		// oneromancentury - `VIII bc?` -->   -750
		return oneromancentury(fp)
	}
	if fp.HasMultiDigitArabic {
		// multislashdate - 109／108／106／105 BC --> -109
		return multislashdate(fp)
	}
	// now we are in the zone where recursive calls might be made; look out for infinite loops
	if fp.HasSlashDate {
		// slashdated - `c.63／2-51／0 bc` --> -57
		return slashdated(fp)
	}
	if fp.HasOR {
		// eitherordate - `618 or 633 ac` --> 618
		return eitherordate(fp)
	}
	if fp.HasBracket {
		// bracketdate - `med Ia [K.87(14)]` --> 50
		return bracketdate(fp)
	}

	// desperate people should try a lookup: we might have "byzantinisch", vel sim

	fp.LacksParser = true
	fmt.Printf("no parser for '%s'\n", fp.OrigDateString)
	return fp
}
