package Hi2

import (
	"fmt"
)

var h2 = h1

func Hi2(name string) string {
	fmt.Println(h1, h2)
	return fmt.Sprintf("Hi, %s", name)
}
