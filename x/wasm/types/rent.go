package types

import (
	"encoding/binary"
	"errors"
)

// custom struct instead of protobuf to minimize state needed
type RentInfo struct {
	Balance          int64
	LastChargedBlock uint64
}

func (r *RentInfo) Marshal() []byte {
	b, l := make([]byte, 8), make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(r.Balance))
	binary.BigEndian.PutUint64(l, r.LastChargedBlock)
	return append(b, l...)
}

func (r *RentInfo) Unmarshal(bz []byte) error {
	if len(bz) != 16 {
		return errors.New("RentInfo must have exactly 16 bytes")
	}
	r.Balance = int64(binary.BigEndian.Uint64(bz[:8]))
	r.LastChargedBlock = binary.BigEndian.Uint64(bz[8:])
	return nil
}
