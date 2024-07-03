package gomer

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Variable struct {
	Type       string
	Name       string
	Visibility string
	IsStatic   bool
	IsConst    bool
}

func NewVariable(_type string, _name string, _vis string, _static bool, _const bool) *Variable {
	return &Variable{
		Type:       _type,
		Name:       _name,
		Visibility: _vis,
		IsStatic:   _static,
		IsConst:    _const,
	}
}

type Function struct {
	Type       string
	Name       string
	Args       []Variable
	Visibility string
	IsStatic   bool
	IsConst    bool
}

type Enum struct {
	Type string
	Name string
}

type Template struct {
	Type       string
	Parameters string
}

type Class struct {
	Name      string
	Variables []Variable
	Functions []Function
	IsStruct  bool
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseClassBody(body *string, currentClass *Class) {
	var rgx_ACCESS string = `^\s*(?:public|private|protected):`
	var rgx_DECLFUNC string = `((?:\w+(?:::\w+)*(?:<(?:[^<>]|(?1))*>)?(?:\s*[*&])?\s+)+)(\w+)\s*\((.*?)\)(\s*const)?`
	var rgx_DECLVAR string = `((?:\w+\s+)*)((?:[\w_:]|<(?:[^<>]|(?2))*>|,\s*)+)( *[*&] +| +[*&] *| +)*([A-Za-z_]\w*)`

	re_ACCESS := regexp.MustCompile(rgx_ACCESS)
	re_DECLFUNC := regexp.MustCompile(rgx_DECLFUNC)
	re_DECLVAR := regexp.MustCompile(rgx_DECLVAR)

	fmt.Printf("Body of %s\n-------\n%s\n", currentClass.Name, *body)

	var currentAccess string = "private"
	if currentClass.IsStruct {
		currentAccess = "public"
	}

	scanner := bufio.NewScanner((strings.NewReader(*body)))
	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}

		access := strings.TrimSpace(re_ACCESS.FindString(line))
		if access != "" {
			currentAccess = access
			continue
		}

		function := re_DECLFUNC.FindString(line)
		if function != "" {
			// parse Function
			var f Function
			f.Visibility = currentAccess
			continue
		}

		variable := re_DECLVAR.FindString(line)
		if variable != "" {
			// parse Variable
			var v Variable
			v.Visibility = currentAccess
			continue
		}

	}

	// match variables
	for _, match := range re_DECLVAR.FindAllStringSubmatch(*body, -1) {
		fmt.Printf("%s, %s\n", match[0], match[1])
	}

}

// naive way to do this parse -- insane assumptions for now:
// - the file isn't too large so it reads in whole file at once... TODO: chunked reading
// - 1 expression per line...  TODO: regex-less parser
func Parse(fileName string) {
	fd, err := os.ReadFile(fileName)
	check(err)

	// Groups
	// 1 - template;  	2 - template params; 3 - class/struct;
	// 4 - class name; 	5 - inheritance;  	 6 - body
	var rgxp_CLASS string = `(?s)(\s*template\s+<((?:typename|class)\s+\w+\s*,?\s*)+>)?\s*(class|struct)\s+(\w+)\s*(:(?:\s*\w+\s+\w+\s*,?)*)?{([ -~\s]+?)};`
	re_CLASS := regexp.MustCompile(rgxp_CLASS)
	for _, match := range re_CLASS.FindAllStringSubmatch(string(fd), -1) {
		var currentClass Class
		if match[3] == "struct" {
			currentClass.IsStruct = true
		}
		currentClass.Name = match[4]
		parseClassBody(&match[len(match)-1], &currentClass)
	}
}
