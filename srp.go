package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "package main\n\nimport (    \"fmt\"\n    \"strings\"\n)\nfunc main() {\n    s := \"%s\"\n    fmt.Println(s, fmt.Sprintf(s, strings.Replace(s, \"\\n\" \"\\n\", -1)))\n}\n"
	fmt.Println(fmt.Sprintf(s, strings.Replace(s, "\n", "\\n", -1)))
}
