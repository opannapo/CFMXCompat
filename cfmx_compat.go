package gocfmx

import (
	"encoding/hex"
	"strings"
)

type CFmxCompat struct {
	m_Key    string
	m_LFSR_A int
	m_LFSR_B int
	m_LFSR_C int
	m_Mask_A int
	m_Mask_B int
	m_Mask_C int
	m_Rot0_A int
	m_Rot0_B int
	m_Rot0_C int
	m_Rot1_A int
	m_Rot1_B int
	m_Rot1_C int
}

func NewCFmxCompat(key string) *CFmxCompat {
	//set default value, attr pindah ke struct
	instance := &CFmxCompat{
		m_Key:    key,
		m_LFSR_A: 0x13579bdf,
		m_LFSR_B: 0x2468ace0,
		m_LFSR_C: 0xfdb97531,
		m_Mask_A: 0x80000062,
		m_Mask_B: 0x40000020,
		m_Mask_C: 0x10000002,
		m_Rot0_A: 0x7fffffff,
		m_Rot0_B: 0x3fffffff,
		m_Rot0_C: 0xfffffff,
		m_Rot1_A: 0x80000000,
		m_Rot1_B: 0xc0000000,
		m_Rot1_C: 0xf0000000,
	}

	instance.setKey(key)

	return instance
}

func (cc *CFmxCompat) Encrypt(input string) (output string) {
	out := cc.transformString([]byte(input))
	output = hex.EncodeToString(out[:])
	return strings.ToUpper(output)
}

func (cc *CFmxCompat) transformString(inBytes []byte) []byte {
	length := len(inBytes)
	outBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		outBytes[i] = cc.transformByte(inBytes[i])
	}
	return outBytes
}

func (cc *CFmxCompat) transformByte(target byte) byte {
	var crypto byte = 0
	b := cc.m_LFSR_B & 1
	c := cc.m_LFSR_C & 1
	for i := 0; i < 8; i++ {
		if 0 != (cc.m_LFSR_A & 1) {
			cc.m_LFSR_A = cc.m_LFSR_A ^ cc.m_Mask_A>>1 | cc.m_Rot1_A
			if 0 != (cc.m_LFSR_B & 1) {
				cc.m_LFSR_B = cc.m_LFSR_B ^ cc.m_Mask_B>>1 | cc.m_Rot1_B
				b = 1
			} else {
				cc.m_LFSR_B = cc.m_LFSR_B >> 1 & cc.m_Rot0_B
				b = 0
			}
		} else {
			cc.m_LFSR_A = cc.m_LFSR_A >> 1 & cc.m_Rot0_A
			if 0 != (cc.m_LFSR_C & 1) {
				cc.m_LFSR_C = cc.m_LFSR_C ^ cc.m_Mask_C>>1 | cc.m_Rot1_C
				c = 1
			} else {
				cc.m_LFSR_C = cc.m_LFSR_C >> 1 & cc.m_Rot0_C
				c = 0
			}
		}

		c1 := (crypto << 1)
		c2 := (b ^ c)
		crypto = c1 | byte(c2)
	}
	target ^= crypto
	return target
}

func (cc *CFmxCompat) setKey(key string) {
	m_Key := key
	if len(key) == 0 {
		key = "Default Seed"
	}

	seedLen := len(key)
	if seedLen < 12 {
		seedLen = 12
	}
	Seed := make([]int, seedLen)
	for i, char := range key {
		Seed[i] = int(char)
	}
	originalLength := len(m_Key)
	for i := 0; originalLength+i < 12; i++ {
		Seed[originalLength+i] = Seed[i]
	}
	for i := 0; i < 4; i++ {
		cc.m_LFSR_A = (cc.m_LFSR_A << 8) | Seed[i+4]
		cc.m_LFSR_B = (cc.m_LFSR_B << 8) | Seed[i+4]
		cc.m_LFSR_C = (cc.m_LFSR_C << 8) | Seed[i+4]
	}
	if 0 == cc.m_LFSR_A {
		cc.m_LFSR_A = 0x13579bdf
	}
	if 0 == cc.m_LFSR_B {
		cc.m_LFSR_B = 0x2468ace0
	}
	if 0 == cc.m_LFSR_C {
		cc.m_LFSR_C = 0xfdb97531
	}
}
