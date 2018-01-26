// - Give the order of execution of the following code
// - What is the output of the program?
// - Write unit tests for this program
package main

import "fmt"

var WhatIsThe = AnswerToLife()

func AnswerToLife() int {
	return 42
}

func init() {
	WhatIsThe = 0
}

func init() {
	WhatIsThe = 10
}

func init() {
	WhatIsThe = 5
}

func main() {
	fmt.Println(WhatIsThe)
}
