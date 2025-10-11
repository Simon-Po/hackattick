package decoder

import (
	"bytes"
	"encoding/binary"
)

// Error is the domain specific error type
type Error string

func (e Error) Error() string {
	return "Error: " + string(e)
}

const (
	// BasicDecodingError Describes an unknown error that happened but could not be determined why.
	BasicDecodingError Error = "Could not decode"
)

type Decoder struct {
	content []byte
}

func New(content []byte) Decoder {
	return Decoder{
		content: content,
	}
}

// Get is the generic get function it slurps the requested amount of bytes and returns a []byte
func (d *Decoder) Get(count int) []byte {
	c := d.content[0:count]
	d.content = d.content[count:]
	return c
}

// GetInt32Le slurps 4 bytes from the Decoders contents and returns an int32
func (d *Decoder) GetInt32Le() int32 {
	var i int32
	binary.Read(bytes.NewReader(d.Get(4)), binary.LittleEndian, &i)
	return i
}

func (d *Decoder) GetInt16Le() int16 {
	var i int16
	binary.Read(bytes.NewReader(d.Get(2)), binary.LittleEndian, &i)
	return i
}

func (d *Decoder) Skip(count int) {
	d.content = d.content[count:]
}

func (d *Decoder) Decode() error {
	return nil
}

func (d Decoder) String() string {
	return "This is a decoder"
}
