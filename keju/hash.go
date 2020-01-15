package keju

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"fmt"
)

func init() {
	gob.Register(MD5Hash{})
}

type Hash interface {
	Equal(hash Hash) bool
	toString() string
}

type MD5Hash [16]byte

func NewMD5Hash(byteData []byte) MD5Hash {
	return md5.Sum(byteData)
}

func NewMD5HashFromString(str string) MD5Hash {
	return NewMD5Hash([]byte(str))
}

func (h MD5Hash) Equal(another Hash) bool {
	anotherMD5, yes := another.(MD5Hash)
	if !yes {
		return false
	}
	return bytes.Equal(h[:], anotherMD5[:])
}

func (h MD5Hash) toString() string {
	return fmt.Sprintf("%x", h)
}
