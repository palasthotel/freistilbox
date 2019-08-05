package websites

type Website_Set struct {
	Cluster_id int `json:"cluster_id"`
	Main_domain string `json:"main_domain"`
	Environment string `json:"environment"`
	Webmaster string `json:"webmaster"`
	Repository_branch string `json:"repository_branch"`
}

type Payload struct {
	Website Website_Set `json:"website"`
}

type Update_Maindomain struct {
	Main_domain string `json:"main_domain"`
}

type Update_env struct {
	Environment string `json:"environment"`
}

type Update_webmaster struct {
	Webmaster string `json:"webmaster"`
}

type Update_branch struct {
	Branch string `json:"repository_branch"`
}

type Update_php struct {
	Php string `json:"php_version"`
}


type Website_Update struct {
	*Update_Maindomain;
	*Update_env;
	*Update_webmaster;
	*Update_branch;
	*Update_php;
}
