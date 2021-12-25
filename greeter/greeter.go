package greeter

import (
	"fmt"
)

func Greet(name string) string {
	message := fmt.Sprintf("Yo, %v. Welcome!", name)
	return message
}
