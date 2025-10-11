import numpy as np
import wave
import struct

DTMF_FREQS = {
    "1": (697.0, 1209.0),
    "2": (697.0, 1336.0),
    "3": (697.0, 1477.0),
    "4": (770.0, 1209.0),
    "5": (770.0, 1336.0),
    "6": (770.0, 1477.0),
    "7": (852.0, 1209.0),
    "8": (852.0, 1336.0),
    "9": (852.0, 1477.0),
    "0": (941.0, 1336.0),
    "*": (941.0, 1209.0),
    "#": (941.0, 1477.0),
}

def generate_dtmf_tone(freq1, freq2, duration_s, sample_rate, amplitude=0.5):
    t = np.linspace(0, duration_s, int(sample_rate * duration_s), endpoint=False)
    s1 = amplitude * np.sin(2 * np.pi * freq1 * t)
    s2 = amplitude * np.sin(2 * np.pi * freq2 * t)
    return (s1 + s2) / 2

def generate_sequence(digits, tone_duration=0.3, pause_duration=0.1,
                      sample_rate=4000, amplitude=0.8):
    parts = []
    for d in digits:
        if d not in DTMF_FREQS:
            raise ValueError(f"Digit '{d}' not supported")
        f1, f2 = DTMF_FREQS[d]
        tone = generate_dtmf_tone(f1, f2, tone_duration, sample_rate, amplitude)
        parts.append(tone)
        # append silence
        parts.append(np.zeros(int(sample_rate * pause_duration)))
    # combine parts
    return np.concatenate(parts)

def write_wave(filename, samples, sample_rate=4000):
    # assume samples are floats in range [-1, +1]
    n = len(samples)
    with wave.open(filename, 'w') as wf:
        wf.setnchannels(1)         # mono
        wf.setsampwidth(2)         # 2 bytes = 16 bits
        wf.setframerate(sample_rate)
        for s in samples:
            # clamp
            if s > 1.0: s = 1.0
            if s < -1.0: s = -1.0
            val = int(s * 32767)
            wf.writeframes(struct.pack('<h', val))

if __name__ == "__main__":
    seq = "*#123456789900*#122234567890"
    sample_rate = 4000
    tone_duration = 0.4
    pause_duration = 0.1
    amplitude = 0.8

    audio = generate_sequence(seq,
                              tone_duration=tone_duration,
                              pause_duration=pause_duration,
                              sample_rate=sample_rate,
                              amplitude=amplitude)
    out_file = "test.wav"
    write_wave(out_file, audio, sample_rate=sample_rate)
    print(f"WAV file generated: {out_file}")
