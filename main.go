package main

import (
	"fmt"
	"os"
	"todofinder/finder"
)

func main() {
	var wd string
	var err error

	args := os.Args[1:]

	// Determines directory to scan, defaults to current working directory
	if len(args) > 0 {
		wd = args[0]
	} else {
		wd, err = os.Getwd()
		if err != nil {
			panic(fmt.Sprintf("failed to get working dir: %s", err.Error()))
		}
	}

	todoFinder := &finder.TodoFinderImpl{}
	todos, err := todoFinder.Find(wd)
	if err != nil {
		panic(fmt.Sprintf("failed to find todos: %s", err.Error()))
	}

	for _, todo := range todos {
		fmt.Println(todo.ToString())
	}
}
