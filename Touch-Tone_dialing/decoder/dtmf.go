package decoder

import (
	"strings"
)

var (
	dtmfLow  = []float64{697, 770, 852, 941}
	dtmfHigh = []float64{1209, 1336, 1477, 1633}
)

var dtmfMap = map[[2]int]string{
	{0, 0}: "1", {0, 1}: "2", {0, 2}: "3", {0, 3}: "A",
	{1, 0}: "4", {1, 1}: "5", {1, 2}: "6", {1, 3}: "B",
	{2, 0}: "7", {2, 1}: "8", {2, 2}: "9", {2, 3}: "C",
	{3, 0}: "*", {3, 1}: "0", {3, 2}: "#", {3, 3}: "D",
}

func DecideDigit(targets []float64, mags []float64, threshold float64, ratio float64) string {
	if len(targets) != len(mags) {
		return ""
	}
	bestLowIdx := -1
	bestLowVal := 0.0
	secondLow := 0.0

	bestHighIdx := -1
	bestHighVal := 0.0
	secondHigh := 0.0

	cutoffLow := len(dtmfLow)
	for i, v := range mags {
		if i < cutoffLow {
			if v > bestLowVal {
				secondLow = bestLowVal
				bestLowVal = v
				bestLowIdx = i
			} else if v > secondLow {
				secondLow = v
			}
		} else {
			if v > bestHighVal {
				secondHigh = bestHighVal
				bestHighVal = v
				bestHighIdx = i
			} else if v > secondHigh {
				secondHigh = v
			}
		}
	}

	if bestLowIdx < 0 || bestHighIdx < 0 {
		return ""
	}
	if bestLowVal < threshold || bestHighVal < threshold {
		return ""
	}
	if secondLow > 0 && bestLowVal/secondLow < ratio {
		return ""
	}
	if secondHigh > 0 && bestHighVal/secondHigh < ratio {
		return ""
	}

	digit, ok := dtmfMap[[2]int{bestLowIdx, bestHighIdx - cutoffLow}]
	if !ok {
		return ""
	}
	return digit
}

func CollapseRepeats(digits []string) string {
	// I though this makes sense but fucked me up pretty bad
	// var out []string
	// for i, d := range digits {
	// 	if d == "" {
	// 		continue
	// 	}
	// 	if i == 0 || d != digits[i-1] {
	// 		out = append(out, d)
	// 	}
	// }
	return strings.Join(digits, "")
}
