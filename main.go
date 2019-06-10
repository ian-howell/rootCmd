package main

import (
	"github.com/ian-howell/rootCmd/cmd"
	barcmd "github.com/ian-howell/barCommand/cmd"
	foocmd "github.com/ian-howell/fooCommand/cmd"
)

func main() {
	rootCmd := cmd.NewRootCommand()
	rootCmd.AddCommand(foocmd.NewFooCommand())
	rootCmd.AddCommand(barcmd.NewFooCommand())

	rootCmd.Execute()
}
