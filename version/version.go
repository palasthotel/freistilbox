package version

import (
	"fmt"
);

type Cmd struct{}

func (_ Cmd) Command() string {
	return "version"
}

func (_ Cmd) Description() string {
	return "shows the version info";
}

func (_ Cmd) Help(args []string) {
	fmt.Println("There is no help available for version");
}

func (_ Cmd) Execute(args []string) {
	fmt.Println("freistilbox version 1.0-dev");
}