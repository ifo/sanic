package sanic

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"strings"
)

func IntToBytes(i int64) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, i)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func IntToString(i int64) (string, error) {
	bts, err := IntToBytes(i)
	if err != nil {
		return "", err
	}
	bts = removeTrailingZeroBytes(bts)
	return MakeURLFriendly(base64.RawStdEncoding.EncodeToString(bts)), nil
}

func RawIntToString(i int64) (string, error) {
	bts, err := IntToBytes(i)
	return base64.RawStdEncoding.EncodeToString(bts), err
}

func removeTrailingZeroBytes(bts []byte) []byte {
	for i := len(bts) - 1; i > 0; i-- {
		if bts[i] != 0 {
			return bts[:i+1]
		}
	}
	return bts
}

// MakeURLFriendly replaces the base64 encoding's "+" and "/" characters
// with "-" and "_" characters
//
// + -> -, / -> _
func MakeURLFriendly(str string) string {
	return strings.Replace(strings.Replace(str, "+", "-", -1), "/", "_", -1)
}
