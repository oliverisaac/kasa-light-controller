package plaintransport

type Encrypter struct{}

func (e Encrypter) EncryptBytes(b []byte) ([]byte, error) {
	return EncryptBytes(b)
}

type Decrypter struct{}

func (e Decrypter) EncryptBytes(b []byte) ([]byte, error) {
	return DecryptBytes(b)
}

func EncryptBytes(b []byte) ([]byte, error) {
	output := make([]byte, len(b))
	for i, b := range b {
		output[i] = b
	}
	return output, nil
}

func DecryptBytes(b []byte) ([]byte, error) {
	return EncryptBytes(b)
}
