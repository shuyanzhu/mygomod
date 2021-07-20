package Hi2

import (
	"fmt"
)

func init() {
	fmt.Println("In Hi3: ")
	fmt.Println(h2)
}

var h1 = 1

func Hi3(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}
