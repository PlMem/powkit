package cuckoo

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"

	"github.com/sencha-dev/powkit/internal/crypto"
)

type Client struct {
	proofSize int
	edgeBits  int
	edgeMask  uint64
	nodeBits  int
	nodeMask  uint64
}

func New(edgeBits, nodeBits, proofSize int) *Client {
	c := &Client{
		proofSize: proofSize,
		edgeBits:  edgeBits,
		edgeMask:  (uint64(1) << edgeBits) - 1,
		nodeBits:  nodeBits,
		nodeMask:  (uint64(1) << nodeBits) - 1,
	}

	return c
}

func NewAeternity() *Client {
	return New(29, 29, 42)
}

func (c *Client) Verify(hash []byte, nonce uint64, sols []uint64) (bool, error) {
	if len(hash) != 32 {
		return false, fmt.Errorf("hash must be 32 bytes")
	} else if len(sols) != 42 {
		return false, fmt.Errorf("sols must be 42 uint64s")
	}

	// encode header
	nonceBytes := make([]uint8, 8)
	binary.LittleEndian.PutUint64(nonceBytes, nonce)
	hashEncoded := []byte(base64.StdEncoding.EncodeToString(hash))
	nonceEncoded := []byte(base64.StdEncoding.EncodeToString(nonceBytes))
	header := append(hashEncoded, append(nonceEncoded, make([]byte, 24)...)...)

	// create siphash keys
	h := crypto.Blake2b256(header)
	keys := [4]uint64{
		binary.LittleEndian.Uint64(h[0:8]),
		binary.LittleEndian.Uint64(h[8:16]),
		binary.LittleEndian.Uint64(h[16:24]),
		binary.LittleEndian.Uint64(h[24:32]),
	}

	return verify(c.proofSize, c.edgeMask, keys, sols)
}
