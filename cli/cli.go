package cli

import (
	"fmt"

	"github.com/peterbourgon/ff/v4"
)

func Run() {
	_ = ff.Command{}
	fmt.Println("Hello, World!")
}
