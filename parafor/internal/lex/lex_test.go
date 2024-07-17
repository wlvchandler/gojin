package lex

import (
	"fmt"
	"localhost/wlvchandler/parafor/internal/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `class name:User
field class:User name:id type:int access:private
method class:User name:getName returnType:string access:public
pattern type:singleton class:User`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.CLASS, "class"},
		{token.NAME, "name"},
		{token.COLON, ":"},
		{token.IDENTIFIER, "User"},
		{token.FIELD, "field"},
		{token.CLASS, "class"},
		{token.COLON, ":"},
		{token.IDENTIFIER, "User"},
		{token.NAME, "name"},
		{token.COLON, ":"},
		{token.IDENTIFIER, "id"},
		{token.TYPE, "type"},
		{token.COLON, ":"},
		{token.IDENTIFIER, "int"},
		{token.ACCESS, "access"},
		{token.COLON, ":"},
		{token.IDENTIFIER, "private"},
		{token.METHOD, "method"},
		{token.CLASS, "class"},
		{token.COLON, ":"},
		{token.IDENTIFIER, "User"},
		{token.NAME, "name"},
		{token.COLON, ":"},
		{token.IDENTIFIER, "getName"},
		{token.IDENTIFIER, "returnType"},
		{token.COLON, ":"},
		{token.IDENTIFIER, "string"},
		{token.ACCESS, "access"},
		{token.COLON, ":"},
		{token.IDENTIFIER, "public"},
		{token.PATTERN, "pattern"},
		{token.TYPE, "type"},
		{token.COLON, ":"},
		{token.IDENTIFIER, "singleton"},
		{token.CLASS, "class"},
		{token.COLON, ":"},
		{token.IDENTIFIER, "User"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		_token := l.NextToken()
		fmt.Printf("tests[%d]: '%s'\n", i, _token)
		if _token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. Expected=%q\tFound=%q", i, tt.expectedType, _token.Type)
		}
		if _token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. Expected=%q\tFound=%q", i, tt.expectedLiteral, _token.Literal)
		}
	}
}

func TestNextTokenWithNumbers(t *testing.T) {
	input := `field class:User name:age type:int value:30`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.FIELD, "field"},
		{token.CLASS, "class"},
		{token.COLON, ":"},
		{token.IDENTIFIER, "User"},
		{token.NAME, "name"},
		{token.COLON, ":"},
		{token.IDENTIFIER, "age"},
		{token.TYPE, "type"},
		{token.COLON, ":"},
		{token.IDENTIFIER, "int"},
		{token.VALUE, "value"},
		{token.COLON, ":"},
		{token.INT, "30"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		_token := l.NextToken()
		fmt.Printf("tests[%d]: '%s'\n", i, _token)
		if _token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. Expected=%q\tFound=%q", i, tt.expectedType, _token.Type)
		}
		if _token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. Expected=%q\tFound=%q", i, tt.expectedLiteral, _token.Literal)
		}
	}
}

func TestParseIntoTokens(t *testing.T) {
	input := `field class:User name:age type:int value:30`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.FIELD, "field"},
		{token.CLASS, "class"},
		{token.COLON, ":"},
		{token.IDENTIFIER, "User"},
		{token.NAME, "name"},
		{token.COLON, ":"},
		{token.IDENTIFIER, "age"},
		{token.TYPE, "type"},
		{token.COLON, ":"},
		{token.IDENTIFIER, "int"},
		{token.VALUE, "value"},
		{token.COLON, ":"},
		{token.INT, "30"},
		{token.EOF, ""},
	}

	l := New(input)

	var lineTokens []*token.Token
	for l.readPosition <= len(input) {
		tok := l.NextToken()
		lineTokens = append(lineTokens, &tok)
	}

	for _, tok := range lineTokens {
		fmt.Printf("%q\n", *tok)
	}
	for _, tok := range tests {
		fmt.Printf("%q\n", tok)
	}

	if len(lineTokens) != len(tests) {
		t.Fatalf("Token count mismatch. Expected=%d\tFound=%d", len(lineTokens), len(tests))
	} else {
		for i, tt := range tests {
			fmt.Printf("lineTokens[%d] - `%s`\n", i, lineTokens[i])
			if lineTokens[i].Type != tt.expectedType {
				t.Fatalf("tests[%d] - tokentype wrong. Expected=%q\tFound=%q", i, tt.expectedType, lineTokens[i].Type)
			}
		}
	}

}
