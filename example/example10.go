// test project main.go
package main

import (
	"fmt"
)

func findEvildoer(name string) string {
	var evildoer string
	for _, r := range name {
		switch {
		case r >= '\u0041' && r <= '\u005a': // a-z
		case r >= '\u0061' && r <= '\u0071': //A-Z
		case r >= '\u4e00' && r <= '\u9fbf': //�����ַ�
			goto L3
		default:
			evildoer = string(r)
			goto L2
		}
	}

L2:
	fmt.Printf("The frist evildoer of name '%s' is '%s'\n", name, evildoer)
L3:
	return evildoer
}

func main() {
	findEvildoer("kkkkk@@@@")
}