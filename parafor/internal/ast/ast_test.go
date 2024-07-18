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

func TestASTCreateField(t *testing.T) {
	n1 := token.Token{Type: token.FIELD, Literal: "field"}
	n2 := token.Token{Type: token.CLASS, Literal: "class"}
	n3 := token.Token{Type: token.IDENTIFIER, Literal: "User"}
	n4 := token.Token{Type: token.NAME, Literal: "name"}
	n5 := token.Token{Type: token.IDENTIFIER, Literal: "age"}
	n6 := token.Token{Type: token.TYPE, Literal: "type"}
	n7 := token.Token{Type: token.IDENTIFIER, Literal: "int"}
	n8 := token.Token{Type: token.VALUE, Literal: "value"}
	n9 := token.Token{Type: token.INT, Literal: "30"}

	var line = []*token.Token{
		&n1, &n2, &n3, &n4, &n5, &n6, &n7, &n8, &n9,
	}

	f := CreateNode_Field(line)

	if f.m_class != nil {
		t.Fatalf("Field class must be nil\n")
	}
	if f.m_name.Literal != n5.Literal {
		t.Fatalf("Field name incorrect. Expected=%s\tFound=%s\n", n5.Literal, f.m_name.Literal)
	}

	if f.m_type.Literal != n7.Literal {
		t.Fatalf("Field type incorrect. Expected=%s\tFound=%s\n", n7.Literal, f.m_type.Literal)
	}
	if f.m_value.Literal != n9.Literal {
		t.Fatalf("Field value incorrect. Expected=%s\tFound=%s\n", n9.Literal, f.m_value.Literal)
	}

}

func TestASTFindClassNode(t *testing.T) {
	/*
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
	*/
}
