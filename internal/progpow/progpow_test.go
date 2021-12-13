// ethash: C/C++ implementation of Ethash, the Ethereum Proof of Work algorithm.
// Copyright 2018-2019 Pawel Bylica.
// Licensed under the Apache License, Version 2.0.

package progpow

import (
	"testing"
)

func TestInitMix(t *testing.T) {
	seed := uint64(0xEE304846DDD0A47B)

	lanesExpected := map[int][32]uint32{
		0: [32]uint32{
			0x10C02F0D, 0x99891C9E, 0xC59649A0, 0x43F0394D,
			0x24D2BAE4, 0xC4E89D4C, 0x398AD25C, 0xF5C0E467,
			0x7A3302D6, 0xE6245C6C, 0x760726D3, 0x1F322EE7,
			0x85405811, 0xC2F1E765, 0xA0EB7045, 0xDA39E821,
			0x79FC6A48, 0x089E401F, 0x8488779F, 0xD79E414F,
			0x041A826B, 0x313C0D79, 0x10125A3C, 0x3F4BDFAC,
			0xA7352F36, 0x7E70CB54, 0x3B0BB37D, 0x74A3E24A,
			0xCC37236A, 0xA442B311, 0x955AB27A, 0x6D175B7E,
		},
		13: [32]uint32{
			0x4E46D05D, 0x2E77E734, 0x2C479399, 0x70712177,
			0xA75D7FF5, 0xBEF18D17, 0x8D42252E, 0x35B4FA0E,
			0x462C850A, 0x2DD2B5D5, 0x5F32B5EC, 0xED5D9EED,
			0xF9E2685E, 0x1F29DC8E, 0xA78F098B, 0x86A8687B,
			0xEA7A10E7, 0xBE732B9D, 0x4EEBCB60, 0x94DD7D97,
			0x39A425E9, 0xC0E782BF, 0xBA7B870F, 0x4823FF60,
			0xF97A5A1C, 0xB00BCAF4, 0x02D0F8C4, 0x28399214,
			0xB4CCB32D, 0x83A09132, 0x27EA8279, 0x3837DDA3,
		},
	}

	mix := initMix(seed)

	for lane := range mix {
		if expected, ok := lanesExpected[lane]; ok {
			actual := mix[lane]

			for reg := range expected {
				if actual[reg] != expected[reg] {
					t.Errorf("failed on lane %d, reg %d: have %d, want %d", lane, reg, actual[reg], expected[reg])
					continue
				}
			}
		}
	}
}
