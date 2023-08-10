/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"lukehagar/tmi/root"
	"os"
)

func main() {
	rootCmd := root.NewRootCmd()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
