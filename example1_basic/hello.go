package example1_basic

import "fmt"

const (
	GenericGreeting = "Hello Dude!"
)

func hello(user string) string {
	if len(user) == 0 {
		return GenericGreeting
	} else {
		return fmt.Sprintf("Hello %v!", user)
	}
}
