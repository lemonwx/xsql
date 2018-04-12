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
