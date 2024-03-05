package plaintransport

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEncryptBytes(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "hello world example",
			input:   "hello world",
			want:    "68656c6c6f20776f726c64",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputBytes := []byte(tt.input)
			got, err := EncryptBytes(inputBytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotBytesAsString := fmt.Sprintf("%x", got)
			if !reflect.DeepEqual(gotBytesAsString, tt.want) {
				t.Errorf("EncryptBytes() = '%s', want '%v'", gotBytesAsString, tt.want)
			}
		})
	}
}

func TestDecryptBytes(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "hello world example",
			input:   "68656c6c6f20776f726c64",
			want:    "hello world",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputBytes, err := hex.DecodeString(tt.input)
			if err != nil {
				panic(err)
			}
			logrus.Errorf("Bytes: % x", inputBytes)
			got, err := DecryptBytes(inputBytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotBytesAsString := string(got)
			if !reflect.DeepEqual(gotBytesAsString, tt.want) {
				t.Errorf("DecryptBytes() = '%s', want '%v'", gotBytesAsString, tt.want)
			}
		})
	}
}
