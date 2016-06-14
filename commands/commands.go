package commands

import (
	"os"
)

var List = []Command{
}

func Run() {
	var args []string=os.Args[1:];
	if(len(args) == 0) {
		args=[]string{"help"}
	}
	for _,cmd := range List {
		if(cmd.Command() == args[0]) {
			var cmdArgs=args[1:];
			cmd.Execute(cmdArgs);
		}
	}
}