package main

import (
	"fmt"
	"github.com/gelleson/gcsv/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}
