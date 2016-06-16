package credentials

import (
	"fmt"
);

type Cmd struct{}

func (_ Cmd) Command() string {
	return "credentials"
}

func (_ Cmd) Description() string {
	return "sets your freistilbox dashboard credentials";

}

func (_ Cmd) Help(args []string) {
	fmt.Println("usage: freistilbox credentials USER TOKEN");
	fmt.Println("without USER and TOKEN: see your current credentials");
}

func (_ Cmd) Execute(args []string) {
	if( len(args)==2) {
		var c=Credentials{User:args[0],Token:args[1]};
		c.Save();
	} else {
		var c=Load();
		fmt.Printf("%s %s\n",c.User,c.Token);
	}
}