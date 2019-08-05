package websites

import (
	"fmt"
	"github.com/palasthotel/freistilbox/api"
	"encoding/json"
	"strconv"
)

type Update struct{}

func(_ Update) Command() string {
	return "update-website";
}

func(_ Update) Description() string {
	return "updates an existing website";
}

func(_ Update) Help(args []string) {
	fmt.Println("usage: freistilbox update-website SITEID [--domain new_main_domain] [--env new_env] [--webmaster new_webmaster] [--branch new_branch] [--php new_php_version]");
}

func performupdates(siteid string, args []string) {
	dataset := Website_Update {}
	for i := 1; i < len(args); i+=2 {
		if (args[i] == "--domain") {
			dataset.Update_Maindomain = &Update_Maindomain{ Main_domain: args[i+1]};

		} else if (args[i] == "--env") {
			dataset.Update_env = &Update_env{Environment: args[i+1]};
		} else if (args[i] == "--webmaster") {
			dataset.Update_webmaster = &Update_webmaster{Webmaster:args[i+1]};
		} else if (args[i] == "--branch") {
			dataset.Update_branch = &Update_branch{Branch:args[i+1]};
		} else if (args[i] == "--php") {
			dataset.Update_php = &Update_php{Php:args[i+1]};
		} else {
			fmt.Println("unknown argument: ",args[i]);
			return;
		}
	}
	if dataset.Update_env == nil && dataset.Update_Maindomain == nil && dataset.Update_webmaster == nil && dataset.Update_branch == nil && dataset.Update_php == nil {
		fmt.Println("Nothing to set. Aborting");
		return;
	}
	encoded,_ := json.Marshal(dataset);
	fmt.Println(string(api.Put(fmt.Sprintf("websites/%s",siteid),encoded)));
}

func(_ Update) Execute(args []string) {
	//first things first: we have to get all clusters.
	clusters := api.FetchClusters();
	website_no,_ := strconv.Atoi(args[0]);
	for _,cluster := range clusters {
		for _,website := range cluster.Relationships.Websites.Data {
			fetchedWebsite := api.FetchWebsite(website.Id);
			if(fetchedWebsite.Attributes.Siteno == website_no) {
				performupdates(website.Id,args);
				return;
			}
		}
	}
}