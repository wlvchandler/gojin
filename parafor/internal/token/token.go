package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL    = "ILLEGAL"
	EOF        = "EOF"
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"
	STRING     = "STRING"
	CONST      = "CONST"

	COLON = ":"

	NAMESPACE   = "NAMESPACE"
	CLASS       = "CLASS"
	INTERFACE   = "INTERFACE"
	FIELD       = "FIELD"
	METHOD      = "METHOD"
	PARAM       = "PARAM"
	CONSTRUCTOR = "CONSTRUCTOR"
	INHERIT     = "INHERIT"
	ENUM        = "ENUM"
	GENERIC     = "GENERIC"
	PATTERN     = "PATTERN"

	NAME         = "NAME"
	TYPE         = "TYPE"
	ACCESS       = "ACCESS"
	STATIC       = "STATIC"
	FINAL        = "FINAL"
	ABSTRACT     = "ABSTRACT"
	TRANSIENT    = "TRANSIENT"
	VOLATILE     = "VOLATILE"
	SYNCHRONIZED = "SYNCHRONIZED"
	EXTENDS      = "EXTENDS"
	CHILD        = "CHILD"
	PARENT       = "PARENT"
	VALUE        = "VALUE"

	YES = "Y"
	NO  = "N"
)

var Keywords = map[string]TokenType{
	"namespace":    NAMESPACE,
	"class":        CLASS,
	"interface":    INTERFACE,
	"field":        FIELD,
	"method":       METHOD,
	"param":        PARAM,
	"constructor":  CONSTRUCTOR,
	"inherit":      INHERIT,
	"enum":         ENUM,
	"generic":      GENERIC,
	"pattern":      PATTERN,
	"name":         NAME,
	"type":         TYPE,
	"access":       ACCESS,
	"static":       STATIC,
	"final":        FINAL,
	"abstract":     ABSTRACT,
	"volatile":     VOLATILE,
	"synchronized": SYNCHRONIZED,
	"extends":      EXTENDS,
	"child":        CHILD,
	"parent":       PARENT,
	"value":        VALUE,
	"Y":            YES,
	"N":            NO,
}

func IdentifierLookup(ident string) TokenType {
	if token, ok := Keywords[ident]; ok {
		return token
	}
	return IDENTIFIER
}
