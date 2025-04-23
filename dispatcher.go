package main

import "fmt"

func pickandrunparser(fp FingerPrint) FingerPrint {
	// order of the tests matters
	if fp.HasTwoArabic && fp.HasOneSpan && (fp.Has2dArabic || fp.Has3dArabic || fp.Has4Arabic) {
		return twoarabicsimple(fp)
	}
	if fp.HasOneArabic && (fp.Has2dArabic || fp.Has3dArabic || fp.Has4Arabic) {
		return onearabicsimple(fp)
	}
	if (fp.HasTH || fp.HasGerman) && fp.HasTwoArabic && (fp.Has1dArabic || fp.Has2dArabic) {
		return twoarabiccenturies(fp)
	}
	if (fp.HasTH || fp.HasGerman) && fp.HasOneArabic && (fp.Has1dArabic || fp.Has2dArabic) {
		return onearabiccentury(fp)
	}
	if fp.HasTwoRoman && fp.HasOneSpan {
		return twoaromancenturies(fp)
	}
	if fp.Has1dArabic && !fp.HasOneSpan {
		// more dangerous, but the way to do 'c 5p
		return onearabiccentury(fp)
	}
	if fp.HasOneRoman {
		return oneromancentury(fp)
	}

	// now we are in the zone where recursive calls might be made; look out for infinite loops
	if fp.HasSlashDate {
		return slashdated(fp)
	}
	if fp.HasOR {
		return eitherordate(fp)
	}
	if fp.HasBracket {
		return bracketdate(fp)
	}

	// desperate people should try a lookup: we might have "byzantinisch", vel sim

	fp.LacksParser = true
	fmt.Printf("no parser for '%s'\n", fp.OrigDateString)
	return fp
}
