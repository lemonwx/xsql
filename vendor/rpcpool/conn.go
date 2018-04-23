/**
 *  author: lim
 *  data  : 18-4-18 下午7:09
 */

package rpcpool

import "net/rpc"

type Conn struct {
	cli *rpc.Client
}

func (conn *Conn) Call(serviceMethod string, args interface{}, reply interface{}) error {
	return conn.cli.Call(serviceMethod, args, reply)
}

func (conn *Conn) Close() error {
	return conn.cli.Close()
}
