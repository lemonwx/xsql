/**
 *  author: lim
 *  data  : 18-3-24 下午4:36
 */

package node

import "fmt"

type Node struct {
	addr     string
	user     string
	password string
	Db string
}

func NewNode(host string, port int, user, password, db string) *Node {

	node := &Node {
		addr : fmt.Sprintf("%s:%d", host, port),
		user: user,
		password: password,
		Db: db,
	}

	return node
}

func (node *Node) Connect() error {
	return nil
}
