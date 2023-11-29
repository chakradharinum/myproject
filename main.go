package main

import (
	"flag"
	"myproject/array"
	"myproject/condition"
	"myproject/forloop"
	"myproject/switchcase"
)

func main() {
	conditionFlag := flag.Bool("if", false, "Run if example")
	forloopFlag := flag.Bool("for", false, "Run for example")
	switchcaseFlag := flag.Bool("switch", false, "Run switch example")
	arrayFlag := flag.Bool("array", false, "Run array example")
	flag.Parse()

	if *conditionFlag {
		condition.RunIfExample()
	}
	if *forloopFlag {
		forloop.RunForExample()
	}
	if *switchcaseFlag {
		switchcase.RunSwitchExample()
	}
	if *arrayFlag {
		array.RunArrayExample()
	}
}
