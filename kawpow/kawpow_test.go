// ethash: C/C++ implementation of Ethash, the Ethereum Proof of Work algorithm.
// Copyright 2018-2019 Pawel Bylica.
// Licensed under the Apache License, Version 2.0.

package kawpow

import (
	"bytes"
	"testing"

	"github.com/sencha-dev/powkit/internal/common/testutil"
)

func TestComputeRavencoin(t *testing.T) {
	tests := []struct {
		height uint64
		nonce  uint64
		hash   []byte
		mix    []byte
		digest []byte
	}{
		{
			height: 1880000,
			nonce:  0x25ca7a0109cf8f2d,
			hash:   testutil.MustDecodeHex("439fe77436016853df3ec5ca24d654da32845f389334cadad356a42ef62e19cd"),
			mix:    testutil.MustDecodeHex("7ff370c848f6f553fa7ca8d68c3515dca5833d330b048dd8a27ccd4b137dc1d7"),
			digest: testutil.MustDecodeHex("00000000000021c9af1f188fcf1c913b2583133b0fe58bb6ff0532dad895fdfc"),
		},
		{
			height: 1888000,
			nonce:  0xdc30900000a4c493,
			hash:   testutil.MustDecodeHex("911a676cf0e5077a24e4917483bcca4bbd461a679b1a780b9d15c8b6bf5bc1d7"),
			mix:    testutil.MustDecodeHex("28605c4c11c72f1d3af0baef8cca10ccf570aa0c89c93596b2bfd485e30bd9f7"),
			digest: testutil.MustDecodeHex("0000000000000d6ea9c4b4b759a157bc2ef74bd9a7ed04ca6ac9611e218ccddc"),
		},
		{
			height: 1888509,
			nonce:  0xf09b0e1342275f3f,
			hash:   testutil.MustDecodeHex("14f2c18d74d48abe637437458c10ff5283a9a5197e8b5e740a161f4411b97a43"),
			mix:    testutil.MustDecodeHex("3dedcc6fb28c6bf3f5d203f29188bef3ff86688be34c93f28bd227eced9226e4"),
			digest: testutil.MustDecodeHex("0000000000005e6585e5e6ab7e4d75a98810204341def05823ad3a5ca1fa0d83"),
		},
	}

	client := NewRavencoin()
	for i, tt := range tests {
		mix, digest := client.Compute(tt.height, tt.nonce, tt.hash)

		if bytes.Compare(mix, tt.mix) != 0 {
			t.Errorf("failed on %d: mix mismatch: have %x, want %x", i, mix, tt.mix)
		}

		if bytes.Compare(digest, tt.digest) != 0 {
			t.Errorf("failed on %d: digest mismatch: have %x, want %x", i, digest, tt.digest)
		}
	}
}