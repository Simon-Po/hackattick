package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"hackattic/Touchtone-dialing/decoder"
)

type Problem struct {
	WavURL string `json:"wav_url"`
}

type Answer struct {
	Sequence string `json:"sequence"`
}

func fetchProblem(accessToken string) (*Problem, error) {
	url := fmt.Sprintf("https://hackattic.com/challenges/touch_tone_dialing/problem?access_token=%s", accessToken)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var p Problem
	if err := json.Unmarshal(body, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: %s", resp.Status)
	}
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

func sendAnswer(accessToken, sequence string) error {
	url := fmt.Sprintf("https://hackattic.com/challenges/touch_tone_dialing/solve?access_token=%s&playground=1", accessToken)
	ans := Answer{Sequence: sequence}
	data, err := json.Marshal(ans)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("Response:", string(body))
	return nil
}

func main() {
	accessToken := "be8466c39c975877"
	prob, err := fetchProblem(accessToken)
	if err != nil {
		log.Fatal("Could not fetch problem:", err)
	}

	wavFile := "sample.wav"
	if err = downloadFile(wavFile, prob.WavURL); err != nil {
		log.Fatal("Could not download WAV:", err)
	}
	b, err := os.ReadFile("sample.wav")
	if err != nil {
		log.Fatal("could not read file:", err)
	}

	d := decoder.New(b)
	_ = d.Decode()
	var sampleRate int32
	var bitsPerSample int16
	var numBytes int

	// Decoding the header and getting out the info that i Actually need
	if true {
		fmt.Printf("RIFF:                   %s\n", string(d.Get(4)))
		fmt.Printf("File Size:              %d\n", d.GetInt32Le())
		fmt.Printf("Form Type:              %s\n", string(d.Get(4)))
		fmt.Printf("Format Chunk ID:        %s\n", string(d.Get(4)))
		fmt.Printf("Format Chunk Size:      %d\n", d.GetInt32Le())
		fmt.Printf("Audio Format:           %d\n", d.GetInt16Le())
		fmt.Printf("Number of Channels:     %d\n", d.GetInt16Le())
		sampleRate = d.GetInt32Le()
		fmt.Printf("Sample Rate:            %d\n", sampleRate)
		fmt.Printf("Byte Rate:              %d\n", d.GetInt32Le())
		fmt.Printf("Block Align:            %d\n", d.GetInt16Le())
		bitsPerSample = d.GetInt16Le()
		fmt.Printf("Bits Per Sample:        %d\n", bitsPerSample)
		fmt.Printf("Data Chunk ID:          %s\n", string(d.Get(4)))
		numBytes = int(d.GetInt32Le())
		fmt.Printf("Data Chunk Size (bytes):%d\n", numBytes)

	} else {
		_ = d.Get(4)
		_ = d.GetInt32Le()
		_ = d.Get(4)
		_ = d.Get(4)
		_ = d.GetInt32Le()
		_ = d.GetInt16Le()
		_ = d.GetInt16Le()
		sampleRate = d.GetInt32Le()
		_ = d.GetInt32Le()
		_ = d.GetInt16Le()
		bitsPerSample = d.GetInt16Le()
		_ = d.Get(4) // “data”
		numBytes = int(d.GetInt32Le())
	}

	raw := d.Get(numBytes)

	samplesCount := numBytes / (int(bitsPerSample) / 8)
	samples := make([]int16, samplesCount)
	for i := 0; i < samplesCount; i++ {
		offset := i * (int(bitsPerSample) / 8)
		samples[i] = int16(binary.LittleEndian.Uint16(raw[offset : offset+2]))
	}

	// simplest windowing logic
	windowSize := int(float32(sampleRate) * 0.03) // 100 ms windows
	threshold := 0.064
	ratio := 4.0

	targets := []float64{697, 770, 852, 941, 1209, 1336, 1477, 1633}
	var digits []string

	for i := 0; i+windowSize < len(samples); i += windowSize {
		block := samples[i : i+windowSize]

		floats := make([]float64, len(block))
		for j, v := range block {
			floats[j] = float64(v) / 32768.0
		}

		mags := decoder.GoertzelMagnitudeScaled(floats, int(sampleRate), targets)
		digit := decoder.DecideDigit(targets, mags, threshold, ratio)
		if digit != "" {
			digits = append(digits, digit)
			i += windowSize
		}
	}

	result := decoder.CollapseRepeats(digits)
	if err := sendAnswer(accessToken, result); err != nil {
		log.Fatal("Could not send answer:", err)
	}
	fmt.Printf("Detected digits: %s\n", result)
}
