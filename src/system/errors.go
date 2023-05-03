package system

import "fmt"

func PanicOnError(err error) {
	if err != nil {
		panic(fmt.Sprintf("an error occurred: %s", err))
	}
}
