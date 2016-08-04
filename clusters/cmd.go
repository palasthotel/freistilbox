package clusters

import (
	"fmt"
	"github.com/palasthotel/freistilbox/api"
);

type Cmd struct{}

func (_ Cmd) Command() string {
	return "clusters"
}

func (_ Cmd) Description() string {
	return "shows a list of all available clusters";
}

func (_ Cmd) Help(args []string) {
	fmt.Println("usage: freistilbox clusters");
}

func (_ Cmd) Execute(args []string) {
	var clusters=api.Clusters()
	for _,value := range clusters {
		fmt.Printf("%s: %s (%d Sites)\n",value.Attributes.Name,value.Attributes.Description,value.Relationships.Websites.Data.Len());
	}

}
