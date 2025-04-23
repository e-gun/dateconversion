package main

import (
	"regexp"
)

var (
	hasonearabic   = regexp.MustCompile(`\d+`)
	hastwoarabic   = regexp.MustCompile(`\d+\D+\d`)
	hasthreearabic = regexp.MustCompile(`\d+\D+\d+\D+\d`)
	has4darabic    = regexp.MustCompile(`\d\d\d\d`)
	has3darabic    = regexp.MustCompile(`\d\d\d`)
	has2darabic    = regexp.MustCompile(`\d\d`)
	has1darabic    = regexp.MustCompile(`\d`)
	hasoneroman    = regexp.MustCompile(`[IXV]+`)
	hastworoman    = regexp.MustCompile(`[IXV]+[^IXV]+[IXV]+`)
	germandate     = regexp.MustCompile(`(nach|vor|Jhdt|Jh)`)
	hasbce         = regexp.MustCompile(`(vor|BC|B\.C\.|bc|v.Chr|a$)`)
	hasce          = regexp.MustCompile(`(nach|AD|A\.D\.|ad|n.Chr|ac$|p$)`)
	has1sthalf     = regexp.MustCompile(`(1st half|1. H.|1. Hälfte)`)
	has2ndhalf     = regexp.MustCompile(`(2nd half|2. H.|2. Hälfte)`)
	hasbeg         = regexp.MustCompile(`(^in |init |init s |early|beg|Anf\.)`)
	hasmid         = regexp.MustCompile(`mid`)
	hasend         = regexp.MustCompile(`(late|end)`)
	hasthird       = regexp.MustCompile(`Drittel`)
	hasquarter     = regexp.MustCompile(`Viertel`)
	hasth          = regexp.MustCompile(`(\dth|2nd)`)
	hasonespan     = regexp.MustCompile(`[\-／]`)
	hastwospans    = regexp.MustCompile(`[\-／][^\-／][\-／]`)
	hasbefore      = regexp.MustCompile(`(before|bef\.)`)
	hasafter       = regexp.MustCompile(`(^p |after|aft\.)`)
)

func TakeFingerprint(text string) FingerPrint {
	fp := FingerPrint{
		OrigDateString: text,
		Calculated:     9999,
	}
	if hasonearabic.MatchString(text) {
		fp.HasOneArabic = true
	}
	if hastwoarabic.MatchString(text) {
		fp.HasTwoArabic = true
	}
	if hasthreearabic.MatchString(text) {
		fp.HasThreeArabic = true
	}
	if has4darabic.MatchString(text) {
		fp.Has4Arabic = true
	}
	if has3darabic.MatchString(text) {
		fp.Has3dArabic = true
	}
	if has2darabic.MatchString(text) {
		fp.Has2dArabic = true
	}
	if has1darabic.MatchString(text) {
		fp.Has1dArabic = true
	}
	if hasoneroman.MatchString(text) {
		fp.HasOneRoman = true
	}
	if hastworoman.MatchString(text) {
		fp.HasTwoRoman = true
	}
	if germandate.MatchString(text) {
		fp.HasGerman = true
	}
	if hasbce.MatchString(text) {
		fp.HasBCE = true
	}
	if hasce.MatchString(text) {
		fp.HasCE = true
	}
	if hasonespan.MatchString(text) {
		fp.HasOneSpan = true
	}
	if hastwospans.MatchString(text) {
		fp.HasTwoSpans = true
	}
	if hasth.MatchString(text) {
		fp.HasTH = true
	}
	if hasend.MatchString(text) || has2ndhalf.MatchString(text) {
		fp.HasEnd = true
	}
	if hasbeg.MatchString(text) || has1sthalf.MatchString(text) {
		fp.HasBeginning = true
	}
	if hasbefore.MatchString(text) {
		fp.HasAnte = true
	}
	if hasafter.MatchString(text) {
		fp.HasPost = true
	}

	return fp
}
