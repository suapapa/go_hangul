package cp949

import (
	"golang.org/x/text/encoding"
)

// CP949Encoding provides the Encoding interface for CP949 encoding.
type CP949Encoding struct{}

// NewCP949Encoding creates a new CP949Encoding.
func NewCP949Encoding() *CP949Encoding {
	return &CP949Encoding{}
}

// NewDecoder returns a Decoder for CP949 encoding.
func (e *CP949Encoding) NewDecoder() *encoding.Decoder {
	return &encoding.Decoder{Transformer: new(cp949Decoder)}
}

// NewEncoder returns an Encoder for CP949 encoding.
func (e *CP949Encoding) NewEncoder() *encoding.Encoder {
	return &encoding.Encoder{Transformer: new(cp949Encoder)}
}

type cp949Decoder struct{}

// Transform converts data from CP949 to UTF-8.
func (d *cp949Decoder) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	// Use the From function to convert CP949 to UTF-8.
	converted, err := From(src)
	if err != nil {
		return 0, 0, err
	}

	// Copy the converted data to dst.
	copy(dst, converted)
	return len(converted), len(src), nil
}

type cp949Encoder struct{}

// Transform converts data from UTF-8 to CP949.
func (e *cp949Encoder) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	// Use the To function to convert UTF-8 to CP949.
	converted, err := To(src)
	if err != nil {
		return 0, 0, err
	}

	// Copy the converted data to dst.
	copy(dst, converted)
	return len(converted), len(src), nil
}

// Reset resets the state of the Transformer.
func (d *cp949Decoder) Reset() {}
func (e *cp949Encoder) Reset() {}
