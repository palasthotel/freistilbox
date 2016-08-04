package api

import "strconv"

type ClusterAttributes struct {
	Description string
	Name string
	Esla bool
}

type Relation struct {
	Id int
	Type string
}

type RelationList []Relation

type RelationCollection struct {
	Data RelationList
}

type ClusterRelations struct {
	Websites RelationCollection
}

type Cluster struct {
	Id int
	Attributes ClusterAttributes
	Relationships ClusterRelations
}

type ClusterList []Cluster

type ClusterCall struct {
	Data ClusterList;
}

type WebsiteAttributes struct {
	Clientid string
	Environment string
	Maindomain string
	Repositorybranch string
	Repositoryurl string
	Siteno int
	Webmaster string
}

type WebsiteRelations struct {
	Auxdomains RelationCollection
	Databases RelationCollection
	Sshkeys RelationCollection
}

type Website struct {
	Attributes WebsiteAttributes
	Id int
}

type WebsiteCall struct {
	Data Website
}

func (data ClusterList) Len() int {
	return len(data);
}

func (data ClusterList) Less(i, j int) bool {
	a := data[i];
	b := data[j];
	var valA,_ = strconv.Atoi(a.Attributes.Name[1:]);
	var valB,_ = strconv.Atoi(b.Attributes.Name[1:]);
	return valA < valB;
}

func (data ClusterList) Swap(i,j int) {
	temp := data[i];
	data[i]=data[j];
	data[j]=temp;
}

func (data RelationList) Len() int {
	return len(data);
}
