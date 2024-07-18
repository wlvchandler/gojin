package ast

import (
	"localhost/wlvchandler/parafor/internal/token"
)

type Node struct {
	Token    *token.Token
	Children []*Node
}

type Node_Class struct {
	m_name  string
	m_vars  []*token.Token
	m_funcs []*token.Token
	m_patts []*token.Token
}

type Node_Field struct {
	m_class *Node_Class
	m_name  *token.Token
	m_type  *token.Token
	m_value *token.Token
}

func CreateRootNode() *Node {
	return &Node{}
}

func New(node_tok token.Token) *Node {
	n := &Node{Token: &node_tok}
	return n
}

// NODE CREATION

func CreateNode(line []*token.Token) *Node {

	switch line[0].Type {
	case token.CLASS:

		CreateNode_Class()
	case token.FIELD:
		CreateNode_Field()
	case token.METHOD:
		CreateNode_Method()

	}

	return &Node{}
}

// func CreateNode_Class(name string, vars []*token.Token, funcs []*token.Token, patterns []*token.Token) *Node_Class {
func CreateNode_Class(line []*token.Token) *Node_Class {
	var n Node_Class
	for i, tok := range line {
		switch tok.Type {
		case token.NAME:
			n.m_name = line[i+1].Literal
		}
	}
	n.m_class = nil // TODO
	return n
	return &Node_Class{
		m_name:  name,
		m_vars:  vars,
		m_funcs: funcs,
		m_patts: patterns,
	}
}

func CreateNode_Field(line []*token.Token) Node_Field {
	var n Node_Field
	for i, tok := range line {
		switch tok.Type {
		case token.NAME:
			n.m_name = line[i+1]
		case token.TYPE:
			n.m_type = line[i+1]
		case token.VALUE:
			n.m_value = line[i+1]
		}
	}
	n.m_class = nil // TODO
	return n
}

func CreateNode_Pattern() *Node {
	return &Node{}
}

func CreateNode_String() *Node {
	return &Node{}
}

func CreateNode_Int() *Node {
	return &Node{}
}

func IsLeafNode(node Node) bool {
	return true
}

// FIND NODES IN TREE

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
