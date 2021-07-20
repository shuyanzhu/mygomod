package Hi2

import (
	"fmt"
)

func init() {
	fmt.Println(h1)
}

var h2 = h1 + 1

func Hi2(name string) string {
	fmt.Println(h1, h2)
	return fmt.Sprintf("Hi, %s", name)
}
