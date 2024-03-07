package xortransport

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
			want:    "0000000bc3a6caa6c9e99ef183ef8b",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputBytes := []byte(tt.input)
			got := EncryptBytes(inputBytes)
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
			input:   "0000000bc3a6caa6c9e99ef183ef8b",
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
			got := DecryptBytes(inputBytes)
			gotBytesAsString := string(got)
			if !reflect.DeepEqual(gotBytesAsString, tt.want) {
				t.Errorf("DecryptBytes() = '%s', want '%v'", gotBytesAsString, tt.want)
			}
		})
	}
}
