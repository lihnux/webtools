package utils

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"time"
)

type UniqueID [12]byte

var NilUniqueID UniqueID

var ErrInvalidHex = errors.New("Hex is invalid")

func random8Bytes() []byte {
	var b [8]byte
	if _, err := io.ReadFull(rand.Reader, b[:]); err != nil {
		slog.With("reason", err.Error()).Error("Cannot generate 8 bytes random array")
	}
	return b[:]
}

func NewUniqueID() UniqueID {
	return NewUniqueIDFromTimestamp(time.Now())
}

func NewUniqueIDFromTimestamp(timestamp time.Time) UniqueID {
	var b [12]byte

	binary.BigEndian.PutUint32(b[0:4], uint32(timestamp.Unix()))
	copy(b[4:12], random8Bytes())
	return b
}

func (id UniqueID) Timestamp() time.Time {
	unixSecs := binary.BigEndian.Uint32(id[0:4])
	return time.Unix(int64(unixSecs), 0).UTC()
}

func (id UniqueID) Hex() string {
	var buf [24]byte
	hex.Encode(buf[:], id[:])
	return string(buf[:])
}

func (id UniqueID) String() string {
	return fmt.Sprintf("UniqueID(%q)", id.Hex())
}

func (id UniqueID) IsZero() bool {
	return id == NilUniqueID
}

func (id UniqueID) MarshalText() ([]byte, error) {
	return []byte(id.Hex()), nil
}

func (id *UniqueID) UnmarshalText(b []byte) error {
	oid, err := UniqueIDFromHex(string(b))
	if err != nil {
		return err
	}
	*id = oid
	return nil
}

func (id UniqueID) MarshalBinary() ([]byte, error) {
	return id.MarshalJSON()
}

func (id UniqueID) UnmarshalBinary(data []byte) error {
	return id.UnmarshalJSON(data)
}

func (id UniqueID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.Hex())
}

func (id *UniqueID) UnmarshalJSON(b []byte) error {
	// Ignore "null" to keep parity with the standard library. Decoding a JSON null into a non-pointer ObjectID field
	// will leave the field unchanged. For pointer values, encoding/json will set the pointer to nil and will not
	// enter the UnmarshalJSON hook.
	if string(b) == "null" {
		return nil
	}

	var err error
	switch len(b) {
	case 12:
		copy(id[:], b)
	default:
		// Extended JSON
		var res interface{}
		err := json.Unmarshal(b, &res)
		if err != nil {
			fmt.Println(err)
			return err
		}
		str, ok := res.(string)
		if !ok {
			fmt.Println("not a valid JSON UniqueID")
			return errors.New("not a valid JSON UniqueID")
		}

		if len(str) == 0 {
			copy(id[:], NilUniqueID[:])
			return nil
		}

		if len(str) != 24 {
			return fmt.Errorf("cannot unmarshal into an UniqueID, the length must be 24 but it is %d", len(str))
		}

		_, err = hex.Decode(id[:], []byte(str))
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return err
}

func UniqueIDFromHex(s string) (UniqueID, error) {
	if len(s) != 24 {
		return NilUniqueID, errors.New("Hex string length is invalid")
	}

	var uid [12]byte
	_, err := hex.Decode(uid[:], []byte(s))
	if err != nil {
		return NilUniqueID, ErrInvalidHex
	}

	return uid, nil
}
