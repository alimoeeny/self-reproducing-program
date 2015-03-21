package main

import "fmt"

func main() {
	s := "package main\n\nimport \"fmt\"\n\nfunc main() {\n\ts := %q\n\tfmt.Println(fmt.Sprintf(s, s))\n}"
	fmt.Println(fmt.Sprintf(s, s))
}
