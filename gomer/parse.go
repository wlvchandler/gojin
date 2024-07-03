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

func parseClassBody(body string, currentClass *Class) error {
	const (
		rgx_ACCESS   = `^\s*(?:public|private|protected):`
		rgx_DECLFUNC = `((?:\w+(?:::\w+)*(?:<(?:[^<>]|(?1))*>)?(?:\s*[*&])?\s+)+)(\w+)\s*\((.*?)\)(\s*const)?`
		rgx_DECLVAR  = `((?:\w+\s+)*)((?:[\w_:]|<(?:[^<>]|(?2))*>|,\s*)+)( *[*&] +| +[*&] *| +)*([A-Za-z_]\w*)`
	)
	re_ACCESS := regexp.MustCompile(rgx_ACCESS)
	re_DECLFUNC := regexp.MustCompile(rgx_DECLFUNC)
	re_DECLVAR := regexp.MustCompile(rgx_DECLVAR)

	var currentAccess string = "private"
	if currentClass.IsStruct {
		currentAccess = "public"
	}

	scanner := bufio.NewScanner((strings.NewReader(body)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		if access := strings.TrimSpace(re_ACCESS.FindString(line)); access != "" {
			currentAccess = access
			continue
		}

		if function := re_DECLFUNC.FindString(line); function != "" {
			f := Function{Visibility: currentAccess}
			_ = f
			continue
		}

		variable := re_DECLVAR.FindString(line)
		if variable != "" {
			// parse Variable
			v := Variable{Visibility: currentAccess}
			_ = v
			continue
		}

	}

	return nil
}

/*
naive way to do this parse -- insane assumptions for now:

  - the file isn't too large so it reads in whole file at once... TODO: chunked reading

  - one expression per line...  TODO: regex-less parser

    Regx groups for CLASS:
    1 - template;  	2 - template params; 3 - class/struct;
    4 - class name; 5 - inheritance;  	 6 - body
*/
func Parse(fileName string) error {
	fd, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("read error: %w", err)
	}

	const rgxp_CLASS = `(?s)(\s*template\s+<((?:typename|class)\s+\w+\s*,?\s*)+>)?\s*(class|struct)\s+(\w+)\s*(:(?:\s*\w+\s+\w+\s*,?)*)?{([ -~\s]+?)};`
	re_CLASS := regexp.MustCompile(rgxp_CLASS)
	for _, match := range re_CLASS.FindAllStringSubmatch(string(fd), -1) {
		currentClass := Class{
			Name:     match[4],
			IsStruct: match[3] == "struct",
		}
		parseClassBody(match[len(match)-1], &currentClass)
	}

	return nil
}
