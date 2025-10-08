package hmu

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrUnmarshal Error = "Could not unmarshal the content"
)

type reqbody struct {
	Bytes string `json:"bytes"`
}

func FetchAndParse(p string) []byte {
	url := p

	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	b, err := unmarshal(body)
	if err != nil {
		panic(err)
	}

	fmt.Println("encoded: ", b.Bytes)
	decoded := decode(b)

	fmt.Println("decoded: ", string(decoded))
	return decoded
}

func decode(b reqbody) []byte {
	decoded, err := base64.StdEncoding.DecodeString(b.Bytes)
	if err != nil {
		panic(err)
	}
	return decoded
}

func unmarshal(body []byte) (reqbody, error) {
	b := reqbody{}
	err := json.Unmarshal(body, &b)
	if err != nil {
		return reqbody{}, ErrUnmarshal
	}

	return b, nil
}

func GetInt(b []byte) (int32, []byte) {
	i := b[:4]
	rest := b[4:]
	return int32(binary.LittleEndian.Uint32(i)), rest
}

func GetUInt(b []byte) (uint32, []byte) {
	i := b[:4]
	rest := b[4:]
	return binary.LittleEndian.Uint32(i), rest
}

func GetShort(b []byte) (int16, []byte) {
	i := b[:2]
	rest := b[2:]
	return int16(binary.LittleEndian.Uint16(i)), rest
}

func Skip(b []byte) []byte {
	return b[2:]
}

func GetFloat(b []byte) (float32, []byte) {
	i := b[:4]
	rest := b[4:]
	return math.Float32frombits(binary.LittleEndian.Uint32(i)), rest
}

func GetDouble(b []byte) (float64, []byte) {
	i := b[:8]
	rest := b[8:]
	return math.Float64frombits(binary.LittleEndian.Uint64(i)), rest
}

func GetDoubleBE(b []byte) (float64, []byte) {
	i := b[:8]
	rest := b[8:]
	return math.Float64frombits(binary.BigEndian.Uint64(i)), rest
}
