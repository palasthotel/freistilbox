package help

import "fmt";
import "github.com/palasthotel/freistilbox/commands";

type Cmd struct{}

func (_ Cmd) Command() string {
	return "help"
}

func (_ Cmd) Description() string {
	return "shows the help screen";
}

func (_ Cmd) Help(args []string) {
	fmt.Println("It seems you're stupid. You just executed freistilbox help help.");
}

func (_ Cmd) Execute(args []string) {
	var cmds=commands.List;
	fmt.Println("freistilbox - administer freistil setups more easily");
	fmt.Println("");
	fmt.Println("supported commands:");
	for _,command := range cmds {
		fmt.Printf("* %s - %s\n",command.Command(),command.Description());
	}
}