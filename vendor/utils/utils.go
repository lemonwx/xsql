/**
 *  author: lim
 *  data  : 18-4-9 下午10:26
 */

package utils

import "bytes"

func BytesContains(b []byte, bs [][]byte) bool {
	for _, tmp := range bs {
		if bytes.Equal(tmp, b) {
			return true
		}
	}
	return false
}

func CompareIntSlice(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for idx := 0; idx < len(s1); idx += 1 {
		if s1[idx] != s2[idx] {
			return false
		}
	}

	return true
}

func ContainsIntSlice(s []int, tgt int) bool {
	for _, item := range s {
		if tgt == item {
			return true
		}
	}

	return false
}

func Uint64ToBytes(n uint64) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
		byte(n >> 32),
		byte(n >> 40),
		byte(n >> 48),
		byte(n >> 56),
	}
}

func Uint64ToString(n uint64) []byte {
	var a [20]byte
	i := 20

	// U+0030 = 0
	// ...
	// U+0039 = 9

	var q uint64
	for n >= 10 {
		i--
		q = n / 10
		a[i] = uint8(n-q*10) + 0x30
		n = q
	}

	i--
	a[i] = uint8(n) + 0x30

	return a[i:]
}


func AppendLengthEncodedInteger(b []byte, n uint64) []byte {
	switch {
	case n <= 250:
		return append(b, byte(n))

	case n <= 0xffff:
		return append(b, 0xfc, byte(n), byte(n>>8))

	case n <= 0xffffff:
		return append(b, 0xfd, byte(n), byte(n>>8), byte(n>>16))
	}
	return append(b, 0xfe, byte(n), byte(n>>8), byte(n>>16), byte(n>>24),
		byte(n>>32), byte(n>>40), byte(n>>48), byte(n>>56))
}
