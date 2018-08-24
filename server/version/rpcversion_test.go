/**
 *  author: lim
 *  data  : 18-8-21 下午10:01
 */

package version

import "testing"

func TestInUseAndNext(t *testing.T) {

	NewRpcPool(10, 100, "192.168.1.2:1235")

	ret, err := InUseAndNext()
	if err != nil {
		t.Error(err)
	}

	t.Log(ret.InUse, ret.Next)
}
