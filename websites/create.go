package websites

import (
	"fmt"
	"github.com/palasthotel/freistilbox/api"
	"encoding/json"
);

type Create struct{}

func (_ Create) Command() string {
	return "create-website"
}

func (_ Create) Description() string {
	return "create a new website";
}

func (_ Create) Help(args []string) {
	fmt.Println("usage: freistilbox create-website CLUSTER DOMAIN ENV WEBMASTER BRANCH");
}

func (_ Create) Execute(args []string) {
	cluster := args[0];
	domain := args[1];
	env := args[2];
	webmaster := args[3];
	repository := args[4];
	clusters := api.FetchClusters();
	var foundCluster api.Cluster;
	for _,value := range clusters {
		if (value.Attributes.Name == cluster) {
			foundCluster = value;
		}
	}
	fmt.Printf("Cluster: %s (%s)\nDomain: %s\nenv: %s\nwebmaster: %s\nbranch: %s\n",cluster,foundCluster.Id,domain,env,webmaster,repository);
	dataset := Payload {}
	dataset.Website.Cluster_id = foundCluster.Id;
	dataset.Website.Main_domain = domain;
	dataset.Website.Environment = env;
	dataset.Website.Webmaster = webmaster;
	dataset.Website.Repository_branch = repository;

	encoded,_ := json.Marshal(dataset);
	fmt.Println(string(api.Post("websites",encoded)));
}
