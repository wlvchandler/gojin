package parafor

import (
	"fmt"
	"localhost/wlvchandler/parafor/internal/lex"
)

func Main() {
	l := lex.New("class name:X\n")

	fmt.Println(l)
}
