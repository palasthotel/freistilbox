package api

import (
	"fmt"
);

type Cmd struct{}

func (_ Cmd) Command() string {
	return "api"
}

func (_ Cmd) Description() string {
	return "provides direct access to the api";
}

func (_ Cmd) Help(args []string) {
	fmt.Println("usage: freistilbox api PATH");
	fmt.Println("");
	fmt.Println("for example: freistilbox api clusters")
}

func (_ Cmd) Execute(args []string) {
	if len(args) != 1 {
		Cmd{}.Help([]string{});
		return
	}
	fmt.Println(string(Perform(args[0])));
}