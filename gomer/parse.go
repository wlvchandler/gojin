package gomer

import (
	"fmt"
	"os"
	"regexp"
)

type Variable struct {
	Type       string
	Name       string
	Visibility string
	isStatic   bool
	isConst    bool
}

type Function struct {
	Type       string
	Name       string
	Args       []Variable
	Visibility string
	isStatic   bool
	isConst    bool
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
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// naive way to do this parse -- assumption is the file isn't too large so it reads in whole file at once...
// TODO: chunked reading
func Parse(fileName string) {
	fd, err := os.ReadFile(fileName)
	check(err)

	var rgxp_CLASS string = `(?s)(?:class|struct)\s+(\w+)\s*(:(?:\s*\w+\s+\w+\s*,?)*)?{([ -~\s]+?)};`
	re_CLASS := regexp.MustCompile(rgxp_CLASS)
	for i, match := range re_CLASS.FindAllStringSubmatch(string(fd), -1) {
		for j, cg := range match {
			fmt.Printf("%d.%d-----------\n`%s`\n", i, j, cg)
		}
	}

	/*
		fd, err := os.Open(fileName)
		check(err)
		defer fd.Close()

		var i int = 0
		scanner := bufio.NewScanner(fd)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Printf("%d:\t%s\n", i, line)
			i++
		}

		re, err := regexp.Compile(`\d+`)
		if err != nil {
			return
		}

		var x = re.FindString("skld234jf")
		fmt.Println(x)
	*/
	//fmt.Println(currentLine)
}
