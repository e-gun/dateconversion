package main

import "fmt"

func pickbasicparser(fp FingerPrint) FingerPrint {
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
	if fp.HasOneRoman {
		return oneromancentury(fp)
	}

	// desperate people should try a lookup: we might have "byzantinisch", vel sim

	fp.LacksParser = true
	fmt.Printf("no parser for '%s'\n", fp.OrigDateString)
	return fp
}
