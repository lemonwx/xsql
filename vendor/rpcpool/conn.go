/**
 *  author: lim
 *  data  : 18-4-18 下午7:09
 */

package rpcpool

import "net/rpc"

type Conn interface {
	Call(serviceMethod string, args interface{}, reply interface{}) error
	Close() error
}

type Client struct {
	cli *rpc.Client
}

func (cli *Client) Call(serviceMethod string, args interface{}, reply interface{}) error {
	return cli.Call(serviceMethod, args, reply)
}