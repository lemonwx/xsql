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