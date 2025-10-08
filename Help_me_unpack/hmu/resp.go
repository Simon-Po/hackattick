package hmu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// int: the signed integer value
// uint: the unsigned integer value
// short: the decoded short value
// float: surprisingly, the float value
// double: the double value - shockingly

type Response struct {
	I        int32   `json:"int"`
	UI       uint32  `json:"uint"`
	Short    int16   `json:"short"`
	Float    float32 `json:"float"`
	Double   float64 `json:"double"`
	DoubleBe float64 `json:"big_endian_double"`
}

func SendResult(r Response) {
	body, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	fmt.Println("Sending: ", string(body))
	resp, err := http.Post("https://hackattic.com/challenges/help_me_unpack/solve?access_token=be8466c39c975877", "application/json", bytes.NewReader(body))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	rBody, _ := io.ReadAll(resp.Body)

	fmt.Println(string(rBody))
}
