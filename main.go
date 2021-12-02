package main

import (
	"os"

	"github.com/shayd3/advent-of-code-2021/days"
)

var funcMap = map[string]interface{}{
	"1": days.Day01,
}

func main() {
	runDay(os.Args[1])
}

func runDay(day string) {
	funcMap[day].(func())()
}