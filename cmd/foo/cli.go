package main

import (
	"encoding/json"
	"fmt"
	"runtime/debug"

	"github.com/peterbourgon/ff/v4"
)

func Run() {
	info, _ := debug.ReadBuildInfo()
	// fmt.Println(info, ok)
	by, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(by))
	_ = ff.Command{}
	fmt.Println("Hello, World!")
}
