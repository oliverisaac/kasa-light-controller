package xortransport

import "fmt"

type Encrypter struct{}

func (e Encrypter) EncryptBytes(b []byte) []byte {
	return EncryptBytes(b)
}

type Decrypter struct{}

func (e Decrypter) EncryptBytes(b []byte) []byte {
	return DecryptBytes(b)
}

func EncryptBytes(b []byte) []byte {
	var xorBytes byte = 0xAB
	output := make([]byte, len(b)+4)
	output[0] = 0x0
	output[1] = 0x0
	output[2] = 0x0
	output[3] = byte(len(b))
	for i, b := range b {
		output[i+4] = xorBytes ^ b
		xorBytes = output[i+4]
	}
	return output
}

func DecryptBytes(b []byte) []byte {
	if len(b) == 0 {
		panic(fmt.Errorf("No bytes to decrypt"))
	}
	var xorBytes byte = 0xAB
	output := make([]byte, len(b)-4)
	for i, b := range b[4:] {
		output[i] = xorBytes ^ b
		xorBytes = b
	}
	return output
}
