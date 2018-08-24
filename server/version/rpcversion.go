/**
 *  author: lim
 *  data  : 18-4-10 下午10:16
 */

package version

import (
	"errors"
	"net/rpc"
	"rpcpool"
	"sync"

	"github.com/lemonwx/VSequence/base"
)

const (
	DIALERR = iota
	USEING
	FREE
	CLOSED
)

var GET_VERSION_CONN_FAILED error = errors.New("GET VERSION CONN FAILED")
var RELEASE_FAILED error = errors.New("RELEASE USED VERSION FAILED ")
var lock sync.Mutex

var pool *rpcpool.ConnPool

func NewRpcPool(initSize, maxSize int, addr string) {
	var err error
	pool, err = rpcpool.NewConnPool(
		func() (*rpc.Client, error) {
			return rpc.DialHTTP("tcp", addr)
		},
		initSize,
		maxSize,
	)
	if err != nil {
		panic(err)
	}
}

func NextVersion() (uint64, error) {
	cli, err := pool.Get()
	if err != nil {
		return 0, GET_VERSION_CONN_FAILED
	}
	defer pool.Put(cli)

	var nextVer uint64
	err = cli.Call("VSeq.NextV", uint8(0), &nextVer)
	if err != nil {
		return 0, err
	}

	return nextVer, nil
}

func ReleaseVersion(version uint64) error {
	cli, err := pool.Get()
	if err != nil {
		return GET_VERSION_CONN_FAILED
	}
	defer pool.Put(cli)

	var ret bool
	err = cli.Call("VSeq.Release", version, &ret)
	if err != nil {
		return err
	}
	if ret {
		return nil
	} else {
		return RELEASE_FAILED
	}
}

func VersionsInUse() (map[uint64]uint8, error) {
	cli, err := pool.Get()
	if err != nil {
		return nil, GET_VERSION_CONN_FAILED
	}
	defer pool.Put(cli)

	var vInuse map[uint64]uint8
	err = cli.Call("VSeq.VInUser", uint8(0), &vInuse)
	if err != nil {
		return nil, err
	}
	return vInuse, nil
}

func InUseAndNext() (*base.UseAndNext, error) {
	cli, err := pool.Get()
	if err != nil {
		return nil, GET_VERSION_CONN_FAILED
	}
	defer pool.Put(cli)

	useAndNext := base.UseAndNext{}
	err = cli.Call("VSeq.InUseAndNext", uint8(0), &useAndNext)
	if err != nil {
		return nil, err
	}
	return &useAndNext, nil
}
