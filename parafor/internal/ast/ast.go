package ast

import "localhost/wlvchandler/parafor/internal/token"

type Node struct {
	Token    *token.Token
	Children []*Node
}

func CreateRootNode() *Node {
	return &Node{}
}

func New(node_tok token.Token) *Node {
	n := &Node{Token: &node_tok}
	return n
}

func IsLeafNode(node Node) bool {
	return true
}

func CreateNode_Class() *Node {
	return &Node{}
}

func findNode(name string, nodeType token.TokenType, root *Node) *Node {
	var foundNode *Node
	foundNode = nil
	if root.Token.Type == nodeType && root.Token.Literal == name {
		foundNode = root
	} else {
		for _, n := range root.Children {
			foundNode = findNode(name, nodeType, n)
			if foundNode != nil {
				break
			}
		}
	}
	return foundNode
}

func FindNode_Class(name string, root *Node) *Node {
	return findNode(name, token.CLASS, root)
}
