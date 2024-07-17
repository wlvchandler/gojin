package ast

import (
	"localhost/wlvchandler/parafor/internal/token"
	"testing"
)

func TestASTCreateRootNode(t *testing.T) {
	root := CreateRootNode()
	if root.Token != nil {
		t.Fatalf("Root internal token must be nil\n")
	}
}

func TestASTFindClassNode(t *testing.T) {
	root := CreateRootNode()

	// field class:User name:age type:int value:30
	n1 := New(token.Token{Type: token.FIELD, Literal: "field"})
	n2 := New(token.Token{Type: token.CLASS, Literal: "class"})
	n3 := New(token.Token{Type: token.IDENTIFIER, Literal: "User"})
	n4 := New(token.Token{Type: token.NAME, Literal: "name"})
	n5 := New(token.Token{Type: token.IDENTIFIER, Literal: "age"})
	n6 := New(token.Token{Type: token.TYPE, Literal: "type"})
	n7 := New(token.Token{Type: token.IDENTIFIER, Literal: "int"})
	n8 := New(token.Token{Type: token.VALUE, Literal: "value"})
	n9 := New(token.Token{Type: token.INT, Literal: "30"})

	if root.Token != nil {
	}
}
