package main

import (
	"encoding/binary"
	"errors"
)

var (
	ErrBadIndex = errors.New("Broken compression index.")
)

const (
	MaxBlockSize   = 64 * 1024
	IndexBlockSize = 16
	TailSize       = 16
)

const (
	AlgoNone = iota
	AlgoSnappy
	// TODO: AlgoLZ4?
)

type Algorithm byte

// Ist Block ein exportierter Typ?
type Block struct {
	rawOff int64
	zipOff int64
}

func (bl *Block) marshal(buf []byte) {
	binary.LittleEndian.PutUint64(buf[0:8], uint64(bl.rawOff))
	binary.LittleEndian.PutUint64(buf[8:16], uint64(bl.zipOff))
}

func (bl *Block) unmarshal(buf []byte) {
	bl.rawOff = int64(binary.LittleEndian.Uint64(buf[0:8]))
	bl.zipOff = int64(binary.LittleEndian.Uint64(buf[8:16]))
}
