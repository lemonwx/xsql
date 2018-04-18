/**
 *  author: lim
 *  data  : 18-4-10 下午10:16
 */

package version


/*
import (
	"errors"
	"net/rpc"
	"sync"
	"rpcpool"
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

var pool rpcpool.Pool

func NewRpcPool(size int, addr string) {
	pool = rpcpool.Pool{
		MaxIdle:     100,
		MaxActive:   10000,
		Dial: func() (rpcpool.Conn, error) {
			return rpc.DialHTTP("tcp", addr)
		},
	}
}

func NextVersion() ([]byte, error) {
	cli := pool.Get()
	if cli == nil {
		return nil, GET_VERSION_CONN_FAILED
	}
	defer cli.Close()

	var nextVer []byte
	err := cli.Call("VSeq.NextV", uint8(0), &nextVer)
	if err != nil {
		return nil, err
	}

	return nextVer, nil
}

func ReleaseVersion(version []byte) error {
	cli := pool.Get()
	if cli == nil {
		return GET_VERSION_CONN_FAILED
	}
	defer cli.Close()

	var ret bool
	err := cli.Call("VSeq.Release", version, &ret)
	if err != nil {
		return err
	}
	if ret {
		return nil
	} else {
		return RELEASE_FAILED
	}
}

func VersionsInUse() (map[string]uint8, error) {
	cli := pool.Get()
	if cli == nil {
		return nil, GET_VERSION_CONN_FAILED
	}
	defer cli.Close()

	var vInuse map[string]uint8
	err := cli.Call("VSeq.VInUser", uint8(0), &vInuse)
	if err != nil {
		panic(err)
	}
	return vInuse, nil
}

*/