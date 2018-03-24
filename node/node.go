/**
 *  author: lim
 *  data  : 18-3-24 下午4:36
 */

package node

type Node struct {
	addr     string
	user     string
	password string
}

func NewNode(addr string) *Node {
	node := new(Node)
	node.connect()

	return node
}

func (node *Node) connect() {

}
