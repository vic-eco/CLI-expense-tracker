package main

import (
	"expense-tracker/cmd"
	"fmt"
)

func main() {
	rootCmd := cmd.NewRootCMD()

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
