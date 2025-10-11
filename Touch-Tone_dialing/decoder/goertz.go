package decoder

import (
	"math"
)

// GoertzelMagnitudeScaled computes "power per sample" magnitudes for each target.
func GoertzelMagnitudeScaled(samples []float64, sampleRate int, targets []float64) []float64 {
	N := float64(len(samples))
	mags := make([]float64, len(targets))

	for i, freq := range targets {
		k := int(math.Round(freq * float64(len(samples)) / float64(sampleRate)))
		omega := 2.0 * math.Pi * float64(k) / float64(len(samples))
		coeff := 2.0 * math.Cos(omega)

		sPrev := 0.0
		sPrev2 := 0.0
		for _, x := range samples {
			s := x + coeff*sPrev - sPrev2
			sPrev2 = sPrev
			sPrev = s
		}
		realPart := sPrev - sPrev2*math.Cos(omega)
		imagPart := sPrev2 * math.Sin(omega)
		magSq := realPart*realPart + imagPart*imagPart

		powerPerSample := magSq / N
		mags[i] = powerPerSample
	}

	return mags
}
