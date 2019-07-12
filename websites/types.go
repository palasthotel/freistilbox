package websites

type Website_Set struct {
	Cluster_id string `json:"cluster_id"`
	Main_domain string `json:"main_domain"`
	Environment string `json:"environment"`
	Webmaster string `json:"webmaster"`
	Repository_branch string `json:"repository_branch"`
}

type Payload struct {
	Website Website_Set `json:"website"`
}