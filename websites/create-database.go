package websites

import (
	"fmt"
	"github.com/palasthotel/freistilbox/api"
	_ "encoding/json"
	"strconv"
);

type CreateDB struct {}

func (_ CreateDB) Command() string {
	return "create-database"
}

func (_ CreateDB) Description() string {
	return "create a new database";
}

func (_ CreateDB) Help(args []string) {
	fmt.Println("usage: freistilbox create-database WEBSITE");
}

func (cmd CreateDB) Execute(args []string) {
		clusters := api.FetchClusters();
	website_no,_ := strconv.Atoi(args[0]);
	for _,cluster := range clusters {
		for _,website := range cluster.Relationships.Websites.Data {
			fetchedWebsite := api.FetchWebsite(website.Id);
			if(fetchedWebsite.Attributes.Siteno == website_no) {
				fmt.Println(string(api.Post(fmt.Sprintf("websites/%s/databases",website.Id),make([]byte,0))));
				return;
			}
		}
	}
}