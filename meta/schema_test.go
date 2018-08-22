/**
 *  author: lim
 *  data  : 18-8-22 下午10:04
 */

package meta

import "testing"

func TestNewMeta(t *testing.T) {
	meta, err := NewMeta("../etc/meta.json")
	if err != nil {
		t.Error(err)
	}
	t.Log(meta.schemas)
}
